package stats

import "dice/dice"//Есть ли более красивый способ указать тип параметра функции?

type statistic struct {
	sum []int
	rep []int
	perc []int
}

/**
*Generates the final data array (type Statistic)
**/
func GetStatistic(diceObj dice.Dice, throws int) statistic {
	stats := statistic {}
	sizeArray := diceObj.GetFaces()*2-1
	firstArray, secondArray := diceObj.RollNTimes(throws)
	sum, rep, persent := make([]int, sizeArray), make([]int, sizeArray), make([]int, sizeArray) //definitions array

	for i:=0; i<throws; i++ {
		rep[sumOfNumb(firstArray[i], secondArray[i])-2]++
	}

	for i:= 0; i<sizeArray; i++ {
		sum[i] = i+2
		persent[i] = probability(rep[i], throws)
	}

	stats.sum = sum
	stats.rep = rep
	stats.perc = persent

	return stats
}

func (s statistic) GetSum() []int {
	return s.sum
}

func (s statistic) GetRepetitions() []int {
	return s.rep
}

func (s statistic) GetPercent() []int {
	return s.perc
}

func sumOfNumb(num1, num2 int) int {
	return num1+num2
}

func probability(part, total int) int {
	return 100*part/total
}
