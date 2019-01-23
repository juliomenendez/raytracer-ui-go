package main

import (
	"fmt"
	"image/color"
	"runtime"

	"github.com/golang/geo/r3"

	wde "github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
)

func main() {
	fmt.Println("Ray tracer started")

	go runUI()
	wde.Run()

	fmt.Println("Ray tracer ended")
}

func runUI() {
	width := 600
	height := 400
	window, err := wde.NewWindow(width, height)
	if err != nil {
		fmt.Println(err)
		return
	}
	window.SetTitle("RayTracer in Golang")
	window.SetSize(width, height)
	window.Show()

	draw(window.Screen(), width, height)
	window.FlushImage()

	events := window.EventChan()

	go func() {
	loop:
		for ei := range events {
			runtime.Gosched()
			switch ei.(type) {
			case wde.CloseEvent:
				window.Close()
				wde.Stop()
				break loop
			}
		}
	}()
}

func draw(im wde.Image, width, height int) {
	bounds := im.Bounds()
	baseColor := 255.99
	lowerLeftCorner := r3.Vector{X: -2.0, Y: -1.0, Z: -1.0}
	horizontal := r3.Vector{X: 4.0, Y: 0.0, Z: 0.0}
	vertical := r3.Vector{X: 0.0, Y: 2.0, Z: 0.0}
	origin := r3.Vector{X: 0.0, Y: 0.0, Z: 0.0}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			xCoords := float64(x) / float64(width)
			yCoords := float64(y) / float64(height)
			ray := Ray{origin: origin, direction: lowerLeftCorner.Add(horizontal.Mul(xCoords)).Add(vertical.Mul(yCoords))}
			rgbColor := getColor(ray)
			im.Set(x, y, color.RGBA{
				uint8(baseColor * rgbColor.X),
				uint8(baseColor * rgbColor.Y),
				uint8(baseColor * rgbColor.Z), 255})
		}
	}
}

func getColor(ray Ray) r3.Vector {
	unitDirection := ray.direction.Normalize()
	t := 0.5 * (unitDirection.Y + 1.0)
	return r3.Vector{X: 1.0, Y: 1.0, Z: 1.0}.Mul(1.0 - t).Add(r3.Vector{X: 0.5, Y: 0.7, Z: 1.0}.Mul(t))
}
