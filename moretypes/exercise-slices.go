package main

import (
	"golang.org/x/tour/pic"
	"math"
)

// import 需要开启 go module 可以使用 goproxy.cn 代理，速度很快
// go mod init project_name 后即可正常使用

//
func Pic(dx, dy int) [][]uint8 {
	
	a := make([][]uint8, dy)
	for x := range a {
		// 里层slice
		b := make([]uint8, dx)
		for y := range b {
			
			//b[y] = uint8((x+y)/2)
			b[y] = uint8(x*y)
			//b[y] = uint8(math.Pow(float64(x), float64(y)))
		}
		a[x] = b
	}
	return a
}

func main() {
	pic.Show(Pic)
}
