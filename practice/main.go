package main

import (
	"fmt"
	"practice/tennis"
)

const matchCount = 3

func main() {
	result := make(map[string][]int)
	oneMatch := tennis.Match {}
	oneMatch.Player1 = tennis.Player{"player 1", 70}
	oneMatch.Player2 = tennis.Player{"player 2", 80}
	
	for i:=0; i < matchCount; i++ {
		oneMatch.Start()
		result[oneMatch.Player1.Name] = append(result[oneMatch.Player1.Name], oneMatch.Score[oneMatch.Player1.Name])
		result[oneMatch.Player2.Name] = append(result[oneMatch.Player2.Name], oneMatch.Score[oneMatch.Player2.Name])
	}
	fmt.Printf("%s: %d-%d-%d \n", oneMatch.Player2.Name, result[oneMatch.Player1.Name][0], result[oneMatch.Player1.Name][1], result[oneMatch.Player1.Name][2])
	fmt.Printf("%s: %d-%d-%d \n", oneMatch.Player2.Name, result[oneMatch.Player2.Name][0], result[oneMatch.Player2.Name][1], result[oneMatch.Player2.Name][2])
}