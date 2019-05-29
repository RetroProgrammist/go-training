package tennis

type Match struct {
	Player1, Player2 Player
	Score map[string]int
}

//Start - Play one match
func (m *Match) Start() map[string]int {
	var winName string
	var setScore map[string]int
	m.Score = make(map[string]int)
	m.Score[m.Player1.Name] = 0 //init Score
	m.Score[m.Player2.Name] = 0 //init Score 

	for ;m.Score[m.Player1.Name] < 6 && m.Score[m.Player2.Name] < 6; {
		setScore = make(map[string]int)
		setScore[m.Player1.Name] = 0 //init Score
		setScore[m.Player2.Name] = 0 //init Score 
		for ;setScore[m.Player1.Name] <= 40 && setScore[m.Player2.Name] <= 40; { 
			channel := make(chan string)
			for true {
				go m.Player1.Play(channel)
				if Name1, ok := <-channel; ok {
					winName = Name1
				} else {
					break
				}
				go m.Player2.Play(channel)
				if Name2, ok := <-channel; ok {
					winName = Name2
				} else {
					break
				}
			}
			if (len(winName) < 1) {continue}
			if (setScore[winName] < 30) {
				setScore[winName] += 15
			} else {
				setScore[winName] += 10
			}
		}

		if setScore[m.Player1.Name] > 40 {
			winName = m.Player1.Name
		} else {
			winName = m.Player2.Name
		}

		m.Score[winName]++
	}
	return m.Score
}