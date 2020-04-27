package main

import (
	"fmt"
	"math"
)

func main(){
	x, y := 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)

	//var f float64 = math.Sqrt(x*x + y*y)
	//var z int = f
	fmt.Println(x, y, z)
}

