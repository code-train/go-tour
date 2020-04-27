package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // type is Vertex
	v2 = Vertex{X: 1} // Y:0 被省略
	v3 = Vertex{}     // X:0 and Y:0
	p1 = &Vertex{1, 2}// type is *Vertex
	p2 = &v1
)

func main(){
	fmt.Println(v1, p1, p2, v2, v3)
}
