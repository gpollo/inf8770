package edges

import (
	"decompose/helper"
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

func From6Edges(e1, e2, e3, e4, e5, e6 Edges) Edges {
	sizeX1, sizeY1 := e1.GetDimensions()
	sizeX2, sizeY2 := e1.GetDimensions()
	sizeX3, sizeY3 := e1.GetDimensions()
	sizeX4, sizeY4 := e1.GetDimensions()
	sizeX5, sizeY5 := e1.GetDimensions()
	sizeX6, sizeY6 := e1.GetDimensions()

	if !helper.AreAllUIntEquals([]uint{sizeX1, sizeX2, sizeX3, sizeX4, sizeX5, sizeX6}) {
		panic("X dimensions of all edge layers must be equal")
	}

	if !helper.AreAllUIntEquals([]uint{sizeY1, sizeY2, sizeY3, sizeY4, sizeY5, sizeY6}) {
		panic("X dimensions of all edge layers must be equal")
	}

	sizeX := sizeX1
	sizeY := sizeY1
	e := NewEdges(2*sizeX, 3*sizeY)

	for y := uint(0); y < sizeY; y++ {
		for x := uint(0); x < sizeX; x++ {
			e[y][x] = e1[y][x]
		}
	}

	for y := uint(0); y < sizeY; y++ {
		for x := uint(0); x < sizeX; x++ {
			e[y][x+sizeX] = e2[y][x]
		}
	}

	for y := uint(0); y < sizeY; y++ {
		for x := uint(0); x < sizeX; x++ {
			e[y+sizeY][x] = e3[y][x]
		}
	}

	for y := uint(0); y < sizeY; y++ {
		for x := uint(0); x < sizeX; x++ {
			e[y+sizeY][x+sizeX] = e4[y][x]
		}
	}

	for y := uint(0); y < sizeY; y++ {
		for x := uint(0); x < sizeX; x++ {
			e[y+2*sizeY][x] = e5[y][x]
		}
	}

	for y := uint(0); y < sizeY; y++ {
		for x := uint(0); x < sizeX; x++ {
			e[y+2*sizeY][x+sizeX] = e6[y][x]
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

func (e Edges) Contains(o Edges) Edges {
	copy := o.Copy()

	sizeX, sizeY := o.GetDimensions()
	for y := uint(0); y < sizeY; y++ {
		for x := uint(0); x < sizeX; x++ {
			if !o[y][x] && !e[y][x] {
				copy[y][x] = true
			}
		}
	}

	return copy
}

func (e Edges) Count() uint {
	count := uint(0)
	sizeX, sizeY := e.GetDimensions()
	for y := uint(0); y < sizeY; y++ {
		for x := uint(0); x < sizeX; x++ {
			if !e[y][x] {
				count += 1
			}
		}
	}

	return count
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
