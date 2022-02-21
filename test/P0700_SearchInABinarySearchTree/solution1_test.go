package P0700_SearchInABinarySearchTree

func searchBST(root *TreeNode, val int) *TreeNode {
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
