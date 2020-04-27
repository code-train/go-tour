package main

import "fmt"

// 变量生命语法的由来
// https://learnku.com/docs/go-blog/gos-declaration-syntax/6587

func add(x int, y int) int{
	return x + y
}

func main(){
	fmt.Println(add(42,13))
}
