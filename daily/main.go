package daily

import "fmt"

// Run ...
func Run() {
	// s := [][]int{[]int{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// s = [][]int{[]int{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	// rotate(s)
	// fmt.Println(s)
	fmt.Println(singleNumbers([]int{1, 2, 3, 1, 2, 4}))
	fmt.Println(singleNumbers([]int{1, 2, 5, 2}))
}
