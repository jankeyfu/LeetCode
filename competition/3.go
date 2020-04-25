package competition

import (
	"sort"
)

func getTriggerTime(increase [][]int, requirements [][]int) []int {
	cur := [3]int{0, 0, 0}

	ret := make([]int, len(requirements), len(requirements))
	for i := 0; i < len(requirements); i++ {
		ret[i] = -1
		requirements[i] = append(requirements[i], i)
	}
	sort.Slice(requirements, func(i, j int) bool {
		if requirements[i][0] == requirements[j][0] {
			if requirements[i][1] == requirements[j][1] {
				return requirements[i][2] <= requirements[j][2]
			}
			return requirements[i][1] < requirements[j][1]
		}
		return requirements[i][0] < requirements[j][0]

	})
	// fmt.Println(requirements, increase)
	for day, inc := range increase {
		cur[0] += inc[0]
		cur[1] += inc[1]
		cur[2] += inc[2]
		for _, req := range requirements {
			// fmt.Println(cur, req)
			if req[req[3]] != -1 {
				continue
			}
			if cur[0] < req[0] {
				break
			}
			if cur[0] == req[0] && cur[1] < req[1] {
				break
			}
			if cur[0] == req[0] && cur[1] == req[1] && cur[2] < req[2] {
				break
			}
			if req[1] == 0 && req[2] == 0 && req[3] == 0 {
				ret[req[3]] = 0
			}
			if cur[0] >= req[0] && cur[1] >= req[1] && cur[2] >= req[2] && ret[req[3]] == -1 {
				ret[req[3]] = day + 1
			}
		}
	}
	return ret
}
