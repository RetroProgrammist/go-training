package bot

import (
	"strconv"
	"strings"
	"practice/dictionary"
)


/*Bot is robot, he understand next command:
* 1 - greeting
* 2 - current time
* 3 - current date
* 4 - 
*
*/
type bot struct {
	name string
	language string
	dic dictionary.Dictionary
}

func Constructor(name string) bot {
	b := bot{name: name}
	return b
}

/*SetLanguage is set Language for bot*/
func (b *bot) SetLanguage(lan string) string {//its createBot
	var err string
	b.language = lan

	b.dic, err = dictionary.GetDictionary(b.language)

	return err
}

/*GetAnswer return answer and status (if true - we continue, else break code), depending on command */
func (b bot) GetAnswer(command int) (string, bool) {
	switch command {
		case 1:
			return strings.Replace(b.dic.GetPhraseByNum(command-1), "#NAME#", b.name, 1), true
		case 2:
			return strings.Replace(b.dic.GetPhraseByNum(command-1), "#TIME#", getTime(b.language) , 1), true
		case 3:
			return strings.Replace(b.dic.GetPhraseByNum(command-1), "#DATE#", getDate(), 1), true
		case 4:
			return strings.Replace(b.dic.GetPhraseByNum(command-1), "#WEEK#", getWeekDay(), 1), true
		case 5:
			return b.dic.GetPhraseByNum(command-1), false
		default:
			return strings.Replace(b.dic.GetException(), "#LEN#", strconv.Itoa(b.dic.GetLen()), 1), true
	}
}
