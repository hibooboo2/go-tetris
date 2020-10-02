package main

import (
	"log"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

func RandomPiece(rootX int32, rootY int32) *Piece {
	p := Piece{}
	p.positionX = 3
	p.rootX = rootX
	p.rootY = rootY
	p.blockType = rand.Intn(7)
	p.blocks = allPieces[p.blockType]
	return &p
}

type Piece struct {
	blocks    [][][]int
	blockType int
	rotation  int
	rootX     int32
	rootY     int32
	positionX int32
	positionY int32
}

func (p *Piece) Rotate(normal bool) {
	switch len(p.blocks) {
	case 4:
		if normal {
			p.rotation += 1
			if p.rotation > 3 {
				p.rotation = 0
			}
		} else {
			p.rotation -= 1
			if p.rotation < 0 {
				p.rotation = 3
			}
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
				SetDrawColor(renderer, colorsToPiece[p.blockType])
				renderer.FillRect(&sdl.Rect{(p.positionX+int32(x))*20 + p.rootX, (p.positionY+int32(y))*20 + p.rootY, 20, 20})
				renderer.SetDrawColor(0, 0, 0, 0)
				renderer.DrawRect(&sdl.Rect{(p.positionX+int32(x))*20 + p.rootX, (p.positionY+int32(y))*20 + p.rootY, 20, 20})
				renderer.SetDrawColor(r, g, b, a)
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

var colorsToPiece = []string{
	"lightblue",
	"blue",
	"orange",
	"yellow",
	"green",
	"purple",
	"red",
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
