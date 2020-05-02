package daily

func singleNumbers(nums []int) []int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}
	mask := 1
	for ret&1 != 1 {
		ret = ret >> 1
		mask = mask << 1
	}
	left := []int{}
	right := []int{}
	for _, num := range nums {
		if (num^ret)&mask == mask {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	// fmt.Println(left, right)
	i0, i1 := 0, 0
	for _, item := range left {
		i0 ^= item
	}

	for _, item := range right {
		i1 ^= item
	}
	return []int{i0, i1}
}

// 011
// 100
// 111
