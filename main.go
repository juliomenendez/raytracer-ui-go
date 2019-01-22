package main

import (
	"fmt"
	"image/color"
	"runtime"

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
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r := toRGBColorComponent(float64(x) / float64(width))
			g := toRGBColorComponent(float64(y) / float64(height))
			b := toRGBColorComponent(0.2)
			im.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
}

func toRGBColorComponent(value float64) uint8 {
	return uint8(255.99 * value)
}
