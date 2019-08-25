package main

/**
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
*/
func restoreIpAddresses(s string) []string {
	if len(s) > 12 || len(s) < 4 {
		return nil
	}
	ips := ipValidate("", s, 1)
	ret := []string{}
	ipMap := map[string]bool{}
	for _, ip := range ips {
		if _, ok := ipMap[ip]; !ok {
			ipMap[ip] = true
			ret = append(ret, ip)
		}
	}
	return ret
}

func ipValidate(ip, s string, depth int) []string {
	var ret []string
	if depth == 4 {
		if !isSegmentValidate(s) {
			return nil
		}
		return []string{ip + "." + s}
	}

	for i := 1; i <= 3 && i < len(s); i++ {
		if !isSegmentValidate(s[:i]) {
			continue
		}
		tmp := ""
		if depth == 1 {
			tmp = s[:i]
		} else {
			tmp = "." + s[:i]
		}
		ip := ipValidate(ip+tmp, s[i:], depth+1)
		if len(ip) != 0 {
			ret = append(ret, ip...)
		}
	}
	return ret
}

func isSegmentValidate(s string) bool {
	if len(s) > 3 {
		return false
	}
	// 超过255
	if len(s) == 3 && !(s[0] <= '2' && s[1] <= '5' && s[2] < '5') {
		return false
	}
	// 01这种不符合条件的
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	return true
}

/**
回溯法，其实也就是递归不断试算哪个符合条件，并且在不满足的情况下不再进行后续的尝试
如25525511135
2 		5525511135
2.5		525511135
2.5.5	25511135
2.5.5.25511135 		不符合条件
2.5.52	5511135
2.5.52.5511135		不符合条件
...
直到找到符合条件的为止，
判断是否符合条件的方法由isSegmentValidate实现，大哥ip小段超过255或者是类似于010这种的都是不合法的。
*/
