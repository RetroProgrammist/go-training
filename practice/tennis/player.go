package tennis

import (
	"math/rand"
	//"fmt"
	"time"
)

type Player struct {
	Name string
	Skill int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Play - hit the ball or not
func (p Player) Play(ch chan string) {
	serve := rand.Intn(100)//his difficulty to hit the ball
	//fmt.Println(serve, p.Name)
	if (p.Skill < serve) {
		close(ch)
	} else {
		ch <- p.Name
	}
}