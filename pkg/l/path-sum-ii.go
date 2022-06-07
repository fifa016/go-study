/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-07 16:54:50
 */
package l

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	list := []int{}
	var tra func(r *TreeNode, target int)
	tra = func(r *TreeNode, target int) {
		if r == nil {
			return
		}
		target -= r.Val
		list = append(list, r.Val)

		defer func() { list = list[:len(list)-1] }()
		if r.Left == nil && r.Right == nil && target == 0 {
			res = append(res, append([]int{}, list...))
			return
		}
		tra(r.Left, target)
		tra(r.Right, target)
	}
	tra(root, targetSum)
	return res

}
