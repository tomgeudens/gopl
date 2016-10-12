/*
description :
  ex0106 is a variation on the lissajous program

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2016/10/12
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White,
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}}

const numberOfColors = 3

func lissajous(out io.Writer) {
	const cycles = 5   // number of complete x oscillator revolutions
	const res = 0.001  // angular resolution
	const size = 100   // image canvas covers [-size..+size]
	const nframes = 64 // number of animation frames
	const delay = 8    // delay between frames in 10ms units

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	freq := r1.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			index := uint8(r1.Int() % numberOfColors)
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				index)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func main() {
	lissajous(os.Stdout)
}
