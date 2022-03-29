package main

import (
	"fmt"
	"sync"
	"time"
)

func withReturn(m *sync.Map) {
	ticker := time.NewTicker(1 * time.Second) //10s 汇总
	for range ticker.C {
		m.Range(func(key, value interface{}) bool {
			val := value.(*sync.Map)
            fmt.Println("key:", key)
			val.Range(func(id, v interface{}) bool {
            fmt.Println("in key:", id)
            fmt.Println("in value:", v)
				return true
			})
			return true
		})

	}
}
func testSyncMapRange(){
    m := &sync.Map{}
    m1 := &sync.Map{}
    m2 := &sync.Map{}
    m1.Store("1", 1)
    m1.Store("2", 2)
    m2.Store("3", 3)
    m2.Store("4", 4)
    m.Store("a", m1)
    m.Store("b", m2)
    m.Range(func(key, value interface{}) bool {
        val := value.(*sync.Map)
        fmt.Println("key:", key)
        val.Range(func(id, v interface{}) bool {
            fmt.Println("in key:", id)
            fmt.Println("in value:", v)
            return true
        })
        return true
    })
}

func withoutReturn() {
	var cnt = 0
	ticker := time.NewTicker(1 * time.Second) //10s 汇总
	for range ticker.C {
		fmt.Println("test count: ", cnt)
		fmt.Println(time.Now().Unix())
		fmt.Println(time.Now())
		cnt++
	}

}
