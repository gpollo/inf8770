package histogram

import (
	"image"
	"math"
)

type Histogram struct {
	bins uint
	R    []uint
	G    []uint
	B    []uint
}

func FromImage(i image.Image, bins uint) *Histogram {
	h := Histogram{
		bins: bins,
		R:    make([]uint, bins),
		G:    make([]uint, bins),
		B:    make([]uint, bins),
	}

	h.Reset()
	h.AddImage(i)

	return &h
}

func (h *Histogram) Bins() uint {
	return h.bins
}

func (h *Histogram) Reset() {
	for i := uint(0); i < h.bins; i++ {
		h.R[i] = 0
		h.G[i] = 0
		h.B[i] = 0
	}
}

func (h *Histogram) AddImage(i image.Image) {
	for y := i.Bounds().Min.Y; y < i.Bounds().Max.Y; y++ {
		for x := i.Bounds().Min.X; x < i.Bounds().Max.X; x++ {
			r, g, b, _ := i.At(x, y).RGBA()

			if r == 0xffff {
				r--
			}

			if g == 0xffff {
				g--
			}

			if b == 0xffff {
				b--
			}

			step := math.Ceil(float64(0xffff) / float64(h.bins))
			h.R[int(math.Floor(float64(r)/step))]++
			h.G[int(math.Floor(float64(g)/step))]++
			h.B[int(math.Floor(float64(b)/step))]++
		}
	}
}

func (h *Histogram) DifferenceNorm(o *Histogram) (float64, float64, float64) {
	rSum := uint(0)
	gSum := uint(0)
	bSum := uint(0)

	for i := uint(0); i < h.bins; i++ {
		rSum += (h.R[i] - o.R[i]) * (h.R[i] - o.R[i])
		gSum += (h.G[i] - o.G[i]) * (h.G[i] - o.G[i])
		bSum += (h.B[i] - o.B[i]) * (h.B[i] - o.B[i])
	}

	rNorm := math.Sqrt(float64(rSum))
	gNorm := math.Sqrt(float64(gSum))
	bNorm := math.Sqrt(float64(bSum))

	return rNorm, gNorm, bNorm
}
