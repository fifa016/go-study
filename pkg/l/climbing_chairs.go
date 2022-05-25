package l

/**
 * 2022-05
 */
func climbStairs2(n int) int {
	fibArr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i <= 2 {
			fibArr[i] = i
			continue
		}
		fibArr[i] = fibArr[i - 1] + fibArr[i - 2]

	}

	return fibArr[n]
}

func climbStairs(n int) int {
	stairs := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			stairs[i] = 1
		} else if i == 1 {
			stairs[i] = 2
		} else {
			stairs[i] = stairs[i-1] + stairs[i-2]
		}

	}
	return stairs[n-1]
}
