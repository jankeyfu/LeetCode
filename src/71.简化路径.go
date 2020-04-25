package src

import "strings"

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 *
 * https://leetcode-cn.com/problems/simplify-path/description/
 *
 * algorithms
 * Medium (37.86%)
 * Likes:    97
 * Dislikes: 0
 * Total Accepted:    24.1K
 * Total Submissions: 62.5K
 * Testcase Example:  '"/home/"'
 *
 * 以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。
 *
 * 在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..）
 * 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。更多信息请参阅：Linux / Unix中的绝对路径 vs 相对路径
 *
 * 请注意，返回的规范路径必须始终以斜杠 / 开头，并且两个目录名之间必须只有一个斜杠 /。最后一个目录名（如果存在）不能以 /
 * 结尾。此外，规范路径必须是表示绝对路径的最短字符串。
 *
 *
 *
 * 示例 1：
 *
 * 输入："/home/"
 * 输出："/home"
 * 解释：注意，最后一个目录名后面没有斜杠。
 *
 *
 * 示例 2：
 *
 * 输入："/../"
 * 输出："/"
 * 解释：从根目录向上一级是不可行的，因为根是你可以到达的最高级。
 *
 *
 * 示例 3：
 *
 * 输入："/home//foo/"
 * 输出："/home/foo"
 * 解释：在规范路径中，多个连续斜杠需要用一个斜杠替换。
 *
 *
 * 示例 4：
 *
 * 输入："/a/./b/../../c/"
 * 输出："/c"
 *
 *
 * 示例 5：
 *
 * 输入："/a/../../b/../c//.//"
 * 输出："/c"
 *
 *
 * 示例 6：
 *
 * 输入："/a//b////c/d//././/.."
 * 输出："/a/b/c"
 *
 */

// unix路径都是用`/`分隔的，所有可以先使用`/`将路径分隔为字符串数组
// 使用栈存储路径：
// 1. 碰到`.`表示当前路径，不需要处理，跳过即可
// 2. 碰到空格，表示是空路径，直接跳过
// 3. 碰到`..`表示路径的上一层，将栈定目录出栈
// 4. 其余字符表示路径目录名称，直接入栈。
// @lc code=start
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

// @lc code=end
