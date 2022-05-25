package l

func InorderTraversal(root *TreeNode) []int {
    res := []int{}
	
	var tra2 func(root *TreeNode)

	tra2 = func(root *TreeNode) {
		if root == nil{
			return 
		}
		tra2(root.Left)
		res = append(res, root.Val)
		tra2(root.Right)
	}

    tra2(root)
    return res
}
