package layer

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type Layer [][]float64

func NewLayer(sizeX, sizeY int) Layer {
	data := make([][]float64, sizeY)
	for j := 0; j < sizeY; j++ {
		data[j] = make([]float64, sizeX)
	}

	return data
}

func FromImageBW(i image.Image) Layer {
	sizeX := i.Bounds().Max.X - i.Bounds().Min.X
	sizeY := i.Bounds().Max.Y - i.Bounds().Min.Y

	bw := NewLayer(sizeX, sizeY)

	for y := i.Bounds().Min.Y; y < i.Bounds().Max.Y; y++ {
		for x := i.Bounds().Min.X; x < i.Bounds().Max.X; x++ {
			r, g, b, _ := i.At(x, y).RGBA()
			bw[y][x] = 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
		}
	}

	return bw
}

func (l Layer) GetDimensions() (int, int) {
	if len(l) == 0 {
		panic("Invalid image data")
	}

	return len(l[0]), len(l)
}

func (l Layer) Copy() Layer {
	sizeX, sizeY := l.GetDimensions()

	copy := NewLayer(sizeX, sizeY)
	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			copy[j][i] = l[j][i]
		}
	}

	return copy
}

func (l Layer) String() string {
	sizeX, sizeY := l.GetDimensions()

	str := ""
	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			str += fmt.Sprintf("%8.2f ", l[j][i])
		}
		str += "\n"
	}

	return str
}

func (l Layer) ColorModel() color.Model {
	return color.Gray16Model
}

func (l Layer) Bounds() image.Rectangle {
	sizeX, sizeY := l.GetDimensions()

	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: sizeX,
			Y: sizeY,
		},
	}
}

func (l Layer) At(x, y int) color.Color {
	v := float64(l[y][x])
	v = math.Min(float64(0xffff), v)
	v = math.Max(float64(0x0000), v)

	return &color.Gray16{Y: uint16(v)}
}
