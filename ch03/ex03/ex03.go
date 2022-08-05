package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var stdColor = color.RGBA{0x00, 0x00, 0xff, 0xff}
var tgtColor = color.RGBA{0xff, 0x00, 0x00, 0xff}

func main() {

	// search maxHeight
	var z, maxHeight float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z = calcAvgHeight(i, j)
			if z > maxHeight {
				maxHeight = z
			}
		}
	}

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			r, g, b := proportion(calcAvgHeight(i, j), maxHeight)
			//fmt.Printf("#%02X, %02X, %02X\n", r, g, b)
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:#%02X%02X%02X'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
	fmt.Println("</svg>")

}

func proportion(height, maxHeight float64) (r, g, b int) {
	r = int(float64(tgtColor.R-stdColor.R)*height/maxHeight + float64(stdColor.R))
	g = int(float64(tgtColor.G-stdColor.G)*height/maxHeight + float64(stdColor.G))
	b = int(float64(tgtColor.B-stdColor.B)*height/maxHeight + float64(stdColor.B))
	if r < 0 {
		r = 0
	} else if r > 255 {
		r = 255
	}
	if g < 0 {
		g = 0
	} else if g > 255 {
		g = 255
	}
	if b < 0 {
		b = 0
	} else if b > 255 {
		b = 255
	}
	return r, g, b
}

func calcAvgHeight(i, j int) float64 {
	h := (calcHeight(i+1, j) + calcHeight(i, j) +
		calcHeight(i, j+1) + calcHeight(i+1, j+1)) / 4
	return h
}

func calcHeight(i, j int) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	return f(x, y)
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
