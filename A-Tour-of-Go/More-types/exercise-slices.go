package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	d := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		d_col := make([]uint8,dx) 
		for x := 0; x < dx; x++ {
			//d_col[x] = (uint8(x) + uint8(y)) / 2
			//d_col[x] = uint8(x) * uint8(y)
			d_col[x] = uint8(math.Pow(float64(x), float64(y))) 
		}
		d[y] = d_col
	}
	return d
}

func main() {
	pic.Show(Pic)
}
