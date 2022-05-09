package l

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	max := 0
	return tra(root, 0, max)
}

func tra(root *TreeNode, depth int, max int) int {
	if root == nil {
		return max
	}

	depth++
	if root.Left == nil && root.Right == nil && depth > max {
		return depth
	}

	leftMax := tra(root.Left, depth, max)
	if leftMax > max {
		max = leftMax
	}
	rightMax := tra(root.Right, depth, max)
	if rightMax > max {
		return rightMax
	} else {
		return max
	}
}
