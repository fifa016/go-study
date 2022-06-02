/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-02 15:07:10
 */
package l

import "math"

func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	cache := []*TreeNode{}
	left := 0
	cache = append(cache, root)
	right := 1
	for left < right {
		max := math.MinInt
		for left < right {
			if cache[left].Val > max {
				max = cache[left].Val
			}
			if cache[left].Left != nil {
				cache = append(cache, cache[left].Left)
			}
			if cache[left].Right != nil {
				cache = append(cache, cache[left].Right)
			}
			left++
		}
		right = len(cache)
		res = append(res, max)
	}
	return res
}
