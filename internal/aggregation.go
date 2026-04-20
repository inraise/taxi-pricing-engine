package internal

import (
	"math"
	"time"
)

const (
	minPrice = 99.0
	maxPrice = 20000.0
)

func (c *PriceCalculator) CalculatePrice(trip TripParameters, now time.Time, weather WeatherData, lat, lng float64) float64 {
	base := CalculateBasePrice(trip)
	timeMult := GetTimeMultiplier(now)
	weatherMult := GetWeatherMultiplier(weather)
	trafficMult := GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMult * weatherMult * trafficMult

	return ApplyPriceLimits(math.Round(finalPrice))
}

func ApplyPriceLimits(price float64) float64 {
	if price < minPrice {
		return minPrice
	}

	if price > maxPrice {
		return maxPrice
	}
	return price
}
