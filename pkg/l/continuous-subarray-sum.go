/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-08 16:37:55
 */
package l

func checkSubarraySum(nums []int, k int) bool {

	kv := map[int]int{}
	kv[0] = 1
	sum := 0
	for i, v := range nums {
		sum += v
		if _, ok := kv[sum-k]; ok && kv[sum-k]-i >= 2 {
			return true
		} else {
			kv[sum-k] = i
		}
	}

	return false
}
