package list

import (
	. "../leetcode"
	"sort"
	"strconv"
)

// LeetCode(id=671,title=二叉树中第二小的节点,difficulty=easy)
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil || root.Right == nil {
		return -1
	}
	vals := []int{root.Val, root.Left.Val, root.Right.Val}
	var getNodeVals func(node *TreeNode)
	getNodeVals = func(node *TreeNode) {
		if node.Left != nil && node.Right != nil {
			vals = append(vals, node.Left.Val)
			vals = append(vals, node.Right.Val)
			getNodeVals(node.Left)
			getNodeVals(node.Right)
		}
	}
	getNodeVals(root.Left)
	getNodeVals(root.Right)
	sort.Ints(vals)
	if vals[0] == vals[len(vals)-1] {
		return -1
	}
	for _, v := range vals {
		if v > vals[0] {
			return v
		}
	}
	return -1
}

// LeetCode(id=669,title=修剪二叉搜索树,difficulty=easy)
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}

var bst []int

// LeetCode(id=653,title=两数之和 IV - 输入 BST,difficulty=easy)
func findTarget(root *TreeNode, k int) bool {
	bst = make([]int, 0)
	bst2Array(root)
	for i, j := 0, len(bst)-1; i < j; {
		sum := bst[i] + bst[j]
		if sum == k {
			return true
		} else if sum > k {
			j--
		} else {
			i++
		}
	}
	return false
}

func bst2Array(node *TreeNode) {
	if node == nil {
		return
	}
	bst2Array(node.Left)
	bst = append(bst, node.Val)
	bst2Array(node.Right)
}

// LeetCode(id=617,title=合并二叉树,difficulty=easy)
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 != nil && t2 != nil {
		t := &TreeNode{Val: t1.Val + t2.Val}
		t.Left = mergeTrees(t1.Left, t2.Left)
		t.Right = mergeTrees(t1.Right, t2.Right)
		return t
	} else if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	return nil
}

// LeetCode(id=606,title=根据二叉树创建字符串,difficulty=easy)
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	s := strconv.Itoa(t.Val)
	if t.Left != nil || t.Right != nil {
		s += "(" + tree2str(t.Left) + ")"
		if t.Right != nil {
			s += "(" + tree2str(t.Right) + ")"
		}
	}
	return s
}

// LeetCode(id=572,title=另一个树的子树,difficulty=easy)
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	sub := compareTree(s, t)
	if sub {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func compareTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val == t.Val {
		return compareTree(s.Left, t.Left) && compareTree(s.Right, t.Right)
	}
	return false
}

var maxDiameter int

// LeetCode(id=543,title=二叉树的直径,difficulty=easy)
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	diameterHelper(root)
	return maxDiameter
}

func diameterHelper(node *TreeNode) int {
	if node == nil {
		return 0
	}
	max := diameterHelper(node.Left)
	right := diameterHelper(node.Right)
	diameter := max + right
	if diameter > maxDiameter {
		maxDiameter = diameter
	}
	if right > max {
		max = right
	}
	return max + 1
}

var mini []int

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
