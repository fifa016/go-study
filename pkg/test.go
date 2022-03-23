package main

import (
	"fmt"
	"go-study/pkg/util"
	"strconv"
)

func main() {
	fmt.Println("=========start run=========")


	a := 1
	//fmt.Printf("%p\n", a)
	//fmt.Printf("%p\n", &a)
	//fmt.Printf("%p\n", *a)
	fmt.Println("asdf","ffff",a)

	//testSliceSwap()
	//testIntSwap()
	//testNilString()
	//testIntArray()
	//testStringJoin()
	//testMap()
}

func testMap() {
	map1 := make(map[int64]string)

	map2 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	map3 := map[int]string{}
	fmt.Println(len(map1))
	fmt.Println(len(map2))
	fmt.Println(len(map3))
	fmt.Printf("map2:%v\n", map2)
	map1[1] = "a"
	map2[2] = "bb"
	map2[4] = "d"
	map3[1] = "a"

	//compile error
	//var map4 map[int]string
	//map4[1] = "a"

	val, exists := map1[2]
	if !exists {
		fmt.Println("key not exists")
	} else {
		fmt.Println("val: " + val)
	}

	map1[1] = ""
	val, exists = map1[1]
	fmt.Printf("exitst: %t\n", exists)
	delete(map2, 2)
	for k, v := range map2 {
		fmt.Println("key:" + strconv.Itoa(k) + " val: " + v)
	}

}

func testForeachDelete() {
	slice := []int{1, 2, 3, 4, 5, 6}
	//会越界
	//for i, _ := range slice {
	//	if slice[i] == 2 {
	//		slice = append(slice[:i], slice[i + 1:]...)
	//	}
	//}

	for i := 0; i < len(slice); i++ {
		if slice[i] == 2 || slice[i] == 3 {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	fmt.Println("len:" + strconv.Itoa(len(slice)))
	fmt.Println("cap:" + strconv.Itoa(cap(slice)))
	for _, v := range slice {
		fmt.Println(v)
	}

}

func testForeachSlice() {
	var slice = []int{1, 2, 3, 4, 5}

	for i, v := range slice {
		fmt.Println("index:" + strconv.Itoa(i) + " value:" + strconv.Itoa(v))
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println("index:" + strconv.Itoa(i) + " value:" + strconv.Itoa(slice[i]))
	}
}
func testSlice() {
	var sli = make([]int, 3)
	fmt.Printf("int[1]:%d\n", sli[1])

	var sli2 = make([]string, 3)
	fmt.Printf("str[1].len:%d\n", len(sli2[1]))

	var sli3 = make([]bool, 3)
	fmt.Printf("sli3[1]:%t\n", sli3[1])

	var sli4 = make([]int, 3, 5)
	fmt.Println(strconv.Itoa(sli4[2]))
	sli4 = append(sli4, 4)
	sli4 = append(sli4, 5)
	sli4 = append(sli4, 6)
	fmt.Println(strconv.Itoa(len(sli4)))
	fmt.Println(strconv.Itoa(sli4[3]))

	var bigSlice = []int{1, 2, 3, 4, 5}
	var smallSlice = bigSlice[1:3]
	fmt.Println(smallSlice[1])
	bigSlice2 := append(bigSlice, 6)
	fmt.Printf("bigSlice addr:%p\n", bigSlice)
	fmt.Printf("bigSlice2 addr:%p\n", bigSlice2)

}

func testIntArray() {
	//var arr = [4]*int{new(int)}
	//fmt.Println(*arr[0])

	var arr1 = []int{1, 2, 3, 4}
	var arr2 = arr1
	arr2[2] = 10
	fmt.Println("arr1[2]:" + strconv.Itoa(arr1[2]))

	var arr3 = [2]string{"abc", "defg"}
	var arr4 = arr3
	arr4[1] = "----"
	fmt.Println(arr3[1])

}
func testStringJoin() {
	fmt.Println("a" + "b")
	var a = "a"
	fmt.Println(&a)
	a = "b"
	fmt.Println(&a)
}

func testSwap() {
	var a = 1
	var b = 2
	fmt.Printf("before swap a:%d, b:%d\n", a, b)
	util.Swap(&a, &b)
	fmt.Printf("after swap a:%d, b:%d\n", a, b)
}

func testIntSwap() {
	fmt.Println("test int swap")
	var a = 1
	var b = 2
	fmt.Printf("before swap a:%d, b:%d\n", a, b)
	//a, b = b, a
	util.FakeSwap(a, b)

	fmt.Printf("after swap a:%d, b:%d\n", a, b)
}

func testSliceSwap() {
	fmt.Println("test slice swap")
	var slice = []int{1, 2, 3}

	fmt.Printf("before swap a:%d, b:%d\n", slice[0], slice[1])
	slice[0], slice[1] = slice[1], slice[0]
	fmt.Printf("before swap a:%d, b:%d\n", slice[0], slice[1])
}

func testNilString() {
	var s string
	fmt.Println(s)
	fmt.Println(len(s))

}
