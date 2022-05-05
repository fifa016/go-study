package l

// func main() {
//     nums := []int{3, 2, 7, 1, 4, 6, 8, 5, 10, 9}
//     fmt.Println("before sort")
//     fmt.Println(nums)
//     mergeSort(nums)
//     fmt.Println("after sort")
//     fmt.Println(nums)
// }

func QuickSort(nums []int) {
    doQuickSort(&nums, 0, len(nums)-1)
}

func doQuickSort(nums *[]int, start int, end int) {
    if start < 0 || end >= len(*nums) || start >= end {
        return
    }
    right := end
    left := start
    for left < right {
        for (*nums)[right] >= (*nums)[start] && left < right {
            right--
        }
        for (*nums)[left] <= (*nums)[start] && left < right {
            left++
        }
        if left < right {
            swap(nums, left, right)
        }
    }
    swap(nums, start, left)
    doQuickSort(nums, start, right-1)
    doQuickSort(nums, right+1, end)

}

func swap(nums *[]int, left int, right int) {
    //(*nums)[left] = (*nums)[left] ^ (*nums)[right]
    //(*nums)[right] = (*nums)[left] ^ (*nums)[right]
    //(*nums)[left] = (*nums)[left] ^ (*nums)[right]
    tmp := (*nums)[left]
    (*nums)[left] = (*nums)[right]
    (*nums)[right] = tmp

}
