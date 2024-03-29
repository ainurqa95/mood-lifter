package scheduler

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/ainurqa95/mood-lifter/internal/service/compliment"
	"github.com/go-co-op/gocron/v2"
)

type ComplimentScheduler struct {
	scheduler  gocron.Scheduler
	massSender compliment.MassSender
	cfg        config.Config
}

func NewComplimentScheduler(cfg config.Config, massSender compliment.MassSender) (*ComplimentScheduler, error) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &ComplimentScheduler{
		massSender: massSender,
		scheduler:  scheduler,
		cfg:        cfg,
	}, nil
}

func (c *ComplimentScheduler) StartScheduler(ctx context.Context) error {
	_, err := c.scheduler.NewJob(
		gocron.CronJob(
			c.cfg.CronSchedule,
			false,
		),
		gocron.NewTask(
			func(ctx context.Context) {
				c.massSender.SendMassCompliments(ctx)
			}, ctx,
		),
	)
	if err != nil {
		return err
	}

	c.scheduler.Start()

	return nil
}

func (c *ComplimentScheduler) ShutDown() error {
	return c.scheduler.Shutdown()
}
