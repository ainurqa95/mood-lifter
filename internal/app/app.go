package app

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/bots"
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/ainurqa95/mood-lifter/internal/service/compliment"
	"github.com/ainurqa95/mood-lifter/internal/service/scheduler"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type App struct {
	serviceProvider *serviceProvider
	bot             bots.BotManager
	massSender      compliment.MassSender
	scheduler       scheduler.ComplimentScheduler
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		a.bot.Start(ctx)
	}()

	go func() {
		err := a.scheduler.StartScheduler(ctx)
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	return a.scheduler.ShutDown()
}

func (a *App) initDeps(ctx context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	inits := []func(context.Context, config.Config) error{
		a.initServiceProvider,
		a.initBot,
		a.initMassSender,
		a.initScheduler,
	}

	for _, f := range inits {
		err := f(ctx, *cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context, cfg config.Config) error {
	dbPool := connectToDb(ctx, cfg)
	a.serviceProvider = newServiceProvider(cfg, dbPool)
	return nil
}

func connectToDb(ctx context.Context, cfg config.Config) *pgxpool.Pool {
	conf, err := pgxpool.ParseConfig(cfg.DbConfig.GetSource())
	if err != nil {
		log.Fatal("Error parse config", err)
	}
	ppool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		log.Fatal("Error parse config", err)
	}

	return ppool
}

func (a *App) initBot(_ context.Context, cfg config.Config) error {
	bot, err := bots.DefineBot(cfg,
		a.serviceProvider.UserService(),
		a.serviceProvider.ComplimentService(),
		a.serviceProvider.MessageService(),
	)

	if err != nil {
		return err
	}
	a.bot = bot
	return nil
}

func (a *App) initMassSender(ctx context.Context, cfg config.Config) error {
	a.massSender = compliment.NewMassSender(
		a.bot,
		a.serviceProvider.UserService(),
	)

	return nil
}

func (a *App) initScheduler(ctx context.Context, cfg config.Config) error {
	cron, err := scheduler.NewComplimentScheduler(cfg, a.massSender)
	if err != nil {
		return err
	}
	a.scheduler = *cron

	return nil
}
