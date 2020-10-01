package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	running := true

	renderer, cancel := GetRenderer(600, 800)
	defer cancel()
	p := RandomPiece()
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.KeyboardEvent:
				switch e.Type {
				case 768: //Key Press
					switch e.Keysym.Sym {
					case 97, 1073741904: //Left
						p.Rotate()
					case 100, 1073741903: //Right
						p.Rotate()
					default:
						log.Printf("Key was: %v", e.Keysym.Sym)
					}
				case 769: //Key release
				default:
					log.Printf("Type was: %v", e.Type)
				}
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		renderer.Clear()
		renderer.SetDrawColor(255, 0, 0, 0)
		p.Draw(renderer)
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
