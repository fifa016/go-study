package testcase

import (
	"fmt"
	"testing"
	"time"
)

func TestClosure() {
	fmt.Println("=========start run=========")
	values := []int64{1, 2, 3}
	// you will see the last element printed for every iteration instead of each value in sequence,
	// because the goroutines will probably not begin executing until after the loop.
	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}
	time.Sleep(1000000)
}

func TestFilterImpl_RelevanceInAdvance(t *testing.T){
	fmt.Println("test case")
}