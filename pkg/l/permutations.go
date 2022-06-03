/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-03 20:31:56
 */
package l

import "sort"

func Permute(nums []int) [][]int {
	res := [][]int{}
	flag := make([]int, len(nums))
	list := []int{}

	var dfs func()
	dfs = func() {
		if len(list) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if flag[i] == 1 {
				continue
			}
			flag[i] = 1
			list = append(list, nums[i])
			dfs()

			flag[i] = 0
			list = list[:len(list)-1]
		}
	}

	dfs()

	return res
}
func PermuteNoRepeat(nums []int) [][]int {
	res := [][]int{}
	flag := make([]int, len(nums))
	list := []int{}

	var dfs func()
	dfs = func() {
		if len(list) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && flag[i-1] == 0 { //如果和前一个值相同，并且前一个还没有flag，肯定会重复
				continue
			}
			if flag[i] == 1 {
				continue
			}
			flag[i] = 1
			list = append(list, nums[i])
			dfs()

			flag[i] = 0
			list = list[:len(list)-1]
		}
	}
	sort.Ints(nums)
	dfs()

	return res
}

func Subsets(nums []int) [][]int {
	res := [][]int{}
	flag := make([]int, len(nums))
	list := []int{}

	var dfs func(start int)
	dfs = func(start int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			if flag[i] == 1 {
				continue
			}
			flag[i] = 1
			list = append(list, nums[i])
			dfs(i + 1)

			flag[i] = 0
			list = list[:len(list)-1]
		}
	}

	dfs(0)

	return res
}
