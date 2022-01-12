package main

import (
	"fmt"
	"sync"
)

/*
func goroutine(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go goroutine("world")
	normal("hello")
}
*/

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		//time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done() //処理が終わったら
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1) //一個の並列処理があることを伝える
	go goroutine("world", &wg)
	normal("hello")
	wg.Wait() //wgがdoneするのを待つ
}
