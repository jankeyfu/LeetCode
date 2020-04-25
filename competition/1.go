package competition

func minCount(coins []int) int {
	num := 0
	for _, coin := range coins {
		num += (coin + 1) / 2
	}
	return num
}
