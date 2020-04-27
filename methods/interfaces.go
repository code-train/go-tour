package main

import (
	"fmt"
	"math"
)

// Abser interface 
type Abser interface{
	Abs() float64
}

// MyFloat float64
type MyFloat float64

// Abs func
func (f MyFloat) Abs() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex struct
type Vertex struct{
	X, Y float64
}

// Abs func
func (v *Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  //a MyFloat 实现了 Abser
	fmt.Println(a.Abs())

	// 此处使用指针指向 *Vertex
	a = &v //a *Vertex 实现了 Abser

	// 此处 v 是一个 Vertex(不是 *Vertex)
	// 所以没有实现 Abser
	a = v
	
	fmt.Println(a.Abs())

}