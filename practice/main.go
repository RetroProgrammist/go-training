package main

import (
	"practice/input"
	weather "practice/openweathermap"
)

const apiKey = "2c19a8c670afc70f2ae7a81f229fce3d" //API key, more https://home.openweathermap.org/api_key

var settings = map[string]string{ //default values
	"city": "Minsk", //city name
	"code": "", //country code
	"lang": "ru", //language code
	"units": "metric", // Standard, metric, and imperial units are available.
}

type s struct {
	test int
}

func main() {
	input.GetSettings(&settings)
	weather.ShowForecast(settings, apiKey)
}