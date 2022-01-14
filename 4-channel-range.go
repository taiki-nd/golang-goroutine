package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum //随時channelに送信
	}
	close(c) //close()しないとエラーが出る。
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s)) //bufferの指定も可能
	go goroutine1(s, c)
	for i := range c { //close()しないとエラーが出る
		fmt.Println(i)
	}
}
