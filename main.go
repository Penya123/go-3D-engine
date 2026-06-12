package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Point2D struct {
	X int
	Y int
}

type Triangle2D struct {
	V1    Point2D
	V2    Point2D
	V3    Point2D
	Color color.Color
}

type Game struct {
	MyTriangle Triangle2D
	DirectionX int
	DirectionY int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255})
	g.MyTriangle.drawTriangle(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (t *Triangle2D) drawTriangle(screen *ebiten.Image) {
	drawLine(screen, t.V1.X, t.V1.Y, t.V2.X, t.V2.Y, t.Color)
	drawLine(screen, t.V2.X, t.V2.Y, t.V3.X, t.V3.Y, t.Color)
	drawLine(screen, t.V3.X, t.V3.Y, t.V1.X, t.V1.Y, t.Color)
}

func drawLine(screen *ebiten.Image, x1, y1, x2, y2 int, color color.Color) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)

	sx, sy := 1, 1
	if x1 > x2 {
		sx = -1
	}
	if y1 > y2 {
		sy = -1
	}

	err := dx - dy

	for {
		screen.Set(x1, y1, color)

		if x1 == x2 && y1 == y2 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	ebiten.SetWindowSize(screenWidth*3, screenHeight*3)
	ebiten.SetWindowTitle("3D engine in Go - phase 1")

	game := &Game{
		MyTriangle: Triangle2D{
			V1:    Point2D{160, 50},
			V2:    Point2D{100, 180},
			V3:    Point2D{220, 180},
			Color: color.RGBA{255, 0, 255, 255},
		},
		DirectionX: 2,
		DirectionY: 2,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
