/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-15 11:06:20
 */
package l

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var res *TreeNode

	var find func(root, p, q *TreeNode) int
	find = func(root, p, q *TreeNode) int {

		if root == nil {
			return 0
		}

		val := find(root.Left, p, q) + find(root.Right, p, q)
		if root.Val == p.Val {
			val += 1
		}
		if root.Val == q.Val {
			val += 2
		}

		if val == 3 && res == nil {
			res = root
		}

		return val
	}

	return res
}
