package l

/**
* 遇到找数想二分
* mid * mid <= x 的最大值
*/
func mySqrt(x int) int {

    res := -1
    left := 0
    right := x

    for left <= right {
        mid := left + (right - left) / 2
        if mid * mid <= x {
            res = mid
            left = mid + 1
        }else {
            right = mid - 1
        }
    }

    return res

}