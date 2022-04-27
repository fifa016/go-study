package main
import (
    "log"
    "runtime"
    "time"
)

func traceMemStats() {
    var ms runtime.MemStats
    runtime.ReadMemStats(&ms)
    log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func f() {
    container := make([]int, 8)
    log.Println("> loop.")
    for i := 0; i < 32*1000*1000; i++ {
        container = append(container, i)
        if i == 16*1000*1000 {
            traceMemStats()

        }
    }
    log.Println("< loop.")
}

// console下执行：go build -o snippet_mem && GODEBUG='gctrace=1' ./snippet_mem
func main() {
    log.Println("start.")
    traceMemStats()
    f()
    traceMemStats()

    log.Println("force gc.")
    runtime.GC()

    log.Println("done.")
    traceMemStats()

    go func() {
        for {
            traceMemStats()
            time.Sleep(10 * time.Second)
        }
    }()

    time.Sleep(1 * time.Hour)
}