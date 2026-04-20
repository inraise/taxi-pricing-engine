package internal

import "time"

func GetTimeMultiplier(t time.Time) float64 {
	hour := t.Hour()
	isWeekend := t.Weekday() == time.Saturday || t.Weekday() == time.Sunday

	switch {
	case hour >= 0 && hour < 5:
		return 1.5
	case hour >= 7 && hour < 10 && !isWeekend:
		return 1.3
	case isWeekend:
		return 1.2
	default:
		return 1.0
	}
}
