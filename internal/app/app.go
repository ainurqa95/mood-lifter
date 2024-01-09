package app

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/bots"
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type App struct {
	serviceProvider *serviceProvider
	bot             bots.BotStarter
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
		complimentService := a.serviceProvider.ComplimentService()
		err := complimentService.MassCreate(ctx)
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	inits := []func(context.Context, config.Config) error{
		a.initServiceProvider,
		a.initBot,
	}

	for _, f := range inits {
		err := f(ctx, *cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context, cfg config.Config) error {
	dbPool := connectToDb(cfg)
	a.serviceProvider = newServiceProvider(cfg, dbPool)
	return nil
}

func (a *App) initBot(_ context.Context, cfg config.Config) error {
	bot, err := bots.DefineBot(cfg, a.serviceProvider.UserService(), a.serviceProvider.ComplimentService())
	if err != nil {
		return err
	}
	a.bot = bot
	return nil
}

func connectToDb(cfg config.Config) *pgxpool.Pool {
	conf, err := pgxpool.ParseConfig(cfg.DbConfig.GetSource())
	if err != nil {
		log.Fatal("Error parse config", err)
	}
	ppool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		log.Fatal("Error parse config", err)
	}

	return ppool
}
