package l

import "fmt"

// func main() {
//     t1 := &TreeNode{1,nil,nil}
//     t2 := &TreeNode{2,nil,nil}
//     t3 := &TreeNode{3,nil,nil}
//     t4 := &TreeNode{4,nil,nil}
//     t5 := &TreeNode{5,nil,nil}
//     t7 := &TreeNode{7,nil,nil}

//     t1.Left = t2
//     t1.Right = t3
//     t2.Left = t4
//     t2.Right = t5
//     t3.Right = t7

//     //preOrderTraverse(t1)
//     postOrderTraverse(t1)
// }

func preOrderTraverse(root *TreeNode) {
    if root == nil{
        return
    }
    fmt.Println(root.id)
    preOrderTraverse(root.Left)
    preOrderTraverse(root.Right)

}
func inOrderTraverse(root *TreeNode) {
    if root == nil{
        return
    }
    inOrderTraverse(root.Left)
    fmt.Println(root.id)
    inOrderTraverse(root.Right)
}

func postOrderTraverse(root *TreeNode) {
    if root == nil{
        return
    }
    postOrderTraverse(root.Left)
    postOrderTraverse(root.Right)
    fmt.Println(root.id)

}


type TreeNode struct {
    id int
    Left *TreeNode
    Right *TreeNode
}

