/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-08 03:03:57
 */
package l

func maxProfit(prices []int) int {
	res := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > res {
			res = prices[i] - min
		}

	}

	return res
}




func maxProfitTest20220608(prices []int) int {
	max := 0
	min := prices[i]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i] - min > max {
			max = prices[i] - min
		}
	}


	return  max

}