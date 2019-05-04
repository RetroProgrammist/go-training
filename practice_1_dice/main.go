package main

import (
	"dice/dice"
	"dice/stats"
	"dice/input"
	"dice/output"
)

const Throws int = 10000

func main() {	
	diceObj := dice.Init(input.GetNumFacesFromConsole()) //6 only
	firstArray, secondArray := diceObj.RollNTimes(Throws)
	result := stats.GetStatistic(firstArray, secondArray, Throws) //its work
	output.PrintRow(result.GetSum())
	output.PrintRow(result.GetRepetitions())
	output.PrintRow(result.GetPercent())
}