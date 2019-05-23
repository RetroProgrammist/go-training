package data

import (
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

/*
	Contains the fields and methods necessary to obtain data from the site openweathermap.org
*/
type Meteorologist struct {
	source string
	apiKey string
	settings map[string]string
}

const source = "http://api.openweathermap.org/data/2.5/weather"

//Construct init Meteorologist obj
func Construct(settings map[string]string, apiKey string) (Meteorologist, error) {
	if _, ok := settings["city"]; (!ok || len(apiKey) < 1) {
		return Meteorologist{}, errors.New("input data not valid")
	}
 	return Meteorologist {
		 source: source,
		 apiKey: apiKey,
		 settings: settings,
		}, nil
}

func (m Meteorologist) GetWeather() (Weather, error) {
	var weatherObj = Weather{}

	//fmt.Println(m.getRequestURL())
	resp, err := http.Get(m.getRequestURL())
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Weather{}, err
	}

	err = json.Unmarshal(data, &weatherObj)
	if err != nil {
		return Weather{}, err
	}

	return weatherObj, nil
}

func (m Meteorologist) getRequestURL() string { //Dislike implementation 
	reqStr := "?" //len = 1
	for key, val := range m.settings {
		if (val != "") {
			if (key == "city") {
				if (m.settings["code"] != "") {
					reqStr = reqStr + "q=" + val + "." + m.settings["code"]
				} else {
					reqStr = reqStr + "q=" + val
				}
			} else if (key != "code") {
				reqStr = reqStr + key + "=" + val
			}
			reqStr = reqStr + "&"
		}	
	}
	if (len(reqStr) > 1) {
		return m.source + reqStr + "appid=" + m.apiKey
	} else {
		return ""
	}
}