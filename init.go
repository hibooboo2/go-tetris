package main

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	rect := sdl.Rect{0, 0, 100, 100}

	running := true

	renderer, cancel := GetRenderer(600, 800)
	defer cancel()
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		renderer.Clear()
		renderer.SetDrawColor(255, 0, 0, 0)
		renderer.DrawRect(&rect)

		renderer.DrawLine(0, 0, 800, 800)
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Present()
	}

}

func GetRenderer(h, w int32) (*sdl.Renderer, func()) {
	runtime.LockOSThread()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	window, r, err := sdl.CreateWindowAndRenderer(w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	window.SetResizable(true)
	window.SetBordered(true)
	window.SetTitle("Go Tetris")
	// window.SetGrab(true)
	// window.SetWindowOpacity(0.4)

	// go eventLoop()
	return r,
		func() {
			window.Destroy()
			sdl.Quit()
			ttf.Quit()
			runtime.UnlockOSThread()
		}

}
