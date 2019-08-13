### 242.有效的字母异位词

#### 题目来源[LeeCode242](https://leetcode-cn.com/problems/valid-anagram/description/)

---

给定两个字符串 *s* 和 *t* ，编写一个函数来判断 *t* 是否是 *s* 的一个字母异位词。

**示例 1:**

```
输入: s = "anagram", t = "nagaram"
输出: true
```

**示例 2:**

```
输入: s = "rat", t = "car"
输出: false
```

**说明:**
你可以假设字符串只包含小写字母。

**进阶:**
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

---

**方法一**：采用HashMap存储各个字符及其数量，然后对另外一个字符进行匹配

```java
import java.util.HashMap;
import java.util.Map;
import java.util.Set;
class Solution {
    public boolean isAnagram(String s, String t) {
		if (s.length() != t.length()) {
			return false;
		}
		Map<Character, Integer> map = new HashMap<>();
		for (int i = 0; i < s.length(); i++) {
			if (map.containsKey(s.charAt(i))) {
				map.put(s.charAt(i), map.get(s.charAt(i)) + 1);
			} else {
				map.put(s.charAt(i), 1);
			}
		}
		for (int i = 0; i < t.length(); i++) {
			if (map.containsKey(t.charAt(i))) {
				map.put(t.charAt(i), map.get(t.charAt(i)) - 1);
			} else {
				return false;
			}
		}
		Set<Character> keys = map.keySet();
		for (Character character : keys) {
			if (map.get(character) != 0) {
				return false;
			}
		}
		return true;
	}
}
```

**方法二**：采用数组进行存储，与采用HashMap同理

```java
public boolean isAnagram2(String s, String t) {
		if (s.length() != t.length()) {
			return false;
		}
		int[] container = new int[26];
		for (int i = 0; i < s.length(); i++) {
			container[s.charAt(i) - 'a']++;
			container[t.charAt(i) - 'a']--;
		}
		for (int i : container) {
			if (i != 0) {
				return false;
			}
		}
		return true;
	}
```

