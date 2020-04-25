package competition

import "fmt"

// Run ...
func Run() {
	// fmt.Println(numWays(5, [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}, 3))
	// fmt.Println(numWays(3, [][]int{{0, 2}, {2, 1}}, 2))
	// fmt.Println(numWays(3, [][]int{{0, 2}, {2, 1}, {1, 2}}, 2))
	fmt.Println(getTriggerTime([][]int{{0, 4, 5}, {4, 8, 8}, {8, 6, 1}, {10, 10, 0}}, [][]int{{12, 11, 16}, {20, 2, 6}, {9, 2, 6}, {10, 18, 3}, {8, 14, 9}}))
	fmt.Println(getTriggerTime([][]int{{1, 1, 1}}, [][]int{{0, 0, 0}}))
	fmt.Println(getTriggerTime([][]int{{2, 8, 4}, {2, 5, 0}, {10, 9, 8}}, [][]int{{2, 11, 3}, {15, 10, 7}, {9, 17, 12}, {8, 1, 14}}))
}

// get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs
