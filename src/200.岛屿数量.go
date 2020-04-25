package src

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (47.98%)
 * Likes:    524
 * Dislikes: 0
 * Total Accepted:    94.7K
 * Total Submissions: 192.9K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 *
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 *
 * 此外，你可以假设该网格的四条边均被水包围。
 *
 * 示例 1:
 *
 * 输入:
 * 11110
 * 11010
 * 11000
 * 00000
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入:
 * 11000
 * 11000
 * 00100
 * 00011
 * 输出: 3
 * 解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
 *
 *
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	num := 0
	stack := [][]int{}
	for i, j := findFirsOne(grid); !(i == len(grid) && j == len(grid[0])); i, j = findFirsOne(grid) {
		// fmt.Println(i, j)
		stack = append(stack, []int{i, j})
		for len(stack) != 0 {
			cij := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			ci, cj := cij[0], cij[1]
			// fmt.Println(ci, cj, grid, stack)
			// time.Sleep(time.Millisecond * 100)
			grid[ci][cj] = '2'
			if ci+1 < len(grid) && grid[ci+1][cj] == '1' {
				stack = append(stack, []int{ci + 1, cj})
			}
			if cj+1 < len(grid[ci]) && grid[ci][cj+1] == '1' {
				stack = append(stack, []int{ci, cj + 1})
			}
			if ci-1 >= 0 && grid[ci-1][cj] == '1' {
				stack = append(stack, []int{ci - 1, cj})
			}
			if cj-1 >= 0 && grid[ci][cj-1] == '1' {
				stack = append(stack, []int{ci, cj - 1})
			}
		}
		num++
		// time.Sleep(time.Second * 1)
	}
	return num
}

func findFirsOne(grid [][]byte) (int, int) {
	// fmt.Println(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				return i, j
			}
		}
	}
	return len(grid), len(grid[0])
}

// @lc code=end
