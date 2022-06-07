/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-07 14:53:15
 */
package l

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(preArr []int, preBegin, preEnd int, inArr []int, inBegin, inEnd int) *TreeNode {

	// get root
	root := &TreeNode{Val: preArr[preBegin]}
	// get left
	index := 0
	for inArr[index] != preArr[preBegin] {
		index++
	}
	root.Left = build(preArr, preBegin + 1, preBegin +  index - inBegin, inArr, inBegin, index -1)
	// get right
	root.Right = build(preArr, preBegin +  index - inBegin + 1, preEnd, inArr,  index  + 1, inEnd )

	return root

}
