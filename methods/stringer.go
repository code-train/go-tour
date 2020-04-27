package main

import "fmt"

// type Stringer struct {
//     String() string
// }

// Person struct
type Person struct{
	Name string
	Age int
}


func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main(){
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	// fmt内部调用了 Stringer 接口中的 String 方法
	fmt.Println(a, z)
}