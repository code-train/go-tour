package main

import (
	"fmt"
	"math"
)

// Vertex struct
type Vertex struct{
	X, Y float64
}

// Scale func
func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

// Abs func
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main(){
	v := &Vertex{3, 4}
	fmt.Println(v, v.Abs())
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
