package main

import "fmt"

func main(){
	
	var i, j int = 1, 2
	// 短声明变量代替 var, 但是此结构不能使用在函数外
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}


