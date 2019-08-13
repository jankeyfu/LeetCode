package main

import "strings"

// unix路径都是用`/`分隔的，所有可以先使用`/`将路径分隔为字符串数组
// 使用栈存储路径：
// 1. 碰到`.`表示当前路径，不需要处理，跳过即可
// 2. 碰到空格，表示是空路径，直接跳过
// 3. 碰到`..`表示路径的上一层，将栈定目录出栈
// 4. 其余字符表示路径目录名称，直接入栈。
func simplifyPath(path string) string {
	pth := strings.Split(path, "/")
	stack := make([]string, 0, len(pth))
	for _, v := range pth {
		switch v {
		case ".":
			continue
		case "":
			continue
		case "..":
			if len(stack) == 0 {
				continue
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}
