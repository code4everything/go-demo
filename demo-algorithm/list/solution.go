package list

import (
	. "../leetcode"
	"strconv"
)

var mini []int

// LeetCode(id=538,title=把二叉搜索树转换为累加树,difficulty=easy)
func convertBST(root *TreeNode) *TreeNode {
	convertHelper(root, 0)
	return root
}

func convertHelper(node *TreeNode, base int) int {
	if node == nil {
		return base
	}
	val := node.Val + convertHelper(node.Right, base)
	node.Val = val
	val = convertHelper(node.Left, val)
	return val
}

// LeetCode(id=520,title=二叉搜索树的最小绝对差,difficulty=easy)
func getMinimumDifference(root *TreeNode) int {
	mini = make([]int, 0)
	inorder(root)
	diff := mini[1] - mini[0]
	for i := 2; i < len(mini); i++ {
		tmp := mini[i] - mini[i-1]
		if tmp < diff {
			diff = tmp
		}
	}
	return diff
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	mini = append(mini, node.Val)
	inorder(node.Right)
}

var treePath []string

// LeetCode(id=257,title=二叉树的所有路径,difficulty=easy)
func binaryTreePaths(root *TreeNode) []string {
	treePath = make([]string, 0)
	if root == nil {
		return treePath
	}
	pathHelper("", "", root)
	return treePath
}

func pathHelper(pre string, sep string, node *TreeNode) {
	str := pre + sep + strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		treePath = append(treePath, str)
		return
	}
	if node.Left != nil {
		pathHelper(str, "->", node.Left)
	}
	if node.Right != nil {
		pathHelper(str, "->", node.Right)
	}
}

// LeetCode(id=226,title=翻转二叉树,difficulty=easy)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// LeetCode(id=111,title=二叉树的最小深度,difficulty=easy)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var min int
	if root.Left != nil {
		min = minDepth(root.Left)
	}
	if root.Right != nil {
		tmp := minDepth(root.Right)
		if min == 0 || tmp < min {
			min = tmp
		}
	}
	return 1 + min
}

var levels [][]int

// LeetCode(id=100,title=二叉树的层次遍历 II,difficulty=easy)
func levelOrderBottom(root *TreeNode) [][]int {
	levels = make([][]int, 0)
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
