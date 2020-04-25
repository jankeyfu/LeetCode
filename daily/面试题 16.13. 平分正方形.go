package daily

func cutSquares(square1 []int, square2 []int) []float64 {
	sqmidx1 := float64(square1[0]) + float64(square1[2])/2
	sqmidy1 := float64(square1[1]) + float64(square1[2])/2
	sqmidx2 := float64(square2[1]) + float64(square2[2])/2
	sqmidy2 := float64(square2[1]) + float64(square2[2])/2
	rate := [2]int{}
	if sqmidx1 == sqmidx2 && sqmidy1 == sqmidy2 { // 45°
		rate = [2]int{1, 1}
	} else {
		rate = [2]int{int(2 * (sqmidx2 - sqmidx1)), int(2 * (sqmidy2 - sqmidy1))}
	}
	
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
