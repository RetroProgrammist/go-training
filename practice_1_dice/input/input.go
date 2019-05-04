package input

import "fmt"

func GetNumFacesFromConsole() int {
	var faces int
	fmt.Print("input number of faces: ")
	fmt.Scan(&faces)
	fmt.Print("\nresult:\n")
	return faces
}