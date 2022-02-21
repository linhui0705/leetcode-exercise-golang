package P0700_SearchInABinarySearchTree

import "leetcode-exercise-golang/structure"

func searchBST(root *structure.TreeNode, val int) *structure.TreeNode {
	if nil == root {
		return nil
	}
	if val == root.Val {
		return root
	} else if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
