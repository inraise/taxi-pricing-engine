package main

import (
	"fmt"
	"taxi-pricing-engine/internal"
	"time"
)

func main() {
	calculator := internal.PriceCalculator{
		TrafficClient: &internal.RealTrafficClient{},
	}

	price := calculator.CalculatePrice(
		internal.TripParameters{Distance: 8.5, Duration: 20},
		time.Now(),
		internal.WeatherData{Condition: internal.HeavyRain, WindSpeed: 10},
		55.751244, 37.618423,
	)

	fmt.Printf("Ваша цена: %.0f руб.\n", price)
}
