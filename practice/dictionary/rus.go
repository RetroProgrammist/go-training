package dictionary

/*
	rus dictionary 
*/

type russian struct {
	phrases []string
	exception string
	len int
}

func (rus russian) init() russian {
	rus.phrases = []string {
		"Привет, я #NAME#",
		"Сейчас в Минске #TIME#",
		"Сегодня #DATE#",
		"Сегодня #WEEK#",
		"До встречи",
	}

	rus.exception = "Вы можете использовать 1-#LEN# комманд"

	rus.len = len(rus.phrases)
	return rus
}

func (rus russian) GetPhraseByNum(num int) string {
	return rus.phrases[num]
}

func (rus russian) GetLen() int {
	return rus.len
}

func (rus russian) GetException() string {
	return rus.exception
}