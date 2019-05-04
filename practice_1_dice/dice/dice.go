package dice

import (
	"math/rand"
	"time"
)

type Dice struct {
	faces int
	lastRandomNum int
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
	time.Sleep(1) //fix
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(d.faces)+1
}

/*
* And here also
*/
func (d Dice) RollNTimes(throws int) []int {
	var result []int 
	time.Sleep(1) //fix
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:= 0; i<throws; i++ {
		result = append(result, r.Intn(d.faces)+1)
	}
	return result
}

/*
* Another way
*/
func (d Dice) hideRoll() int { 
	time.Sleep(1) //fix
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(d.faces)+1
}

/*Interface*/
func GetRandomNum(d Dice) int {
	return d.hideRoll()
}

func (d Dice) GetFaces() int{
	return d.faces
}