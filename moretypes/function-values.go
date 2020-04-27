package main

import(
	"fmt"
	"math"
)

func main(){
	// 将函数作为值
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

}