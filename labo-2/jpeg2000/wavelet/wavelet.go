package wavelet

import (
	"errors"
	"jpeg2000/data"
	"strconv"
	"strings"
)

type Wavelet interface {
	WaveletTransform(l data.Layer) data.Layer
	WaveletInverse(l data.Layer) data.Layer
	ToProtobuf() *data.WaveletConfig
}

func scaleX(f1, f2 data.Layer) data.Layer {
	sizeX1, sizeY1 := f1.GetDimensions()
	sizeX2, sizeY2 := f2.GetDimensions()

	if sizeX1 != sizeX2 || sizeY1 != sizeY2 {
		panic("Image dimensions aren't equal")
	}

	sizeX := sizeX1
	sizeY := sizeY1
	data := data.NewLayer(2*sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			data[j][2*i+0] = f1[j][i] + f2[j][i]
			data[j][2*i+1] = f1[j][i] - f2[j][i]
		}
	}

	return data
}

func scaleY(f1, f2 data.Layer) data.Layer {
	sizeX1, sizeY1 := f1.GetDimensions()
	sizeX2, sizeY2 := f2.GetDimensions()

	if sizeX1 != sizeX2 || sizeY1 != sizeY2 {
		panic("Image dimensions aren't equal")
	}

	sizeX := sizeX1
	sizeY := sizeY1
	data := data.NewLayer(sizeX, 2*sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			data[2*j+0][i] = f1[j][i] + f2[j][i]
			data[2*j+1][i] = f1[j][i] - f2[j][i]
		}
	}

	return data
}

func copyIntoQuadrant(from, into data.Layer, quadrant int) {
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

func copyFromQuadrant(from, into data.Layer, quadrant int) {
	sizeFromX, sizeFromY := from.GetDimensions()
	sizeIntoX, sizeIntoY := into.GetDimensions()

	if 2*sizeIntoX > sizeFromX {
		panic("Invalid X size for copying into")
	}

	if 2*sizeIntoY > sizeFromY {
		panic("Invalid Y size for copying into")
	}

	var offsetX int
	var offsetY int
	switch quadrant {
	case 1:
		offsetX = sizeIntoX
		offsetY = 0
	case 2:
		offsetX = 0
		offsetY = 0
	case 3:
		offsetX = 0
		offsetY = sizeIntoY
	case 4:
		offsetX = sizeIntoX
		offsetY = sizeIntoY
	default:
		panic("Invalid quadrant selected for copying")
	}

	for j := 0; j < sizeIntoY; j++ {
		for i := 0; i < sizeIntoX; i++ {
			into[j][i] = from[j+offsetY][i+offsetX]
		}
	}
}

func FromCommandLine(arg string) (Wavelet, error) {
	splited := strings.Split(arg, ":")
	if len(splited) < 1 {
		return nil, errors.New("Invalid number of argument for parsing wavelet")
	}

	switch splited[0] {
	case "haar":
		if len(splited) != 2 {
			return nil, errors.New("Invalid number of argument for parsing haar wavelet")
		}

		level, err := strconv.ParseInt(splited[1], 0, 32)
		if err != nil {
			return nil, err
		}

		return NewHaarWavelet(level)
	case "daubechies":
		if len(splited) != 3 {
			return nil, errors.New("Invalid number of argument for parsing daubechies wavelet")
		}

		level, err := strconv.ParseInt(splited[1], 0, 32)
		if err != nil {
			return nil, err
		}

		coefficient, err := strconv.ParseInt(splited[2], 0, 32)
		if err != nil {
			return nil, err
		}

		return NewDaubechiesWavelet(level, coefficient)
	case "dummy":
		if len(splited) != 1 {
			return nil, errors.New("Invalid number of argument for parsing dummy wavelet")
		}

		return NewDummyWavelet(), nil
	default:
		return nil, errors.New("Unrecognized wavelet type")
	}
}

func FromProtobuf(d *data.WaveletConfig) (Wavelet, error) {
	switch d.Data.(type) {
	case *data.WaveletConfig_Haar:
		haar := HaarWavelet{}
		if err := haar.FromProtobuf(d); err != nil {
			return nil, err
		} else {
			return &haar, nil
		}
	case *data.WaveletConfig_Daubechies:
		daubechies := DaubechiesWavelet{}
		if err := daubechies.FromProtobuf(d); err != nil {
			return nil, err
		} else {
			return &daubechies, nil
		}
	case *data.WaveletConfig_Dummy:
		dummy := DummyWavelet{}
		if err := dummy.FromProtobuf(d); err != nil {
			return nil, err
		} else {
			return &dummy, nil
		}
	case nil:
		return nil, errors.New("Wavelet configuration not found in protobuf data")
	default:
		return nil, errors.New("Unexpected wavelet configuration in protobuf data")
	}
}
