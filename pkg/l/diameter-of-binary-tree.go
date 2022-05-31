/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-31 21:40:02
 */

package l

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	var tra func(root *TreeNode) int
	tra = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			return 1
		}

		left := tra(root.Left)
		right := tra(root.Right)
		if left+right > res {
			res = left + right
		}
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}

	tra(root)

	return res
}
