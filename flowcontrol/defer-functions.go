package main

import "fmt"

func deferFunc(i int) (t int){
	fmt.Println("t = ", t)
	fmt.Println(i)
	return 2
}

func main(){
	defer deferFunc(10)
}

