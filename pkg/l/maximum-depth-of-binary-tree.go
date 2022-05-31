/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-31 16:40:31
 */
package l

func maxDepth2(root *TreeNode) int {
	res := 0
	var tra func(*TreeNode, int)
	tra = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		depth++
		if depth > res {
			res = depth
		}
		tra(root.Left, depth)
		tra(root.Right, depth)
		depth--

	}

	tra(root, 0)
	return res
}
