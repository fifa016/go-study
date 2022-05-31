/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-31 16:25:56
 */
package l

func preorderTraversal(root *TreeNode) []int {
	res := []int{}

	var tra func(*TreeNode)
	tra = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)
		tra(root.Left)
		tra(root.Right)

	}

	tra(root)
	return res
}
