package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock() //片方のgoroutineが書き換えている間にもう片方のgoroutineが書き換えられないようにする
	c.v[key]++
	c.mux.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mux.Lock() //片方のgoroutineが書き換えている間にもう片方のgoroutineが書き換えられないようにする
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := Counter{v: make(map[string]int)} //2つのgoroutineから1つのmapを書き換えようとするとエラーになりうる
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c.v, c.v["key"])
}
