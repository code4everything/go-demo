package list

import (
	. "../leetcode"
)

var levels = make([][]int, 0)

// LeetCode(id=100,title=二叉树的层次遍历 II,difficulty=easy)
func levelOrderBottom(root *TreeNode) [][]int {
	levelOrderHelper(root, 0)
	for i, j := 0, len(levels)-1; i < j; i++ {
		tmp := levels[i]
		levels[i] = levels[j]
		levels[j] = tmp
		j--
	}
	return levels
}

func levelOrderHelper(node *TreeNode, asc int) {
	if node == nil {
		return
	}
	if asc < len(levels) {
		levels[asc] = append(levels[asc], node.Val)
	} else {
		levels = append(levels, []int{node.Val})
	}
	levelOrderHelper(node.Left, asc+1)
	levelOrderHelper(node.Right, asc+1)
}

// LeetCode(id=100,title=相同的树,difficulty=easy)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

// LeetCode(id=83,title=删除排序链表中的重复元素,difficulty=easy)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	val := head.Val
	node := head.Next
	for node != nil {
		if node.Val == val {
			node = node.Next
			pre.Next = node
		} else {
			pre = node
			val = node.Val
			node = node.Next
		}
	}
	return head
}
