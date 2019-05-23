package formatdata

import (
	"fmt"
	"time"
	"practice/openweathermap/data"
)

type weatherForecast struct {
	dictionary []string
	formatWeather string
}

func Construct(settings map[string]string, apiKey string) (weatherForecast, error) {
	dic := []string{
		"Сегодня в городе %s",
		" %s",
		", температура воздуха %+.2f°С",
		", ветер %s",
		" %.0fм/с",
		" c порывами до %.0fм/с",
		". Влажность воздуха %d%%",
		". Восход солнца %s, заход солнца %s",
	}

	wf := weatherForecast{dictionary: dic}
	meteo, err := data.Construct(settings, apiKey)
	if (err != nil) {
		return weatherForecast{}, err
	}
	weatherObj, err := meteo.GetWeather()
	if (err != nil) {
		return weatherForecast{}, err
	}
	wf.setFormatWeather(weatherObj)

	return wf, nil
}

//FormatWeather write response in predefined format
func (wf *weatherForecast) FormatWeather() string {
	return (*wf).formatWeather
}

func (wf *weatherForecast) setFormatWeather(w data.Weather) {
	var formatWeather string
	if l:=len(w.GetNameCity()); l > 0 {
		formatWeather = fmt.Sprintf(wf.dictionary[0], w.GetNameCity())

		if l:=len(w.GetCloudiness()); l > 0 {
			formatWeather = formatWeather + fmt.Sprintf(wf.dictionary[1], w.GetCloudiness())
		}

		if temp, tempMin, tempMax := w.GetTemperature(); (temp != tempMin || temp != tempMax) {
			formatWeather = formatWeather + fmt.Sprintf(wf.dictionary[2], temp)
		}

		if  speed, gust, direc := w.GetWind(); direc != "" {
			formatWeather = formatWeather + fmt.Sprintf(wf.dictionary[3], direc)
			if (speed > 0) {
				formatWeather = formatWeather + fmt.Sprintf(wf.dictionary[4], speed)
			}

			if (gust > 0) {
				formatWeather = formatWeather + fmt.Sprintf(wf.dictionary[5], gust)
			}
		}

		if h:=w.GetHumidity(); h > 0 {
			formatWeather = formatWeather + fmt.Sprintf(wf.dictionary[6], w.GetHumidity())
		}

		if sunrise, sunset := w.GetTimeSunriseAndSunset(); (sunrise > 0 && sunset > 0) {
			formatWeather = formatWeather + fmt.Sprintf(wf.dictionary[7], time.Unix(sunrise, 0).Format("15:04"), time.Unix(sunset, 0).Format("15:04"))
		}
	}
	wf.formatWeather = formatWeather + "."
}
