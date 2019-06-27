package list

import (
	. "../leetcode"
	"testing"
)

func TestLevelOrderBottom(test *testing.T) {
	node := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}
	levelOrderBottom(node)
}
