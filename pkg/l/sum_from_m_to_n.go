/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-01 16:39:07
 */

package l

func getSumMN(nums []int, m, n int) int {

	res := 0

	for i := m; i <= n; i++ {
		res += nums[i]
	}

	return res
}

var cache []int = []int{}

func getSumMN2(nums []int, m, n int) int {

	if n > len(nums) || m > len(nums) || m > n {
		return -1
	}

	len := len(cache)

	for i := len; i <= n - len + 1; i++ {
		cache[i] = cache[i - 1] + nums[i]
	}

	return cache[n] - cache[m]

}
