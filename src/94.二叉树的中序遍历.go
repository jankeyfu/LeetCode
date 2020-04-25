package src

import "fmt"

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (69.18%)
 * Likes:    378
 * Dislikes: 0
 * Total Accepted:    97.5K
 * Total Submissions: 139.5K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [1,3,2]
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 */

// @lc code=start
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
1. 从根节点开始，节点入栈，不断读取左节点，直到左节点为空
*/
func inorderTraversal(root *TreeNode) []int {
	cur := root
	stack := []*TreeNode{}
	ret := []int{}
	isVisted := false
	for cur != nil {
		if !isVisted {
			ret = append(ret, cur.Val)
			stack = append(stack, cur)
		}
		fmt.Println(cur, cur.Left, stack)

		if !isVisted && cur.Left != nil {
			cur = cur.Left
			isVisted = false
		} else if cur.Right != nil {
			// 出栈cur,入栈cur.Right
			stack = stack[:len(stack)-1]
			cur = cur.Right
			isVisted = false
		} else {
			if len(stack) == 0 {
				break
			}
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			isVisted = true
		}
	}
	return ret
}

// @lc code=end
