/*
description :
  ex0304 is a variation on the surface program

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/08/01
*/
package main

import "fmt"
import "math"
import "log"
import "net/http"
import "strconv"

func f(x float64, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func corner(i int, j int, xyrange float64, sin30 float64, cos30 float64, cells int, width int, height int) (float64, float64) {
	var xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
	var zscale = float64(height) * 0.4         // pixels per z unit

	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func surfaceHandler(rw http.ResponseWriter, req *http.Request) {
	var width, height = 600, 320 // canvas size in pixels
	var cells = 100              // number of grid cells
	var xyrange = 30.0           // axis ranges (-xyrange .. +xyrange)
	var angle = math.Pi / 6      // angle of x, y axes (=30°)

	var sin30 = math.Sin(angle) // sin(30°)
	var cos30 = math.Cos(angle) // cos(30°)
	var color string = "black"  // color of the image

	err := req.ParseForm()
	if err != nil {
		log.Print(err)
	}

	widthForm, ok := req.Form["width"]
	if ok {
		widthInt, err := strconv.Atoi(widthForm[0])
		if err != nil {
			log.Print(err)
		} else {
			width = widthInt
		}
	}

	heightForm, ok := req.Form["height"]
	if ok {
		heightInt, err := strconv.Atoi(heightForm[0])
		if err != nil {
			log.Print(err)
		} else {
			height = heightInt
		}
	}

	colorForm, ok := req.Form["color"]
	if ok {
		color = colorForm[0]
	}

	rw.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(rw, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style= 'stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", color, width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, xyrange, sin30, cos30, cells, width, height)
			bx, by := corner(i, j, xyrange, sin30, cos30, cells, width, height)
			cx, cy := corner(i, j+1, xyrange, sin30, cos30, cells, width, height)
			dx, dy := corner(i+1, j+1, xyrange, sin30, cos30, cells, width, height)

			fmt.Fprintf(rw, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintln(rw, "</svg>")
}

func main() {
	http.HandleFunc("/", surfaceHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
