package main

import(
	"fmt"
	"math"
)

// Vertex struct
type Vertex struct{
	X, Y float64
}

// Abs func
func (v *Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	v := &Vertex{3,4}
	fmt.Println(v.Abs())

}