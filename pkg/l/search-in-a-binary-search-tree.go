/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-02 11:38:27
 */
package l

func searchBST(root *TreeNode, val int) *TreeNode {
	var res *TreeNode = nil
	var tra func(root *TreeNode, val int)
	tra = func(root *TreeNode, val int) {
		if root == nil {
			return
		}
		if root.Val == val {
			res = root
		} else if root.Val < val {
			tra(root.Right, val)
		} else {
			tra(root.Left, val)
		}
	}

	tra(root, val)

	return res

}
