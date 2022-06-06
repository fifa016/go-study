/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-04-29 22:23:51
 */
package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	testGoroutineReturn()
	fmt.Println("after test return")
}

func testGoroutineReturn() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		fmt.Println("goroutine begin")
		time.Sleep(time.Duration(5) * time.Second)
		fmt.Println("goroutine end")
		defer wg.Done()
	}()

	if code := 1; code == 1 {
		return
	}
	wg.Wait()
	fmt.Println("merge")

}
