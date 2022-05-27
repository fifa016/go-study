/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 22:55:51
 */
package l

func HasPathSum(root *TreeNode, targetSum int) bool {
	sum := 0
	has := false
	var tra func(*TreeNode)

	tra = func(root *TreeNode) {
		if root == nil {
			return
		}

		sum = sum + root.Val
		if root.Left == nil && root.Right == nil && sum == targetSum {
			has = true
		}
		tra(root.Left)
		tra(root.Right)

		sum = sum - root.Val
	}

	tra(root)

	return has
}
