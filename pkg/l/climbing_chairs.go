package l

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
