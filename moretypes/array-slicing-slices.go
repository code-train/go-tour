package main

import "fmt"

func main(){
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)
	// 表示从 1 到 4 - 1 的元素
	fmt.Println("p[1:4] ==", p[1:4])

	// 省略开始的下标代表从0开始
	fmt.Println("p[:3]", p[:3])
	// 省略结束的下标代表从 len(p) 结束
	fmt.Println("p[4:] ==", p[4:])

}
