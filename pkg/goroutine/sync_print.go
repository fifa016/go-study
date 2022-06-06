/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-06 23:34:52
 */
package goroutine

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var state int = 0
var wg sync.WaitGroup

func TestPrint() {

	wg.Add(4)
	go printNum(0)
	go printNum(1)
	go printNum(2)
	go printNum(3)

	wg.Wait()
}

func printNum(num int) {

	lock.Lock()
	if state == num {
		fmt.Println(num)
		state++
	}

	defer lock.Unlock()
	defer wg.Done()
}
