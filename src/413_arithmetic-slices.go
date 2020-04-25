package src

import "fmt"

/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sum := 0
	left := 0
	is := false
	for i := 0; i < len(A)-2; i++ {
		if A[i]+A[i+2] == 2*A[i+1] {
			if !is {
				left = i
				is = true
			}
		} else {
			fmt.Println(i, A[i], is, left)
			if is {
				sum += (i + 1 - left) * (i - left) / 2
				is = false
			}
		}

	}
	if is {
		sum += (len(A) - left - 1) * (len(A) - left - 2) / 2
	}
	return sum
}

// @lc code=end
