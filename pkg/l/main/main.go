/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 23:12:29
 */
package main

import (
	"fmt"
	"go-study/pkg/l"
)

func main() {
	fmt.Println("============= start run ============")
	testSort()
	fmt.Println("============= end ============")
}

func testSort() {
	nums := []int{8, 4, 5, 7, 1, 3, 6, 2}
	// nums := []int{9, 3, 4, 8, 1, 2, 5, 7}
	l.MergeSort(nums)
	fmt.Println(nums)
}

func testPathSum() {
	root := &l.TreeNode{

		Val: 1,
		Left: &l.TreeNode{
			Val: 2,
		},
		Right: &l.TreeNode{
			Val: 3,
		},
	}
	res := l.HasPathSum(root, 4)
	fmt.Println(res)

}

func testInorderTraversal() {
	root := &l.TreeNode{
		Val:  1,
		Left: nil,
		Right: &l.TreeNode{
			Val: 2,
			Left: &l.TreeNode{
				Val: 3,
			},
		},
	}
	l.InorderTraversal(root)
}

func testBinarySearch() {
	nums := []int{1, 3, 4, 7, 8}
	target := 3
	fmt.Println(l.BinarySearch(nums, target))
}

func testQuickSort() {
	nums := []int{2, 1, 3, 4, 5}
	l.QuickSort(nums)
	fmt.Println(nums)
}

func testIsValidParenthenes() {
	s := "{}"
	fmt.Println(l.IsValid(s))
}
