package bot

import "time"

/*getTime returns time, depending on language*/
func getTime(lang string) string {
	var local string
	switch lang {
		case "en":
			local = "Europe/London"
		case "rus":
			local = "Europe/Minsk"
		default:
			return time.Now().Local().Format("15:04:05")
	}
	location, err := time.LoadLocation(local)
	if(err != nil) {panic(err)} //catch error
	return time.Now().In(location).Format("15:04:05")
}

/*getDate returns date in formate Month day, year*/
func getDate() string {
	return time.Now().Format("January _2, 2006")
}

/*getWeekDay returns day of the week*/
func getWeekDay() string {
	return time.Now().Format("Monday")
}