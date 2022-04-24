package main

func twoSum(arr []int, target int) []int {
    res := []int{0, 0}

    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i]+arr[j] == target {
                return []int{i, j}
            }
        }
    }

    return res
}
