/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-01 15:49:44
 */
package l

import "sync"

var rwMutex sync.RWMutex

func checkCoin(arrCoin []int, start, end int) int {
	res := -1

	var check func(arrCoin []int, start, end int)
	check = func(arrCoin []int, start, end int) {

		rwMutex.RLock()
		mid := start + (end-start)/2
		rwMutex.RUnlock()
		leftSum := getSum(arrCoin, start, mid)
		rightSum := getSum(arrCoin, mid+1, end)
		rwMutex.Lock()
		if leftSum < rightSum {
			if end - start == 1 {
				res = start
				return
			}
			checkCoin(arrCoin, start, mid)
			} else if leftSum > rightSum {
			if end - start == 1 {
				res = end
				return
			}
			checkCoin(arrCoin, mid+1, end)
		}
		rwMutex.Unlock()
	}

	check(arrCoin, start, end)

	return res
}

func getSum(arrCoin []int, start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += arrCoin[i]
	}
	return sum
}
