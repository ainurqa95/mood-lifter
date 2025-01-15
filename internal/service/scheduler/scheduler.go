package scheduler

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/ainurqa95/mood-lifter/internal/service/compliment"
	"github.com/go-co-op/gocron/v2"
	"log"
)

const everyCronOur = "0 * * * *"

type ComplimentScheduler struct {
	scheduler         gocron.Scheduler
	massSender        compliment.MassSender
	periodTypeDefiner PeriodTypeDefiner
	cfg               config.Config
}

func NewComplimentScheduler(cfg config.Config, massSender compliment.MassSender, periodTypeDefiner PeriodTypeDefiner) (*ComplimentScheduler, error) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &ComplimentScheduler{
		massSender:        massSender,
		scheduler:         scheduler,
		periodTypeDefiner: periodTypeDefiner,
		cfg:               cfg,
	}, nil
}

func (c *ComplimentScheduler) StartScheduler(ctx context.Context) error {
	_, err := c.scheduler.NewJob(
		gocron.CronJob(
			everyCronOur,
			false,
		),
		gocron.NewTask(
			func(ctx context.Context) {
				periodTypes := c.periodTypeDefiner.DefinePeriods()
				err := c.massSender.SendMassCompliments(ctx, periodTypes)
				if err != nil {
					log.Printf("error sending mass compliments: %v", err)
				}
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
