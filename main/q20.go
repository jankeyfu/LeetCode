package main

func isValid(s string) bool {
	stack := []rune{}
	opMp := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if len(stack) == 0 {
			return false
		} else {
			if opMp[stack[len(stack)-1]] == c {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
