package src

/*
实现函数 ToLowerCase()，该函数接收一个字符串参数 str，并将该字符串中的大写字母转换成小写字母，之后返回新的字符串。

示例 1：
输入: "Hello"
输出: "hello"
示例 2：

输入: "here"
输出: "here"
示例 3：

输入: "LOVELY"
输出: "lovely"
*/

func toLowerCase(str string) string {
	bits := make([]byte, 0, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] <= 'Z' && str[i] >= 'A' {
			bits = append(bits, str[i]+32)
		} else {
			bits = append(bits, str[i])
		}
	}
	return string(bits)
}

/*
题解：
	主要利用的是AscaII码的 'A'与'a' 值相差32
*/
