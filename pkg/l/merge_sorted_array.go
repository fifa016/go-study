/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 19:19:19
 */
package l

/**
 * @name: 2022-05
 * @param {[]int} nums1
 * @param {int} m
 * @param {[]int} nums2
 * @param {int} n
 * @return {*}
 */
func merge3(nums1 []int, m int, nums2 []int, n int) {

	n1 := m - 1
	n2 := n - 1

	index := m + n - 1

	for n1 >= 0 && n2 >= 0 {
		if n1 == -1 {
			nums1[index] = nums2[n2]
			n2--
		} else if n2 == -1 {
			nums1[index] = nums1[n1]
			n1--
		} else if nums1[n1] > nums2[n2] {
			nums1[index] = nums1[n1]
			n1--
		} else {
			nums1[index] = nums2[n2]
			n2--
		}
		index--
	}

}
