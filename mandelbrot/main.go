/*
description :
  mandelbrot emits a PNG image of the Mandelbrot fractal

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/08/28
*/
package main

import "image"
import "image/color"
import "image/png"
import "math/cmplx"
import "os"

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const constrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - constrast*n}
		}
	}

	return color.Black
}

func main() {
	const xmin, ymin, xmax, ymax = -2, -2, +2, +2
	const width, height = 1024, 1024

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			// image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
