/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-26 15:29:02
 */
package l

// 2022-05
func candy(ratings []int) int {

	candy := make([]int, len(ratings))
	for i := 0; i < len(candy); i++ {
		candy[i] = 1
	}

	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			candy[i+1] = candy[i] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candy[i] <= candy[i+1] {
			candy[i] = candy[i+1] + 1
		}
	}

	res := 0
	for _, val := range candy {
		res = res + val
	}

	return res
}
