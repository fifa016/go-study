/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-03 18:04:34
 */


package l



func KthSmallest(root *TreeNode, k int) int {
	res := -1
	
	var tra func(root *TreeNode)
	tra = func (root *TreeNode)  {
		if root == nil {
			return 
		}
		tra(root.Left )
		
		k --
		if 0 == k {
			res = root.Val
			return 
		}


		tra(root.Right)
	}

	tra(root)



	return res

}