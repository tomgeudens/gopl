/*
description :
  lissajous generates GIF animations of random Lissajous figures

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/23
*/
package main

import "image"
import "image/color"
import "image/gif"
import "io"
import "math"
import "math/rand"
import "time"
import "os"

var palette = []color.Color{color.White, color.Black}

const whiteIndex = 0 // first color in palette
const blackIndex = 1 // next color in palette

func lissajous(out io.Writer) {
	const cycles = 5   // number of complete x oscillator revolutions
	const res = 0.001  // angular resolution
	const size = 100   // image canvas covers [-size..+size]
	const nframes = 64 // number of animation frames
	const delay = 8    // delay between frames in 10ms units

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // seed

	freq := r.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
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
