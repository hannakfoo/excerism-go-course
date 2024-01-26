// Package weather is an package easy to use .
package weather

// CurrentCondition : string value to give to current condition.
var CurrentCondition string

// CurrentLocation : string value that represents the current location.
var CurrentLocation string

// Forecast : function that returns the current weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
