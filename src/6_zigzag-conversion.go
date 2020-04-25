package src

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/

// func convert(s string, numRows int) string {
// 	if numRows == 1 {
// 		return s
// 	}
// 	// get actual numRows
// 	var rows [][]rune
// 	for i := 0; i < int(math.Min(float64(numRows), float64(len(s)))); i++ {
// 		rows = append(rows, []rune{})
// 	}
// 	cr := 0
// 	// for get the first rowstep +1，and later to be -1
// 	var rowStep = -1
// 	for _, c := range s {
// 		rows[cr] = append(rows[cr], c)
// 		if cr == 0 || cr == numRows-1 {
// 			rowStep = -rowStep
// 		}
// 		cr += rowStep
// 	}
// 	var buff bytes.Buffer
// 	for _, row := range rows {
// 		buff.WriteString(string(row))
// 	}
// 	return buff.String()
// }

/**
解题思路一：逐个字符解析到每一行的数组中
LEETCODEISHIRING
L   C   I   R
E T O E S I I G
E   D   H   N

如：设定一个长度为numrows的slice数组rows
依次读取L E E T
cr表示当前行，当cr==0的时候，cr++,当cr == numrows时，cr--
rows[0]:{L} 	cr = 0
rows[1]:{E} 	cr = 1
rows[2]:{E} 	cr = 2
rows[1]:{E, T} 	cr = 1
rows[0]:{L, C} 	cr = 0
cr来回的在[0, numrows-1]递增或者递减
*/

/**
解题思路二：分析每一行的字符在字符串中索引的位置
LEETCODEISHIRING
L     D     R
E   O E   I I
E C   I H   N
T     S     G

按照上述进行编号得
0        6	 	  12
1     5  7     11 13
2  4     8  10	  14
3	     9	      15

由上图可知：每一行都含有 (numRows - 1) * 2 * n + cr 这个编号的数据
处理第0行和第numRows-1行之外，其余行在上述每两个编号数据中间还有一个数：(numRows - cr - 1) * 2 + cr
第0行： 0, 3*2*1, 3*3*2
第1行： 1, (4-1-1)*2+1, 3*2*1+1, (4-1-1)*2*2+1
...
*/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ret := make([]byte, len(s))
	idx := 0
	commonStep := (numRows - 1) << 1
	for cr := 0; cr < numRows; cr++ {
		for i := cr; i < len(s); i += commonStep {
			ret[idx] = s[i]
			idx++
			if cr != 0 && cr != (numRows-1) {
				secStep := (numRows - cr - 1) << 1
				if i+secStep < len(s) {
					ret[idx] = s[i+(numRows-cr-1)<<1]
					idx++
				}
			}
		}
	}
	return string(ret)
}
