package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //sumをchannelに送信
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //sumをchannelに送信
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine1(s, c) //呼び出す
	go goroutine2(s, c)
	x := <-c //受け取る,自身するまでここでコードがストップするのでsyncWaitを使わなくて良い
	fmt.Println(x)
	y := <-c //受け取る,自身するまでここでコードがストップするのでsyncWaitを使わなくて良い
	fmt.Println(y)
}
