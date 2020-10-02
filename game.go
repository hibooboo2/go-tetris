package main

import "github.com/veandco/go-sdl2/sdl"

func NewGame(rootX int32, rootY int32) *Game {
	g := Game{
		fallingPiece: RandomPiece(rootX, rootY),
		RootX:        rootX,
		RootY:        rootY,
	}
	for y := 0; y < 24; y++ {
		row := []int{}
		for x := 0; x < 10; x++ {
			row = append(row, 0)
		}
		g.board = append(g.board, row)
	}
	return &g
}

type Game struct {
	fallingPiece *Piece
	board        [][]int
	RootX        int32
	RootY        int32
}

func (g *Game) Draw(renderer *sdl.Renderer) {

	r, gc, b, a, _ := renderer.GetDrawColor()
	renderer.SetDrawColor(120, 50, 120, 255)
	renderer.DrawRect(&sdl.Rect{g.RootX, g.RootY, 10 * 20, 24 * 20})

	for y := range g.board {
		for x, val := range g.board[y] {
			if val > 0 {
				renderer.SetDrawColor(120, 50, 120, 255)
				renderer.DrawRect(&sdl.Rect{g.RootX + (int32(x) * 20), g.RootY + (int32(y) * 20), 20, 20})
			} else {
				renderer.SetDrawColor(120, 50, 10, 255)
				renderer.DrawRect(&sdl.Rect{g.RootX + (int32(x) * 20), g.RootY + (int32(y) * 20), 20, 20})
			}
		}
	}

	g.fallingPiece.Draw(renderer)

	renderer.SetDrawColor(r, gc, b, a)
}
