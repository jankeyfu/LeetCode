package main

import "strings"

func letterCombinations(digits string) []string {
	var ret []string
	digitMp := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	for _, s := range strings.Split(digits, "") {
		for i, c := range digitMp[s] {
			if len(ret) < i {
				ret = append(ret, c)
			} else {
				for i := range ret {
					ret[i] = ret[i] + c
				}
			}
		}
	}
	return ret
}
