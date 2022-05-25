package main

import (
    "fmt"
    "go-study/pkg/l"
)

func main() {
    fmt.Println("============= start run ============")
    testIsValidParenthenes()
    fmt.Println("============= end ============")
}

func testBinarySearch() {
    nums := []int{1,3,4,7,8}
    target := 3
    fmt.Println(l.BinarySearch(nums, target))
}

func testQuickSort() {
    nums := []int{2,1,3,4,5}
    l.QuickSort(nums)
    fmt.Println(nums)
}

func testIsValidParenthenes() {
	s:="{}"
	fmt.Println(l.IsValid(s))
}
