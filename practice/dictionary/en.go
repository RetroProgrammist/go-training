package dictionary

/*
	en dictionary 
*/

type english struct {
	phrases []string
	exception string
	len int
}

func (en english) init() english {
	en.phrases = []string {
		"Hello, I am #NAME#",
		"Now in London is #TIME#",
		"Today is #DATE#",
		"Today is #WEEK#",
		"Bye",
	}

	en.exception = "You can use 1-#LEN# commands"

	en.len = len(en.phrases)
	return en
}

func (en english) GetPhraseByNum(num int) string {
	return en.phrases[num]
}

func (en english) GetLen() int {
	return en.len
}

func (en english) GetException() string {
	return en.exception
}