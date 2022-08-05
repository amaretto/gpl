package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	newImg := supersampling(img)
	png.Encode(os.Stdout, newImg)
}

func supersampling(img *image.RGBA) *image.RGBA {
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < img.Rect.Max.Y-1; py++ {
		for px := 0; px < img.Rect.Max.X-1; px++ {
			newImg.Set(px, py, avg4Colors(img.RGBAAt(px, py), img.RGBAAt(px+1, py), img.RGBAAt(px+1, py+1), img.RGBAAt(px, py+1)))
		}
	}
	return newImg
}

// a b
// c d
func avg4Colors(a, b, c, d color.RGBA) color.RGBA {
	n := color.RGBA{(a.R + b.R + c.R + d.R) / 4, (a.G + b.G + c.G + d.G) / 4,
		(a.B + b.B + c.B + d.B) / 4, (a.A + b.A + c.A + d.A) / 4}
	return n
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := 255 - contrast*n
			g := 0 + contrast*n
			b := 0 + contrast*n
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}
