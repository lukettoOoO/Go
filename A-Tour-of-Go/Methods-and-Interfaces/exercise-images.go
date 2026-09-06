package main

import (
	"golang.org/x/tour/pic"
	"math"
	"image"
	"image/color"
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

type Image struct{
	width, height int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
