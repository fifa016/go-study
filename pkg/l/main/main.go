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
	testValidParenWithPriority()
	fmt.Println("============= end ============")
}

func testValidParenWithPriority() {

	arr := []string{

		"{}{}<><>()()[][]",
		"{{}}{{}}<<>><<>>(())(())[[]][[]]",
		"{{}}{{}}<<>><<>>(())(())[[]][[]]",
		"{<>}{[]}<<<>><<>>>((<>))(())[[(<>)]][[]]",
		"><}{{[]}<<<>><<>>>((<>))(())[[(<>)]][[]]",
	}

	for _, s := range arr {
		fmt.Println(l.IsValidWithPriority(s))

	}
}

func testLongestCommonSubstring() {
	s1 := "abcd"
	s2 := "cbca"
	res := l.LongestCommonSubstring(s1, s2)
	fmt.Println(res)
}

func testLongestPalindromicSubstring() {
	s := "aaaaa"
	res := l.LongestPalindrome(s)
	fmt.Println(res)
}

func testMaxSum() {
	list := []int{1, -5, 7, -4, 6, -9, 3, 6, -7}
	fmt.Println(l.MaxSumOfTwoSubArray(list))
}

func testPermute() {
	arr := []int{1, 2, 3, 4}
	res := l.Permute(arr)
	fmt.Println(res)
}

func testKthSmallest() {
	root := &l.TreeNode{
		Val: 3,
		Left: &l.TreeNode{
			Val:   1,
			Right: &l.TreeNode{Val: 2},
		},
		Right: &l.TreeNode{
			Val: 4,
		},
	}

	l.KthSmallest(root, 1)

}

func testReverseLinkedList() {
	head := &l.ListNode{
		Val: 3,
		Next: &l.ListNode{
			Val: 5,
		},
	}
	l.ReverseBetween(head, 1, 2)
}

func testLinkedListSort() {
	head := &l.ListNode{
		Val: 3,
		Next: &l.ListNode{
			Val: 2,
			Next: &l.ListNode{
				Val: 4,
			},
		},
	}

	res := l.LinkedListBubbleSort(head)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func testStrToInt() {
	input := "+-2"
	res := l.StrToInt(input)
	fmt.Println(res)

}

func testFindKth() {
	nums := []int{3, 2, 1, 5, 6, 4}
	res := l.FindKthLargest(nums, 2)
	fmt.Println(res)

}

func testSort() {
	nums := []int{8, 4, 5, 7, 1, 3, 6}
	// nums := []int{9, 3, 4, 8, 1, 2, 5, 7}
	l.MergeSortTest20220605(nums)
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

func testPreorderTraversal() {
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
	l.PreorderTraTest20220606(root)
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
func testPostorderTraversal() {
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
	l.PostorderTraversal(root)
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
