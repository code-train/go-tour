package main

import "fmt"

func main(){
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// 当函数返回时，会按照后进先出的顺序调用被延迟的函数
		// https://learnku.com/docs/go-blog/defer-panic-and-recover/6585
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
