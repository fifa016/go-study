package l

func PlusOne(digits []int) []int {
    jinwei := 1
    tmp := 0
    for i := len(digits) - 1; i >= 0 && jinwei == 1; i-- {
        tmp = digits[i] + jinwei
        if tmp >= 10 {
            tmp = tmp - 10
            jinwei = 1
        } else {
            jinwei = 0
        }
        digits[i] = tmp
    }
    if jinwei == 1 {
        return append([]int{1}, digits...)
    } else {
        return digits
    }
}
