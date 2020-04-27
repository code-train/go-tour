package main

import "fmt"

func main() {
	// slice 是另外一种类型，基于数组来提供能强大的功能和便利
	p := []int{2, 3, 5, 7, 11, 13}

	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

}
