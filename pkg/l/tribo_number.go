package l

func tribonacci(n int) int {
    stairs := make([]int, n + 1)

    for i := 0; i <= n; i++ {
        if i == 0 {
            stairs[i] = 0
        } else if i == 1 {
            stairs[i] = 1
        } else if i == 2 {
            stairs[i] = 1
        } else {
            stairs[i] = stairs[i-1] + stairs[i-2] + stairs[i-3]
        }

    }
    return stairs[n]
}
