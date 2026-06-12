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

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255})

	//pixelX := screenWidth / 2
	//xelY := screenHeight / 2

	//et(pixelX, pixelY, color.RGBA{255, 0, 0, 255})
	//screen.Set(100, 100, color.RGBA{255, 0, 0, 255})

	//drawLine(screen, 100, 100, pixelX, pixelY, color.RGBA{255, 0, 0, 255})
	drawLine(screen, 160, 50, 100, 180, color.RGBA{215, 95, 255, 255})
	drawLine(screen, 100, 180, 220, 180, color.RGBA{215, 95, 255, 255})
	drawLine(screen, 220, 180, 160, 50, color.RGBA{215, 95, 255, 255})

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
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

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
