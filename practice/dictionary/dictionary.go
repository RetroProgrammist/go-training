package dictionary

import (
	"strings"
)

type Dictionary interface {
	GetPhraseByNum(num int) string 
	GetLen() int
	GetException() string
}

func GetDictionary(lang string) (Dictionary, string) {
	switch strings.ToLower(lang) {
	case "en":
		return english{}.init(), ""
	case "rus":
		return russian{}.init(), ""
	default:
		return nil, "Try input 'en' or 'rus'"
	}
}