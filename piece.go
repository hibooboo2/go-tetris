package main

import (
	"log"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

func RandomPiece() Piece {
	p := Piece{}
	p.blocks = allPieces[rand.Intn(7)]
	p.color = "lightblue"
	return p
}

type Piece struct {
	blocks    [][][]int
	color     string
	rotation  int
	positionX int32
	positionY int32
}

func (p *Piece) Rotate() {
	switch len(p.blocks) {
	case 4:
		switch p.rotation {
		case 0, 1, 2:
			log.Println("Rotated", p.rotation)
			p.rotation++
		case 3:
			p.rotation = 0
			log.Println("Rotated to 0")
		default:
			log.Fatalf("Invalid rotation state")
		}
	default:
		log.Fatalf("Invalid blocks passed to rotate %v", p.blocks)
	}
}

func (p Piece) Draw(renderer *sdl.Renderer) {
	for y := range p.blocks[p.rotation] {
		for x, box := range p.blocks[p.rotation][y] {
			if box > 0 {
				r, g, b, a, _ := renderer.GetDrawColor()
				renderer.SetDrawColor(colors[p.color][0], colors[p.color][1], colors[p.color][2], colors[p.color][3])
				renderer.FillRect(&sdl.Rect{p.positionX + int32(x*20), p.positionY + int32(y*20), 20, 20})
				renderer.SetDrawColor(0, 0, 0, 0)
				renderer.DrawRect(&sdl.Rect{p.positionX + int32(x*20), p.positionY + int32(y*20), 20, 20})
				renderer.SetDrawColor(r, g, b, a)
			} else {
				renderer.SetDrawColor(colors[p.color][3], colors[p.color][2], colors[p.color][1], colors[p.color][0])
				renderer.FillRect(&sdl.Rect{p.positionX + int32(x*20), p.positionY + int32(y*20), 20, 20})
				renderer.SetDrawColor(0, 0, 0, 0)
				renderer.DrawRect(&sdl.Rect{p.positionX + int32(x*20), p.positionY + int32(y*20), 20, 20})
			}

		}
	}
}

var allPieces = [][][][]int{
	Line_piece,
	BlueL_piece,
	OrangeL_piece,
	Square_piece,
	GreenS_piece,
	PurpleT_piece,
	RedS_piece,
}

var Line_piece = [][][]int{
	{
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
	},
	{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
	},
	{
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
	},
}

var BlueL_piece = [][][]int{
	{
		{1, 0, 0},
		{1, 1, 1},
		{0, 0, 0},
	},
	{
		{0, 1, 1},
		{0, 1, 0},
		{0, 1, 0},
	},
	{
		{0, 0, 0},
		{1, 1, 1},
		{0, 0, 1},
	},
	{
		{0, 1, 0},
		{0, 1, 0},
		{1, 1, 0},
	},
}

var OrangeL_piece = [][][]int{
	{
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	},
	{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 1},
	},
	{
		{0, 0, 0},
		{1, 1, 1},
		{1, 0, 0},
	},
	{
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	},
}

var Square_piece = [][][]int{
	{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	},
	{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	},
	{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	},
	{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	},
}
var GreenS_piece = [][][]int{
	{
		{0, 1, 1},
		{1, 1, 0},
		{0, 0, 0},
	},
	{
		{0, 1, 0},
		{0, 1, 1},
		{0, 0, 1},
	},
	{
		{0, 0, 0},
		{0, 1, 1},
		{1, 1, 0},
	},
	{
		{1, 0, 0},
		{1, 1, 0},
		{0, 1, 0},
	},
}

var PurpleT_piece = [][][]int{
	{
		{0, 1, 0},
		{1, 1, 1},
		{0, 0, 0},
	},
	{
		{0, 1, 0},
		{0, 1, 1},
		{0, 1, 0},
	},
	{
		{0, 0, 0},
		{1, 1, 1},
		{0, 1, 0},
	},
	{
		{0, 1, 0},
		{1, 1, 0},
		{0, 1, 0},
	},
}

var RedS_piece = [][][]int{
	{
		{1, 1, 0},
		{0, 1, 1},
		{0, 0, 0},
	},
	{
		{0, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	},
	{
		{0, 0, 0},
		{1, 1, 0},
		{0, 1, 1},
	},
	{
		{0, 1, 0},
		{1, 1, 0},
		{1, 0, 0},
	},
}

var colors = map[string][]uint8{
	"lightblue": []uint8{0, 100, 125, 255},
}
