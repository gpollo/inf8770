package edges

import (
	"decompose/layer"
	"image"
	"image/color"
)

type Edges [][]bool

func NewEdges(sizeX, sizeY uint) Edges {
	data := make([][]bool, sizeY)
	for j := uint(0); j < sizeY; j++ {
		data[j] = make([]bool, sizeX)
	}

	return data
}

func FromLayer(l layer.Layer, crop float64) Edges {
	sizeX, sizeY := l.GetDimensions()

	e := NewEdges(uint(sizeX), uint(sizeY))
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			e[y][x] = (l[y][x] > crop)
		}
	}

	return e
}

func (e Edges) ExpandRadius(r uint) Edges {
	copy := e.Copy()

	sizeX, sizeY := e.GetDimensions()
	for y := uint(0); y < sizeY; y++ {
		for x := uint(0); x < sizeX; x++ {
			if !e[y][x] {
				copy.DrawCircle(x, y, r)
			}
		}
	}

	return copy
}

func (e Edges) DrawCircle(x, y, r uint) {
	sizeX, sizeY := e.GetDimensions()

	minX := int(x) - int(r)
	if minX < 0 {
		minX = 0
	}

	minY := int(y) - int(r)
	if minY < 0 {
		minY = 0
	}

	maxX := int(x) + int(r)
	if maxX > int(sizeX) {
		maxX = int(sizeX)
	}

	maxY := int(y) + int(r)
	if maxY > int(sizeY) {
		maxY = int(sizeY)
	}

	for j := minY; j < maxY; j++ {
		for i := minX; i < maxX; i++ {
			dX := int(x) - i - 1
			dY := int(y) - j - 1

			if dX*dX+dY*dY < int(r)*int(r) {
				e[j][i] = false
			}
		}
	}
}

func (e Edges) GetDimensions() (uint, uint) {
	if len(e) == 0 {
		panic("Invalid image data")
	}

	return uint(len(e[0])), uint(len(e))
}

func (e Edges) Copy() Edges {
	sizeX, sizeY := e.GetDimensions()

	copy := NewEdges(sizeX, sizeY)
	for y := uint(0); y < sizeY; y++ {
		for x := uint(0); x < sizeX; x++ {
			copy[y][x] = e[y][x]
		}
	}

	return copy
}

func (e Edges) ColorModel() color.Model {
	return color.Gray16Model
}

func (e Edges) Bounds() image.Rectangle {
	sizeX, sizeY := e.GetDimensions()

	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: int(sizeX),
			Y: int(sizeY),
		},
	}
}

func (e Edges) At(x, y int) color.Color {
	if e[y][x] {
		return &color.Gray16{Y: 0xffff}
	} else {
		return &color.Gray16{Y: 0x0000}
	}
}
