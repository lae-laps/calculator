package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const edgeStartDeviation = 10

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	emptyImage = ebiten.NewImage(3, 3)
)

func init() {
	emptyImage.Fill(color.White)
}

type Game struct {
	count int
}

func line(x0, y0, x1, y1 float32, clr color.RGBA) ([]ebiten.Vertex, []uint16) {
	const width = 1

	theta := math.Atan2(float64(y1-y0), float64(x1-x0))
	theta += math.Pi / 2
	dx := float32(math.Cos(theta))
	dy := float32(math.Sin(theta))

	r := float32(clr.R) / 0xff
	g := float32(clr.G) / 0xff
	b := float32(clr.B) / 0xff
	a := float32(clr.A) / 0xff

	return []ebiten.Vertex{
		{
			DstX:   x0 - width*dx/2,
			DstY:   y0 - width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x0 + width*dx/2,
			DstY:   y0 + width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1 - width*dx/2,
			DstY:   y1 - width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1 + width*dx/2,
			DstY:   y1 + width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
	}, []uint16{0, 1, 2, 1, 2, 3}
}

func (g *Game) Update() error {
	g.count++
	g.count %= 240
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	src := emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

    startPointX := float32(edgeStartDeviation)
    startPointY := float32(screenHeight - edgeStartDeviation)

	v, i := line(startPointX, startPointY, float32(vec.x), float32(vec.y), color.RGBA{0xff, 0xff, 0xff, 0xff})
	screen.DrawTriangles(v, i, src, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func draw_vector_gui() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Vectorial - GUI")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

