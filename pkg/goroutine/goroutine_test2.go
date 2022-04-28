package main

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

    go func() {
        wg.Add(1)
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
