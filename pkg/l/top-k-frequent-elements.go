/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-09 22:52:06
 */
package l

import (
	"fmt"
	"sort"
)

func TopKFrequent(nums []int, k int) []int {
	res := []int{}

	countMap := map[int]int{}

	for _, num := range nums {
		countMap[num]++
	}

	for num := range countMap {
		res = append(res, num)
	}

	fmt.Println(res)
	fmt.Println(countMap)

	sort.Slice(res, func(n1, n2 int) bool {
		// if countMap[n1] == countMap[n2] {
		// 	return res[n1] <= res[n2]
		// } else {
		return countMap[res[n1]] > countMap[res[n2]]
		// }
	})

	return res[0:k]

}
