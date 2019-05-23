package data

//Weather is response from the site openweathermap.org
type Weather struct {
	Coord struct {
		Lon float64 //City geo location, longitude
		Lan float64 //City geo location, latitude
	} `json:"-"`
	Weather []struct { //(more info Weather condition codes)
		Id          int `json:"-"`//Weather condition id
		Main        string //Group of weather parameters (Rain, Snow, Extreme etc.)
		Description string //Weather condition within the group
		Icon        string `json:"-"`//Weather icon id
	}
	Base string `json:"-"` //Internal parameter
	Main struct {
		Temp      float64 //Temperature. Unit Default: Kelvin, Metric: Celsius, Imperial: Fahrenheit.
		Pressure  float64 `json:"-"`//Atmospheric pressure (on the sea level, if there is no sea_level or grnd_level data), hPa
		Humidity  int //Humidity, %
		TempMin   float64 `json:"Temp_min"` //Minimum temperature at the moment.
		TempMax   float64 `json:"Temp_man"` //Maximum temperature at the moment. 
		SeaLevel  float64 `json:"sea_level"` //Atmospheric pressure on the sea level, hPa
		GrndLevel float64 `json:"grnd_level"` //Atmospheric pressure on the ground level, hPa
	}
	Wind struct { 
		Speed float64 //Wind speed. Unit Default: meter/sec, Metric: meter/sec, Imperial: miles/hour.
		Deg   float64 //Wind direction, degrees (meteorological)
		Gust  float64 //Not sure about it param
	}
	Clouds struct {
		All int //Cloudiness, %
	}
	Rain struct { 
		H1 int `json:"1h"`//Rain volume for the last 1 hour, mm
		H3 int `json:"3h"`//Rain volume for the last 3 hours, mm
	} `json:"-"`
	Snow struct { 
		H1 int `json:"1h"`//Snow volume for the last 1 hour, mm
		H3 int `json:"3h"`//Snow volume for the last 3 hours, mm
	} `json:"-"`
	Dt  int `json:"-"` //Time of data calculation, unix, UTC
	Sys struct {
		Type    int `json:"-"`//Internal parameter
		Id      int `json:"-"`//Internal parameter
		Message float64 `json:"-"`//Internal parameter
		Country string `json:"-"`//Country code (GB, JP etc.)
		Sunrise int64 //Sunrise time, unix, UTC
		Sunset  int64 //Sunset time, unix, UTC
	}
	Id   int `json:"-"`// ity ID
	Name string //City name
	Cod  int //Internal parameter
}


func (w Weather) GetNameCity() string {
	return w.Name
}

func (w Weather) GetTimeSunriseAndSunset() (sunrise int64, sunset int64) {
	return w.Sys.Sunrise, w.Sys.Sunset
}

func (w Weather) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	return w.Main.Temp, w.Main.TempMin, w.Main.TempMax
}

func (w Weather) GetCloudiness() (description string) {
	if (len(w.Weather) > 0) {
		for _, v := range w.Weather {
			if (len(v.Description) > 0) {
				return v.Description
			}
		}
	}
	return ""
}

func (w Weather) GetHumidity() (humidity int) {
	return w.Main.Humidity
}

func (w Weather) GetWind() (speed float64, gust float64, direction string) {

	blowDirection := map[float64]string{
		0:   "северный", //if missing
		45:  "северо-восточный",
		90:  "восточный",
		135: "юго-восточный",
		180: "южный",
		225: "юго-западный",
		270: "западный",
		315: "северо-западный",
	}

	//w.Wind.Deg - can be 0, when be missing
	deg := float64(w.Wind.Deg) - 22
	for key, val := range blowDirection {
		if ((key-22) <= deg && deg <= (key+23)) {
			direction = val
		}
	}

	return float64(w.Wind.Speed), w.Wind.Gust, direction
}