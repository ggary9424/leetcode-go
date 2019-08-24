package leetcode_go

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val >= L && root.Val <= R {
		treeLeft := trimBST(root.Left, L, R)
		root.Left = treeLeft

		treeRight := trimBST(root.Right, L, R)
		root.Right = treeRight

		return root
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}

	return trimBST(root.Left, L, R)
}
