/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-06 23:38:37
 */
package main

import (
	"fmt"
	"go-study/pkg/goroutine"
)


func main() {
	fmt.Println("start")
	goroutine.TestPrint()
	fmt.Println("end")
}