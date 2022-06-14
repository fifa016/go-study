/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-14 22:30:20
 */
package l

import "sort"

func merge(intervals [][]int) [][]int {
	res := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {

		right := intervals[i][1]

		if len(res) > 0 && res[len(res)-1][1] >= intervals[i][0] {
			right = max(res[len(res)-1][1], right)
			res[len(res)-1][1] = right
		} else {
			res = append(res, intervals[i])
		}

	}

	return res
}
