package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

//
var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var width int = 600
var height int = 320
var color string = "FFFFFF"
var xyscale float64 = float64(width) / 2 / xyrange
var zscale float64 = float64(height) * 0.4

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var err error
	for k, v := range params {
		switch k {
		case "width":
			width, err = strconv.Atoi(v[0])
			xyscale = float64(width) / 2 / xyrange
		case "height":
			height, err = strconv.Atoi(v[0])
			zscale = float64(height) * 0.4
		case "color":
			color = v[0]
		}
		if err != nil {
			fmt.Fprint(w, "bad query parameters\n")
			return
		}
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w)
}

func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:#%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
