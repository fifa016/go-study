package l

/**
 * 2022-05
 */
func twoSum2(nums []int, target int) []int {

	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if value, ok := numMap[target-nums[i]]; ok {
			return []int{i, value}
		}
		numMap[nums[i]] = i
	}
	return nil
}

func twoSum(arr []int, target int) []int {
	res := []int{0, 0}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}

	return res
}
