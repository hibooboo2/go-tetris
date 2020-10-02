package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func SetDrawColor(renderer *sdl.Renderer, color string) {
	switch color {
	case "lightblue":
		renderer.SetDrawColor(0x69, 0xff, 0xff, 255)
	case "green":
		renderer.SetDrawColor(0x60, 0xff, 0x02, 255)
	case "orange":
		renderer.SetDrawColor(0xf4, 0xaa, 0x02, 255)
	case "blue":
		renderer.SetDrawColor(0x2c, 0x03, 0xff, 255)
	case "red":
		renderer.SetDrawColor(0xee, 0x0b, 0x02, 255)
	case "purple":
		renderer.SetDrawColor(0x99, 0x00, 0xfe, 255)
	case "yellow":
		renderer.SetDrawColor(0xfc, 0xff, 0x00, 255)
	default:
		renderer.SetDrawColor(120, 20, 40, 255)
		log.Printf("No case for %s color so default used", color)
	}
}
