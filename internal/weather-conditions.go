package internal

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	multiplier := 1.0

	switch weather.Condition {
	case Rain:
		multiplier += 0.125
	case HeavyRain:
		multiplier += 0.2
	case Snow:
		multiplier += 0.15
	case Clear:
	}

	if weather.WindSpeed > 15 {
		multiplier += 0.1
	}

	return multiplier
}
