### 344. Reverse String
#### 题目来源 [LeetCode 344](https://leetcode.com/problems/reverse-string/)
---
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Write a function that takes a string as input and returns the string reversed.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;编写一个将输入的字符串反向输出的方法。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Given s = "hello",&nbsp; return "olleh".输入s="hello" 返回值为"olleh";

**方法一**：通过StringBuffer集成的方法reverse()方法。先将String类型的字符串通过new StringBuffer(String str)的构造方法实例化一个StringBuffer对象，然后调用其中的方法reverse()实现反转，调用toString()方法形成一个String类型的字符串赋值给str，最后返回str;
```
public class Solution {
    public String reverseString(String s) {
        String str = null;
		StringBuffer sBuffer = new StringBuffer(s);
		str = sBuffer.reverse().toString();
		return str;
    }
}
```
**方法二**：通过toArrayChar()方法，将字符串转化为字符数组。然后循环交换字符数组的位置，最后将字符数组通过new String(char[] c)的方法转化为字符串。
```
public class Solution {
    public String reverseString(String s) {
        char[] c =s.toCharArray();
        char tempC;
        for(int i=0;i<c.length/2;i++){
        	tempC = c[i];
        	c[i]=c[c.length-i-1];
        	c[c.length-i-1] = tempC;
        }
        s = new String(c);
        return s;
    }
}
```