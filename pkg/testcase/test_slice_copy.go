package testcase

import (
    "fmt"
)

func SliceCopyWithoutPtr() {
	var sli []int
	sli = []int{1, 2, 3, 4, 5}
	modifySlice(sli)
	fmt.Println(sli)
}

func modifySlice(sli []int) {
	sli[0] = -1
}


