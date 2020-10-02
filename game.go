package main

import (
	"log"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func NewGame(rootX int32, rootY int32) *Game {
	g := Game{
		fallingPiece: RandomPiece(rootX, rootY),
		RootX:        rootX,
		RootY:        rootY,
		level:        1,
	}
	for y := 0; y < 24; y++ {
		row := []int{}
		for x := 0; x < 10; x++ {
			row = append(row, -1)
		}
		g.board = append(g.board, row)
	}
	return &g
}

type Game struct {
	fallingPiece *Piece
	heldPiece    *Piece
	board        [][]int
	RootX        int32
	RootY        int32
	lock         sync.RWMutex
	level        int
}

func (g *Game) Draw(renderer *sdl.Renderer) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	r, gc, b, a, _ := renderer.GetDrawColor()
	renderer.SetDrawColor(120, 50, 120, 255)
	renderer.DrawRect(&sdl.Rect{g.RootX, g.RootY, 10 * 20, 24 * 20})

	for y := range g.board {
		for x, val := range g.board[y] {
			if val > -1 {
				SetDrawColor(renderer, colorsToPiece[val])
				renderer.FillRect(&sdl.Rect{g.RootX + (int32(x) * 20), g.RootY + (int32(y) * 20), 20, 20})
				SetDrawColor(renderer, "black")
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

func (g *Game) Start() {
	go func() {
		for {
			time.Sleep(time.Second / time.Duration(g.level+4))
			g.Update()
		}
	}()
}

func (g *Game) Update() {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.fallingPiece.positionY++
	if !g.canMoveTo() {
		g.fallingPiece.positionY--
		for y := range g.fallingPiece.blocks[g.fallingPiece.rotation] {
			for x, box := range g.fallingPiece.blocks[g.fallingPiece.rotation][y] {
				if box > 0 {
					yloc := int(g.fallingPiece.positionY) + y
					xloc := int(g.fallingPiece.positionX) + x
					g.board[yloc][xloc] = g.fallingPiece.blockType
				}
			}
		}
		clearedLines := []int{}
		for y := range g.board {
			allFilled := true
			for x := range g.board[y] {
				if g.board[y][x] < 0 {
					allFilled = false
				}
			}
			if allFilled {
				clearedLines = append(clearedLines, y)
			}
		}
		for _, line := range clearedLines {
			log.Println("Cleared:", line)
			g.board = append(g.board[:line], g.board[line+1:]...)
			g.board = append([][]int{{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}, g.board...)
		}

		g.fallingPiece = RandomPiece(g.RootX, g.RootY)
	}

}

func (g *Game) Rotate() {
	g.lock.Lock()
	g.fallingPiece.Rotate(true)
	if !g.canMoveTo() {
		g.fallingPiece.Rotate(false)
	}
	g.lock.Unlock()
}

func (g *Game) MoveLeft() {
	g.lock.Lock()
	g.fallingPiece.positionX--
	if !g.canMoveTo() {
		g.fallingPiece.positionX++
	}
	g.lock.Unlock()
}

func (g *Game) MoveRight() {
	g.lock.Lock()
	g.fallingPiece.positionX++
	if !g.canMoveTo() {
		g.fallingPiece.positionX--
	}
	g.lock.Unlock()
}
func (g *Game) MoveDown() {
	g.lock.Lock()
	g.fallingPiece.positionY++
	if !g.canMoveTo() {
		g.fallingPiece.positionY--
	}
	g.lock.Unlock()
}

func (g *Game) HoldPiece() {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.heldPiece == nil {
		g.heldPiece = g.fallingPiece
		g.fallingPiece = RandomPiece(g.RootX, g.RootY)
	} else {
		tmp := g.heldPiece
		g.heldPiece = g.fallingPiece
		g.fallingPiece = tmp
	}
	g.heldPiece.positionX = 3
	g.heldPiece.positionY = 0
}

func (g *Game) canMoveTo() bool {
	for y := range g.fallingPiece.blocks[g.fallingPiece.rotation] {
		for x, box := range g.fallingPiece.blocks[g.fallingPiece.rotation][y] {
			if box > 0 {
				yloc := int(g.fallingPiece.positionY) + y
				xloc := int(g.fallingPiece.positionX) + x
				if xloc < 0 || xloc > 9 {
					return false
				}
				if yloc >= 24 {
					return false
				} else {
					boardCell := g.board[yloc][xloc]
					if boardCell >= 0 {
						return false
					}
					// log.Println(y, x, boardCell)
				}
			}
		}
	}
	return true
}
