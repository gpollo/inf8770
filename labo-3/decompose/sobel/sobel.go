package sobel

import (
	"decompose/layer"
	"math"
)

type Sobel struct {
	XKernel   [][]float64
	YKernel   [][]float64
	Crop      float64
	MinCrop   float64
	MaxCrop   float64
	MergeFunc func(uint, float64, float64) float64
}

func With33Kernel() *Sobel {
	return &Sobel{
		XKernel: [][]float64{
			{-1, 0, 1},
			{-2, 0, 2},
			{-1, 0, 1},
		},
		YKernel: [][]float64{
			{1, 2, 1},
			{0, 0, 0},
			{-1, -2, 1},
		},
		MinCrop:   20000,
		MaxCrop:   40000,
		MergeFunc: SumMerge,
	}
}

func With55Kernel() *Sobel {
	return &Sobel{
		XKernel: [][]float64{
			{-2, -1, 0, 1, 2},
			{-3, -2, 0, 2, 3},
			{-4, -3, 0, 3, 4},
			{-3, -2, 0, 2, 3},
			{-2, -1, 0, 1, 2},
		},
		YKernel: [][]float64{
			{2, 3, 4, 3, 2},
			{1, 2, 3, 2, 1},
			{0, 0, 0, 0, 0},
			{-1, -2, -3, -2, -1},
			{-2, -3, -4, -3, -2},
		},
		MinCrop:   20000,
		MaxCrop:   50000,
		MergeFunc: EuclideanDistanceMerge,
	}
}

func EuclideanDistanceMerge(count uint, gx, gy float64) float64 {
	return math.Sqrt(gx*gx+gy*gy) / float64(count)
}

func SumMerge(count uint, gx, gy float64) float64 {
	return (math.Abs(gx) + math.Abs(gy)) / float64(count/2)
}

func (s *Sobel) GetKernelDimension() uint {
	if len(s.XKernel) != len(s.YKernel) {
		panic("X and Y kernel sizes differ")
	}

	if len(s.XKernel) == 0 {
		return uint(0)
	}

	if len(s.XKernel[0]) != len(s.YKernel[0]) {
		panic("X and Y kernel sizes differ")
	}

	if len(s.XKernel)%2 == 0 {
		panic("X and Y kernel dimensions cannot be even")
	}

	if len(s.XKernel[0])%2 == 0 {
		panic("X and Y kernel dimensions cannot be even")
	}

	if len(s.XKernel) != len(s.XKernel[0]) {
		panic("X and Y kernel dimensions aren't equal")
	}

	return uint(len(s.XKernel))
}

func (s *Sobel) Apply(l layer.Layer) layer.Layer {
	sizeX, sizeY := l.GetDimensions()
	kernelDim := int(s.GetKernelDimension())
	border := int(kernelDim / 2)

	copy := layer.NewLayer(sizeX, sizeY)
	for y := border; y < sizeY-border; y++ {
		for x := int(border); x < sizeX-border; x++ {
			gx := float64(0)
			gy := float64(0)

			for j := -border; j <= border; j++ {
				for i := -border; i <= border; i++ {
					gx += s.XKernel[j+border][i+border] * l[y-j][x-i]
					gy += s.YKernel[j+border][i+border] * l[y-j][x-i]
				}
			}

			copy[y][x] = s.MergeFunc(uint(kernelDim*kernelDim), gx, gy)

			if copy[y][x] < s.MinCrop {
				copy[y][x] = float64(0x0000)
			} else if copy[y][x] > s.MaxCrop {
				copy[y][x] = float64(0xffff)
			}

			copy[y][x] = float64(0xffff) - copy[y][x]
		}
	}

	for y := 0; y <= border; y++ {
		for x := 0; x < sizeX; x++ {
			copy[y][x] = 0
		}
	}

	for y := sizeY - border; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			copy[y][x] = 0
		}
	}

	for y := 1; y < sizeY-1; y++ {
		for x := 0; x < int(border); x++ {
			copy[y][x] = 0
		}

		for x := sizeX - int(border); x < sizeX; x++ {
			copy[y][x] = 0
		}
	}

	return copy
}
