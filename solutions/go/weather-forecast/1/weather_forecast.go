// Package weather provides tools for reporting the current weather conditions.
package weather

var (
	// CurrentCondition is the current weather condition.
	CurrentCondition string
	// CurrentLocation is the current city name.
	CurrentLocation string
)

// Forecast returns a message with the current weather condition of the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}