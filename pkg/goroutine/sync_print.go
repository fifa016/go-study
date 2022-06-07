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

	wg.Wait()
}

func printNum(num int) {

	for i := 0; i < 2; {

		lock.Lock()
		if state %3 == num {
			fmt.Println(num)
			state++
			i++
		}

		lock.Unlock()
	}
	wg.Done()
}
