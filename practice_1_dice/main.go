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
	result := stats.GetStatistic(diceObj.RollNTimes(Throws), Throws) //why? error "not enough arguments in call to stats.GetStatistic"
	output.PrintRow(result.GetSum())
	output.PrintRow(result.GetRepetitions())
	output.PrintRow(result.GetPercent())
}