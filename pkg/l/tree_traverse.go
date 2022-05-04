package l

import "fmt"

// func main() {
//     t1 := &TreeNode{1,nil,nil}
//     t2 := &TreeNode{2,nil,nil}
//     t3 := &TreeNode{3,nil,nil}
//     t4 := &TreeNode{4,nil,nil}
//     t5 := &TreeNode{5,nil,nil}
//     t7 := &TreeNode{7,nil,nil}

//     t1.left = t2
//     t1.right = t3
//     t2.left = t4
//     t2.right = t5
//     t3.right = t7

//     //preOrderTraverse(t1)
//     postOrderTraverse(t1)
// }

func preOrderTraverse(root *TreeNode) {
    if root == nil{
        return
    }
    fmt.Println(root.id)
    preOrderTraverse(root.left)
    preOrderTraverse(root.right)

}
func inOrderTraverse(root *TreeNode) {
    if root == nil{
        return
    }
    inOrderTraverse(root.left)
    fmt.Println(root.id)
    inOrderTraverse(root.right)
}

func postOrderTraverse(root *TreeNode) {
    if root == nil{
        return
    }
    postOrderTraverse(root.left)
    postOrderTraverse(root.right)
    fmt.Println(root.id)

}


type TreeNode struct {
    id int
    left *TreeNode
    right *TreeNode
}

