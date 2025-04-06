package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLCA(curr *TreeNode) (*TreeNode, int) {
	if curr == nil {
		return nil, 0
	}

	leftLCA, leftDepth := findLCA(curr.Left)
	rightLCA, rightDepth := findLCA(curr.Right)

	if leftDepth == rightDepth {
		return curr, leftDepth + 1
	} else if leftDepth > rightDepth {
		return leftLCA, leftDepth + 1
	} else {
		return rightLCA, rightDepth + 1
	}
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	lca, _ := findLCA(root)
	return lca
}
