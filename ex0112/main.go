/*
description :
  ex0112 serves a random Lissajous in the browser

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/24
*/
package main

import "image"
import "image/color"
import "image/gif"
import "io"
import "log"
import "math"
import "math/rand"
import "net/http"
import "strconv"
import "time"

var palette = []color.Color{color.White, color.Black}

func lissajous(out io.Writer, cycles int, size int) {
	// const cycles = 5   // number of complete x oscillator revolutions
	const res = 0.001 // angular resolution
	// const size = 100   // image canvas covers [-size..+size]
	const nframes = 64 // number of animation frames
	const delay = 8    // delay between frames in 10ms units

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	freq := r.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5), 1)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func lissajousHandler(rw http.ResponseWriter, req *http.Request) {
	var cycles int = 5 // default number of complete x oscillator revolutions
	var size int = 100 // default image canvas covers [-size..+size]

	err := req.ParseForm()
	if err != nil {
		log.Print(err)
	}

	cyclesForm, ok := req.Form["cycles"]
	if ok {
		cyclesInt, err := strconv.Atoi(cyclesForm[0])
		if err != nil {
			log.Print(err)
		} else {
			cycles = cyclesInt
		}
	}

	sizeForm, ok := req.Form["size"]
	if ok {
		sizeInt, err := strconv.Atoi(sizeForm[0])
		if err != nil {
			log.Print(err)
		} else {
			size = sizeInt
		}
	}
	lissajous(rw, cycles, size)
}

func main() {
	http.HandleFunc("/", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
