package src

import (
	"math"
)

/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
*/
func minimumTotal(triangle [][]int) int {
	for row := 1; row < len(triangle); row++ {
		triangle[row][0] = triangle[row][0] + triangle[row-1][0]
		for col := 1; col < len(triangle[row])-1; col++ {
			triangle[row][col] = min(triangle[row][col]+triangle[row-1][col-1], triangle[row][col]+triangle[row-1][col])
		}
		triangle[row][row] = triangle[row][row] + triangle[row-1][row-1]
	}
	m := math.MaxInt32
	for i := 0; i < len(triangle); i++ {
		m = min(m, triangle[len(triangle)-1][i])
	}
	return m
}
