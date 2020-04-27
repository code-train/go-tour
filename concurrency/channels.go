package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println("set sum:", sum)
	c <- sum // 将和传入c
	
}

func main(){
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	
	x, y := <-c, <-c //从c中取值

	fmt.Println(x, y, x+y)


}