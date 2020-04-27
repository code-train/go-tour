package main

import "fmt"

func main(){
	sum := 1
	// go 中 while 使用 for实现
	for sum < 1000 {
	//	fmt.Println(sum)
		sum += sum
	//	fmt.Println(sum)
	}
	fmt.Println(sum)
}
