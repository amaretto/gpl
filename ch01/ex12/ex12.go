package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}
var cycles int = 5
var res float64 = 0.001
var size int = 100
var nframes int = 64
var delay int = 8

const (
	whiteIndex = 0
	blackIndex = 1
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
		case "cycles":
			cycles, err = strconv.Atoi(v[0])
		case "res":
			res, err = strconv.ParseFloat(v[0], 64)
		case "size":
			size, err = strconv.Atoi(v[0])
		case "nframes":
			nframes, err = strconv.Atoi(v[0])
		case "delay":
			delay, err = strconv.Atoi(v[0])
		}
		if err != nil {
			fmt.Fprint(w, "bad query parameters\n")
		}
	}
	lissajous(w)
}

func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
