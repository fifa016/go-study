/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-02 15:00:46
 */
package l

/**
 * @name:
 * @param {*TreeNode} root
 * @param {int} k
 * @return {*}
 */
func kthLargest(root *TreeNode, k int) int {
	var res *TreeNode = nil
	var tra func(root *TreeNode, k int)
	count := 0
	tra = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		tra(root.Right, k)
		count++
		if count == k {
			res = root
			return
		}

		tra(root.Left, k)
	}
	tra(root, k)

	return res.Val
}
