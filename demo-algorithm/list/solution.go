package list

import . "../leetcode"

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
