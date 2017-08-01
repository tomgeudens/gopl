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

const width, height = 600, 320      // canvas size in pixels
const cells = 100                   // number of grid cells
const xyrange = 30.0                // axis ranges (-xyrange .. +xyrange)
const xyscale = width / 2 / xyrange // pixels per x or y unit
const zscale = height * 0.4         // pixels per z unit
const angle = math.Pi / 6           // angle of x, y axes (=30°)

var sin30 = math.Sin(angle) // sin(30°)
var cos30 = math.Cos(angle) // cos(30°)

func f(x float64, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func corner(i int, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func surfaceHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(rw, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style= 'stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Fprintf(rw, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintln(rw, "</svg>")
}

func main() {
	http.HandleFunc("/", surfaceHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
