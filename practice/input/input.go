package input

import "fmt"

//GetSettings fills the hash table with values
func GetSettings(settings *map[string]string) {
	var param string
	fmt.Println("You can input value or skip push on enter")
	for key, val := range (*settings) {
		param = "empty"
		if (val == "") {
			val = "Not required"
		} else {
			val = "default " + val
		}
		fmt.Printf("input %s param (%s): ", key, val)
		fmt.Scanln(&param)
		if(param != "empty") {
			(*settings)[key] = param
		}
	}
}