package P0098_ValidateBinarySearchTree

func isValidBST(root *TreeNode) bool {
	return traverse(root, nil, nil)
}

func traverse(node *TreeNode, min *TreeNode, max *TreeNode) bool {
	if nil == node {
		return true
	}
	if nil != min && min.Val >= node.Val {
		return false
	}
	if nil != max && node.Val >= max.Val {
		return false
	}
	return traverse(node.Left, min, node) && traverse(node.Right, node, max)
}
