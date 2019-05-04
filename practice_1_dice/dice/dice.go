package dice

import (
	"math/rand"
	"time"
)

type Dice struct {
	faces int
}

/**
* It returns 'obj' type dice. 
**/
func Init(faces int) Dice {
	return Dice {faces: faces}
}

/**
* It returns an array with the results of dice throws for 2 dices
**/
func (d Dice) RollNTimes(throws int) ([]int, []int) { //фигня с рандомом, будто кэшируется результаты
	var resultFirstDice, resultSecondDice []int 
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:= 0; i<throws; i++ {
		resultFirstDice = append(resultFirstDice, r.Intn(d.faces)+1)
		resultSecondDice = append(resultSecondDice, r.Intn(d.faces)+1)
	}
	return resultFirstDice, resultSecondDice
}

func (d Dice) GetFaces() int{
	return d.faces
}