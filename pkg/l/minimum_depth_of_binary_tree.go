package l

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	return traMinDepth(root, 0, 0)
}

func traMinDepth(root *TreeNode, depth int, min int) int {
	if root == nil {
		return min
	}

	depth++
	if root.Left == nil && root.Right == nil && (min == 0 || depth < min) {
		return depth
	}

	leftMin := traMinDepth(root.Left, depth, min)
	if min == 0 || leftMin < min {
		min = leftMin
	}
	rightMin := traMinDepth(root.Right, depth, min)
	if min == 0 || rightMin < min {
		return rightMin
	} else {
		return min
	}
}
