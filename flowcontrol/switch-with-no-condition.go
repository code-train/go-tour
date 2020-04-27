package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.Now()
	// 没有条件的 switch 可以在 case 直接使用条件
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening")
	}

}
