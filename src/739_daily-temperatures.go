package src

/**
https://leetcode-cn.com/problems/daily-temperatures/
739. 每日温度
根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高的天数。如果之后都不会升高，请输入 0 来代替。
例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的都是 [30, 100] 范围内的整数。
*/
func dailyTemperatures(T []int) []int {
	// save idx of element in array T
	stack := []int{}
	ret := make([]int, len(T))
	for i, v := range T {
		for top := len(stack) - 1; len(stack) != 0 && v > T[stack[top]]; top = len(stack) - 1 {
			ret[stack[top]] = i - stack[top]
			stack = stack[:top]
		}
		stack = append(stack, i)
	}
	for _, v := range stack {
		ret[v] = 0
	}
	return ret
}
