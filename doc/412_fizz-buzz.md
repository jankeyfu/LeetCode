### 412. Fizz Buzz
#### 题目来源 [LeetCode 412](https://leetcode.com/problems/fizz-buzz/)
---

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Write a program that outputs the string representation of numbers from 1 to n. But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;编写一个程序,输出的字符串表示数字从1到n。
但对3的倍数应该输出"Fizz",而不是输出数字，对于5的倍数的数应该输出"Buzz",对数字3和5的倍数应该输出"FizzBuzz".

```
Example:例子

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
```
**源码如下：**
```

public class Solution {
   public List<String> fizzBuzz(int n) {
		List<String> list = new ArrayList<String>();
		for(int i=1;i<=n;i++){
			
			if(i%3==0 && i%5==0){
				list.add("FizzBuzz");
			}else if(i%3 == 0){
				list.add("Fizz");
			}else if(i%5==0){
				list.add("Buzz");
			}else{
				list.add(""+i);
			}
		}
		return list;
	}
}
```