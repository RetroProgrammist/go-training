package main

import (
	"dice/dice"
	"dice/output"
	"dice/stats"
)

const Throws int = 10000

func main() {
	diceObj := dice.Init(6)
	result := stats.GetStatistic(diceObj.RollNTimes(Throws))
	output.PrintRow(result.GetSum())
	output.PrintRow(result.GetRepetitions())
	output.PrintRow(result.GetPercent())
}