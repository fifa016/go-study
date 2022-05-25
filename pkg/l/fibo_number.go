package l

func fib2(n int) int{

	fibArr := make([]int, n + 1)

	for i := 0 ;i <= n ; i ++ {
		if i <= 1 {
			fibArr[i] = i
			continue
		}
		fibArr[i] = fibArr[i - 1] + fibArr[i - 2]
	}
	return fibArr[n]
}


func fib(n int) int {
    stairs := make([]int, n + 1)

    for i := 0; i <= n; i++ {
        if i == 0 {
            stairs[i] = 0
        } else if i == 1 {
            stairs[i] = 1
        } else {
            stairs[i] = stairs[i-1] + stairs[i-2]
        }

    }
    return stairs[n]
}
