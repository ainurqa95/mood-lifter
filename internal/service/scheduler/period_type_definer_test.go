package scheduler

import (
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/ainurqa95/mood-lifter/internal/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

type TestCase struct {
	name            string
	hour            int
	expectedPeriods []int
}

var testCases = []TestCase{
	{
		name:            "six",
		hour:            6,
		expectedPeriods: nil,
	},
	{
		name:            "seven",
		hour:            7,
		expectedPeriods: nil,
	},
	{
		name:            "eight",
		hour:            8,
		expectedPeriods: []int{1, 2, 4},
	},
	{
		name:            "nine",
		hour:            9,
		expectedPeriods: []int{1, 3},
	},
	{
		name:            "ten",
		hour:            10,
		expectedPeriods: []int{1, 2, 5},
	},
	{
		name:            "eleven",
		hour:            11,
		expectedPeriods: []int{1},
	},
	{
		name:            "twelve",
		hour:            12,
		expectedPeriods: []int{1, 2, 3, 4, 6, 12},
	},
	{
		name:            "thirteen",
		hour:            13,
		expectedPeriods: []int{1},
	},
	{
		name:            "fourteen",
		hour:            14,
		expectedPeriods: []int{1, 2},
	},
	{
		name:            "fifteen",
		hour:            15,
		expectedPeriods: []int{1, 3, 5},
	},
	{
		name:            "seventeen",
		hour:            17,
		expectedPeriods: []int{1},
	},
	{
		name:            "eighteen",
		hour:            18,
		expectedPeriods: []int{1, 2, 3, 6},
	},
	{
		name:            "nineteen",
		hour:            19,
		expectedPeriods: nil,
	},
	{
		name:            "twelve",
		hour:            20,
		expectedPeriods: nil,
	},
}

func TestPeriodTypeDefiner_DefinePeriods(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	clock := mock.NewMockClock(ctr)
	cfg := config.Config{
		EndHour:   18,
		StartHour: 8,
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			currentDate := time.Now()
			newTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), tc.hour, 0, 0, 0, currentDate.Location())
			clock.EXPECT().Now().Return(newTime)
			periodDefiner := NewSchedulerPeriodDefiner(cfg, clock)
			periods := periodDefiner.DefinePeriods()
			assert.Equal(t, tc.expectedPeriods, periods)
		})
	}
}
