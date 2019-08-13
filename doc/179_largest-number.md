### 179.最大的数

题目来源：[LeetCode179](https://leetcode-cn.com/problems/largest-number/description/)

---

给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

**示例 1:**

```
输入: [10,2]
输出: 210
```

**示例 2:**

```
输入: [3,30,34,5,9]
输出: 9534330
```

**说明:** 输出结果可能非常大，所以你需要返回一个字符串而不是整数。

---

**解题思路**：本题的主要思路是判断哪个数在前面得到的整数更大，比如`3`和`34`很明显我们知道`343`比`334`要大，因此`34`应该大于`3`，可以通过将两个整数组合后比较来判断哪个数在这种情况下更大。

```java
public String largestNumber(int[] nums) {
		for (int i = 0; i < nums.length - 1; i++) {
			for (int j = i + 1; j < nums.length; j++) {
				if (compare(nums[i], nums[j]) <= 0) {
					int tmp = nums[i];
					nums[i] = nums[j];
					nums[j] = tmp;
				}
			}
		}
		StringBuffer sb = new StringBuffer();
		for (int i = 0; i < nums.length; i++) {
			sb.append(nums[i]);
		}

		return sb.toString().startsWith("0") ? "0" : sb.toString();
	}

	public int compare(int num1, int num2) {
		String result1 = String.valueOf(num1) + String.valueOf(num2);
		String result2 = String.valueOf(num1) + String.valueOf(num2);
		return result1.compareTo(result2);
	}
```

