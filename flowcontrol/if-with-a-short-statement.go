package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// if 可在前面定义一个变量， 此变量只能在 if 作用域内使用
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// return v will show undefined: v error message
	return lim
}

func main(){
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

