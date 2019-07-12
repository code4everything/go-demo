package list

import (
	. "../leetcode"
	"fmt"
	"testing"
)

func TestLevelOrderBottom(test *testing.T) {
	node := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}
	levelOrderBottom(node)
}

func TestGetMinimumDifference(test *testing.T) {
	node := &TreeNode{Val: 1, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 3}}}
	fmt.Println(getMinimumDifference(node))
}
