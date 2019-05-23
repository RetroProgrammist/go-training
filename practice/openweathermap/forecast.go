package openweathermap

import (
	"fmt"
	"practice/openweathermap/formatdata"
)

//ShowForecast print Current weather data in userfriendly format
func ShowForecast(settings map[string]string, apiKey string) {
	wf, err := formatdata.Construct(settings, apiKey)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println(wf.FormatWeather())
}