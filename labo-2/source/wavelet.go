package main

type Wavelet interface {
	WaveletTransform(d ImageData) ImageData
}

type HaarWavelet struct {
	level int
}

func (w *HaarWavelet) GetXLowPassFilter(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX/2, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX/2; i++ {
			value := 0
			value += int(d[j][2*i+0])
			value += int(d[j][2*i+1])
			data[j][i] = byte(value / 2)
		}
	}

	return data
}

func (w *HaarWavelet) GetYLowPassFilter(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX, sizeY/2)

	for j := 0; j < sizeY/2; j++ {
		for i := 0; i < sizeX; i++ {
			value := 0
			value += int(d[2*j+0][i])
			value += int(d[2*j+1][i])
			data[j][i] = byte(value / 2)
		}
	}

	return data
}

func (w *HaarWavelet) GetXHighPassFilter(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX/2, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX/2; i++ {
			value := 0
			value += int(d[j][2*i+0])
			value -= int(d[j][2*i+1])
			if value < 0 {
				data[j][i] = byte(0)
			} else {
				data[j][i] = byte(value / 2)
			}
		}
	}

	return data
}

func (w *HaarWavelet) GetYHighPassFilter(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX, sizeY/2)

	for j := 0; j < sizeY/2; j++ {
		for i := 0; i < sizeX; i++ {
			value := 0
			value += int(d[2*j+0][i])
			value -= int(d[2*j+1][i])
			if value < 0 {
				data[j][i] = byte(0)
			} else {
				data[j][i] = byte(value / 2)
			}
		}
	}

	return data
}

func (w *HaarWavelet) CopyIntoQuadrant(from, into ImageData, quadrant int) {
	sizeFromX, sizeFromY := from.GetDimensions()
	sizeIntoX, sizeIntoY := into.GetDimensions()

	if 2*sizeFromX > sizeIntoX {
		panic("Invalid X size for copying from")
	}

	if 2*sizeFromY > sizeIntoY {
		panic("Invalid Y size for copying from")
	}

	var offsetX int
	var offsetY int
	switch quadrant {
	case 1:
		offsetX = sizeFromX
		offsetY = 0
	case 2:
		offsetX = 0
		offsetY = 0
	case 3:
		offsetX = 0
		offsetY = sizeFromY
	case 4:
		offsetX = sizeFromX
		offsetY = sizeFromY
	default:
		panic("Invalid quadrant selected for copying")
	}

	for j := 0; j < sizeFromY; j++ {
		for i := 0; i < sizeFromX; i++ {
			into[j+offsetY][i+offsetX] = from[j][i]
		}
	}
}

func (w *HaarWavelet) WaveletTransform(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX, sizeY)

	level := w.level
	if level == 0 {
		level = 2
	}

	for i := 0; i < level; i++ {
		fll := w.GetYLowPassFilter(w.GetXLowPassFilter(d))
		flh := w.GetYHighPassFilter(w.GetXLowPassFilter(d))
		fhl := w.GetYLowPassFilter(w.GetXHighPassFilter(d))
		fhh := w.GetYHighPassFilter(w.GetXHighPassFilter(d))

		w.CopyIntoQuadrant(fll, data, 2)
		w.CopyIntoQuadrant(flh, data, 1)
		w.CopyIntoQuadrant(fhl, data, 3)
		w.CopyIntoQuadrant(fhh, data, 4)

		d = fll
	}

	return data
}
