package input

import "fmt"

func GetLanguageFromConsole() string {
	var Language string
	fmt.Print("Language: ")
	fmt.Scan(&Language)
	return Language
}

func GetCommandFromConsole() int {
	var Command int
	fmt.Print("You: ")
	fmt.Scan(&Command)
	return Command
}