package dice

import (
	"math/rand"
	"time"
)

type dice struct {
	faces int
}

/**
* It returns 'obj' type dice. 
**/

func Init(faces int) dice {
	return dice {faces: faces}
}

/**
* It returns an array with the results of dice throws for 2 dices
**/
func (d dice) RollNTimes(throws int) ([]int, []int) {
	var resultFirstDice, resultSecondDice []int 
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:= 0; i<throws; i++ {
		resultFirstDice = append(resultFirstDice, r.Intn(d.faces)+1)
		resultSecondDice = append(resultSecondDice, r.Intn(d.faces)+1)
	}
	return resultFirstDice, resultSecondDice
}
