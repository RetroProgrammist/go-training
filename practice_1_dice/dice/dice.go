package dice

import (
	"math/rand"
	"time"
)

type Dice struct {
	faces int
	lastRandomNum int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
* It returns 'obj' type dice. 
*/
func Init(faces int) Dice {
	return Dice {faces: faces}
}

/* 
* It must return random number (with most entropia)
* But we have the same value.
*/
func (d Dice) Roll() int { 
	return rand.Intn(d.faces)+1
}

/*
* And here also
*/
func (d Dice) RollNTimes(throws int) []int {
	var result []int 
	for i:= 0; i<throws; i++ {
		result = append(result, rand.Intn(d.faces)+1)
	}
	return result
}

/*
* Another way
*/
func (d Dice) hideRoll() int { 
	return rand.Intn(d.faces)+1
}

/*Interface*/
func GetRandomNum(d Dice) int {
	return d.hideRoll()
}

func (d Dice) GetFaces() int{
	return d.faces
}