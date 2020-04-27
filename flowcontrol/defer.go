package main

import "fmt"

func main(){
	// 延迟调用
	defer fmt.Println("world")

	fmt.Println("hello")
}
