/*
 * @Descripttion: 
 * @Author: jzh
 * @Date: 2022-06-02 16:59:37
 */
package l


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


type NaryTreeNode struct {
	Val int
	children []*NaryTreeNode
}