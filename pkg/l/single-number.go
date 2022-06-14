/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-06 16:49:47
 */

package l

//一个数字出现了一次，其他的数字出现了两次
func singleNumber(nums []int) int {
	res := 0
	for _, val := range nums {
		res ^= val
	}

	return res

}

//一个元素出现1次，其他出现3次
//
func singleNumberII(nums []int) int {
	res := int32(0)
	for i := 0; i < 32; i++ {
		bitSum := 0
		for _, val := range nums {
			bitSum += val >> i & 1
		}
		res |= int32(bitSum) % 3 << i
	}

	return int(res)
}

//两个元素出现1次，其他出现2次
func singleNumberIII(nums []int) []int {
	res1, res2 := 0, 0

	xorRes := 0
	for _, val := range nums {
		xorRes ^= val
	}
	firstOneBit := -1
	for i := 0; i < 32; i++ {
		if xorRes>>i&1 == 1 {
			firstOneBit = i
			break
		}
	}

	for _, val := range nums {
		if val>>firstOneBit&1 == 1 {
			res1 ^= val
		} else {
			res2 ^= val
		}
	}

	return []int{res1, res2}
}
