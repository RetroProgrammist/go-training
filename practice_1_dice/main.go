package main

import (
	"dice/dice"
	"dice/stats"
	"dice/input"
	"dice/output"
)

const Throws int = 10000

func main() {	
	diceObj := dice.Init(input.GetNumFacesFromConsole())
	result := stats.GetStatistic(diceObj, Throws)
	output.PrintRow(result.GetSum())
	output.PrintRow(result.GetRepetitions())
	output.PrintRow(result.GetPercent())
}