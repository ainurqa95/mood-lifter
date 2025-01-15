package model

type SchedulePeriodType uint8

const (
	EveryHour       = 1
	EveryTwoHours   = 2
	EveryThreeHours = 3
	EveryFourHours  = 4
	EveryFiveHours  = 5
	EverySixHours   = 6
	EveryDay        = 12
)

var AvailableSchedulePeriodTypes = []int{
	EveryHour,
	EveryTwoHours,
	EveryThreeHours,
	EveryFourHours,
	EveryFiveHours,
	EverySixHours,
	EveryDay,
}
