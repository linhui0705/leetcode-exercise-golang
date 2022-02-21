package P0700_SearchInABinarySearchTree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return insertNode(root, val)
}

func insertNode(node *TreeNode, val int) *TreeNode {
	if nil == node {
		return &TreeNode{Val: val}
	}
	if val < node.Val {
		if nil == node.Left {
			node.Left = &TreeNode{Val: val}
			// return node
		} else {
			insertNode(node.Left, val)
		}
	} else {
		if nil == node.Right {
			node.Right = &TreeNode{Val: val}
			// return node
		} else {
			insertNode(node.Right, val)
		}
	}
	return node
}
