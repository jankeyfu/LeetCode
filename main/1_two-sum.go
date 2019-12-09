package main

/**
map + 一遍遍历解决问题
*/
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		if v, ok := mp[target-num]; ok {
			return []int{v, i}
		}
		mp[num] = i
	}
	return nil
}
