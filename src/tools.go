package src

// max 获取两个数中大的一个
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

// 两者取最小
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateLinkedList ...
func CreateLinkedList(in []int) *ListNode {
	var head, tmp *ListNode
	for i, val := range in {
		if i == 0 {
			tmp = &ListNode{
				Val: val,
			}
			head = tmp
		} else {
			tmp.Next = &ListNode{
				Val: val,
			}
			tmp = tmp.Next
		}
	}
	return head
}

// WalkLinkedList ...
func WalkLinkedList(head *ListNode) []int {
	ret := []int{}
	for tmp := head; tmp != nil; tmp = tmp.Next {
		ret = append(ret, tmp.Val)
	}
	return ret
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CreateTree ...
func CreateTree(in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}
	var root, tmp *TreeNode
	tmp = &TreeNode{
		Val: in[0],
	}
	queue := []*TreeNode{tmp}
	root = tmp
	for i := 1; i <= len(in)/2; i++ {
		cur := queue[0]
		if len(queue) == 1 {
			queue = []*TreeNode{}
		} else {
			queue = queue[1:]
		}
		cur.Left = &TreeNode{
			Val: in[i*2-1],
		}
		queue = append(queue, cur.Left)
		if i*2 < len(in) {
			cur.Right = &TreeNode{
				Val: in[i*2],
			}
			queue = append(queue, cur.Right)
		}
	}
	return root
}

func walkTree(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		if len(queue) == 1 {
			queue = []*TreeNode{}
		} else {
			queue = queue[1:]
		}
		ret = append(ret, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return ret
}
