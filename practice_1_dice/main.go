package main

import (
	"dice/dice"
	"dice/input"
	"dice/output"
	"fmt"
)

const Throws int = 10

func main() {	
	diceObj := dice.Init(input.GetNumFacesFromConsole()) //input 1000

	output.PrintRow(diceObj.RollNTimes(Throws)) //They will be different
	output.PrintRow(diceObj.RollNTimes(Throws)) //They will be different
	output.PrintRow(diceObj.RollNTimes(Throws+5)) //They will be different, but first 10 values are equal to the previous arrays
	for i:=0; i<Throws; i++ {
		fmt.Printf("%v ", diceObj.Roll())
	}
	fmt.Print("\n")
	for i:=0; i<Throws; i++ {
		fmt.Printf("%v ", dice.GetRandomNum(diceObj)) 
	}
}