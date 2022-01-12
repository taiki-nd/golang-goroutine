package main

import "fmt"

func main() {
	ch := make(chan int, 2) //bufferを指定
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	//ch <- 300
	//fmt.Println(len(ch))bufferが2なので3つめがあるとエラーが出る

	x := <-ch //1つ取り出す
	fmt.Println(x)
	fmt.Println(len(ch)) //chの中には1つだけ
	ch <- 300            //追加できる
	fmt.Println(len(ch))

	close(ch) //chをcloseする必要がある。これなしでfor文を回すとあるはずのない3つめの値を拾いにいってしまう
	for c := range ch {
		fmt.Println(c)
	}
}
