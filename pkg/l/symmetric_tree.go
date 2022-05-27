/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 22:28:47
 */
package l

func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)

}
