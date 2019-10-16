package data

import (
	"fmt"
	"image"
)

type Layer [][]float32

func NewLayer(sizeX, sizeY int) Layer {
	data := make([][]float32, sizeY)
	for j := 0; j < sizeY; j++ {
		data[j] = make([]float32, sizeX)
	}

	return data
}

func GetLayers(image image.Image) (Layer, Layer, Layer, uint, uint) {
	sizeX := image.Bounds().Max.X - image.Bounds().Min.X
	sizeY := image.Bounds().Max.Y - image.Bounds().Min.Y

	r := NewLayer(sizeX, sizeY)
	g := NewLayer(sizeX, sizeY)
	b := NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			rv, gv, bv, _ := image.At(i, j).RGBA()
			r[j][i] = float32(rv)
			g[j][i] = float32(gv)
			b[j][i] = float32(bv)
		}
	}

	return r, g, b, uint(sizeX), uint(sizeY)
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

func (l Layer) Times(factor float32) {
	sizeX, sizeY := l.GetDimensions()

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			l[j][i] = factor * l[j][i]
		}
	}
}

func (l Layer) ScaleInteger(scaleX, scaleY int) Layer {
	sizeX, sizeY := l.GetDimensions()

	scaled := NewLayer(sizeX*scaleX, sizeY*scaleY)
	for j := 0; j < sizeY*scaleY; j++ {
		for i := 0; i < sizeX*scaleX; i++ {
			scaled[j][i] = l[j/scaleY][i/scaleX]
		}
	}

	return scaled
}

func (l Layer) ToProtobuf() ImageData {
	sizeX, sizeY := l.GetDimensions()

	rows := make([]*ImageRow, sizeY)
	for j := 0; j < sizeY; j++ {
		row := ImageRow{Values: make([]float32, sizeX)}
		for i := 0; i < sizeX; i++ {
			row.Values[i] = l[j][i]
		}
		rows[j] = &row
	}

	return ImageData{Rows: rows}
}

func (l *Layer) FromProtobuf(data ImageData) {
	sizeX := len(data.Rows[0].Values)
	sizeY := len(data.Rows)

	(*l) = NewLayer(sizeX, sizeY)
	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			(*l)[j][i] = data.Rows[j].Values[i]
		}
	}
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
