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
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgbVector := r3.Vector{X: float64(x) / float64(width), Y: float64(y) / float64(height), Z: 0.2}
			rgbColor := color.RGBA{
				uint8(baseColor * rgbVector.X),
				uint8(baseColor * rgbVector.Y),
				uint8(baseColor * rgbVector.Z), 255}
			im.Set(x, y, rgbColor)
		}
	}
}
