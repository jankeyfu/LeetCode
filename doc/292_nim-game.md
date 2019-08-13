### 292. Nim Game
#### 题目来源 [LeetCode 292](https://leetcode.com/problems/nim-game/)
---
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;You are playing the following Nim Game with your friend: There is a heap of stones on the table, each time one of you take turns to remove 1 to 3 stones. The one who removes the last stone will be the winner. You will take the first turn to remove the stones.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Both of you are very clever and have optimal strategies for the game. Write a function to determine whether you can win the game given the number of stones in the heap.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;For example, if there are 4 stones in the heap, then you will never win the game: no matter 1, 2, or 3 stones you remove, the last stone will always be removed by your friend.

**翻译**：

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;你正在和你的朋友玩一个叫Nim的游戏，游戏规则如下：桌上有一堆石头，每次你可以从桌上拿1~3块石头。谁拿走桌上最后一块石头谁就赢了。你先手开始拿石头。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;你们两个都是很聪明的人，编写一个方法，给定桌上的石头数量，判断你是否能赢得这次游戏。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;例如，如果在那堆石头一共有四块，那么你就无法赢得比赛，无论你拿1，2,3块石头，留下的石头你的朋友都能一次拿走。

---

**分析**：

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;你第一个拿，而且可以拿1~3块石头，所以桌上石头少于等于三块的时候你赢。当桌上有四块石头的时候，无论你拿几块都赢不了。当桌上有五块石头的时候，你拿掉石头后只要保证桌上有四块石头，你就赢了，你的朋友此时的状况就相当于你先拿的时候有4块石头。而当桌上有5块石头的时候，你同样保证桌上剩4块石头就赢了。以此类推，当桌上有8块石头的时候，你无法让桌上剩下4块石头，所以你无法赢得游戏,以此类推，在石头数大于8的时候，只要保证你拿完石头后能留下8块就赢了。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;总结上面的分析，当桌上石头的数量是4的倍数的时候，你永远无法赢。即n%4==0时，return false;其他情况return true;

**源代码**：

```
public class Solution {
    public boolean canWinNim(int n) {
		boolean isWin = false;
		if(n%4!=0){
			isWin = true;
		}		
		return isWin;
	}
}
```