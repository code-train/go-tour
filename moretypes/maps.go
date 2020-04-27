package main

import "fmt"

//
type Vertex struct{
	Lat, Long float64
}

var m map[string]Vertex

func main(){

	// map 在使用之前必须由 make 来创建 值为 nil 的 map 是空的，并且不能赋值
	fmt.Println(m)

	m = make(map[string]Vertex)
	fmt.Println(m)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}

	fmt.Println(m["Bell Labs"])

	m["Tom Labs"] = Vertex{-40.68433, 74.39967}

	fmt.Println(m)
}