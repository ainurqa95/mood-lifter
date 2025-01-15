package scheduler

import (
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/ainurqa95/mood-lifter/internal/service/clock"
)

type PeriodTypeDefiner struct {
	cfg   config.Config
	clock clock.Clock
}

func NewSchedulerPeriodDefiner(cfg config.Config, clock clock.Clock) PeriodTypeDefiner {
	return PeriodTypeDefiner{
		cfg:   cfg,
		clock: clock,
	}
}

func (c *PeriodTypeDefiner) DefinePeriods() []int {
	now := c.clock.Now()
	currentHour := now.Hour()
	if currentHour < c.cfg.StartHour || currentHour > c.cfg.EndHour {
		return nil
	}
	var periodTypes []int
	for _, periodType := range model.AvailableSchedulePeriodTypes {
		if currentHour%periodType == 0 {
			periodTypes = append(periodTypes, periodType)
		}
	}

	return periodTypes
}
