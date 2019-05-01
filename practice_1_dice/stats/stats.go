package stats

type statistic struct {
	sum []int
	rep []int
	perc []int
}

/**
*Generates the final data array (type Statistic)
**/
func GetStatistic(firstArray, secondArray []int) statistic {
	stats := statistic {}
	throws := len(firstArray) //No matter which array for now
	sum := make([]int, 11)
	rep := make([]int, 11)
	persent := make([]int, 11)

	for i:=0; i<throws; i++ {
		rep[sumOfNumb(firstArray[i], secondArray[i])-2]++
	}

	for i:= 0; i<11; i++ {
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
