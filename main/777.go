package main

import (
	"strings"
)

// "XXXLXXXXXX"
// "XXXLXXXXXX"
// 777
// 1. 长度必须相等
// 2. LR的顺序不能改变
// 3. end 中的对应L只能比其当前位置更靠左边，R只能比其当前位置更靠右边
func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}
	if strings.ReplaceAll(start, "X", "") != strings.ReplaceAll(end, "X", "") {
		return false
	}
	j := 0
	k := 0
	for i, sc := range start {
		if sc == 'L' {
			for end[j] != 'L' {
				j++
			}
			if i < j {
				return false
			}
			j++
		}
		if sc == 'R' {
			for end[k] != 'R' {
				k++
			}
			if i > k {
				return false
			}
			k++
		}
	}

	return true
}
