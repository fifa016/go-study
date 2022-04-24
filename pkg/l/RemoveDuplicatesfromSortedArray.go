package main

import "fmt"

func main() {
    nums := []int{0,0,1,1,2,2}
    fmt.Println(removeDuplicates(nums))
    fmt.Println(nums)
}
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return 1
    }
    endIndex := len(nums) - 1
    deleteCnt := 0
    cur := nums[0]
    for i := 1; i <= endIndex; i++ {
        if nums[i] != cur {
            cur = (nums)[i]
        }else {
            deleteCnt ++
            moveForward(nums, i, endIndex)
            endIndex --
            i --
        }
    }
    return len(nums) - deleteCnt
}

func moveForward(nums []int, toIndex , endIndex int) {
    for i := toIndex ;i < endIndex; i ++ {
        nums[i] = nums[i + 1]
    }
}

