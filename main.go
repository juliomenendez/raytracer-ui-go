package main

import (
	"fmt"
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
	size := 400
	window, err := wde.NewWindow(size, size)
	if err != nil {
		fmt.Println(err)
		return
	}
	window.SetTitle("RayTracer in Golang")
	window.SetSize(size, size)
	window.Show()

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
