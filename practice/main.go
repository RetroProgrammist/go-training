package main

import (
	"practice/bot"
	"practice/input"
	"fmt"
)


func main() {
	
	robot := bot.Constructor("josy") //setName

	for err := robot.SetLanguage(input.GetLanguageFromConsole()); err != ""; err = robot.SetLanguage(input.GetLanguageFromConsole()) {
		if( err != "" ) { fmt.Println(err) }
	}

	for status:=true; status; {
		var answer string
		answer, status = robot.GetAnswer(input.GetCommandFromConsole())
		fmt.Println(answer)
	}
}