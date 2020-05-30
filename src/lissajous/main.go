package main

import (
	"image/color"
	"image/gif"
	"io"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //nomber of complete x oscillator revolutions
		res     = 0.001 //angular resolution
		site    = 100   //image canvas cover
		nframes = 64    //number of animation frames
		delay   = 8     //delay between frames
	)
	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase
}
