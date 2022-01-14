/*
Q1. 以下のように文字列を連結させて出力したいコードがありますが、

test1!
test1!test2!
test1!test2!test3!
test1!test2!test3!test4!
以下のコードには間違いがあります。上記の出力になるように修正してください。

package main

import "fmt"

func goroutine(s []string, c chan int){
    sum := ""
    for _, v := range s{
        sum += v
        c <- sum
    }
}

func main(){
    words := []string{"test1!", "test2!", "test3!", "test4!"}
    c := make(chan int)
    go goroutine(words, c)
    for w := range c{
        fmt.Println(w)
    }
}
*/

package main

import (
	"fmt"
)

func goroutine(s []string, c chan string) {
	ss := ""
	for _, v := range s {
		ss += v
		c <- ss
	}
	close(c)
}

func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
