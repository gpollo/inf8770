package main

import (
	"errors"
	"jpeg2000/data"
	"strconv"
	"strings"
)

type Wavelet interface {
	WaveletTransform(d ImageData) ImageData
	WaveletInverse(d ImageData) ImageData
	ToProtobuf() *data.WaveletConfig
}

func WaveletFromCommandLine(arg string) (Wavelet, error) {
	splited := strings.Split(arg, ":")
	if len(splited) <= 1 {
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
	default:
		return nil, errors.New("Unrecognized wavelet type")
	}
}

func WaveletFromProtobuf(d *data.WaveletConfig) (Wavelet, error) {
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
	case nil:
		return nil, errors.New("Wavelet configuration not found in protobuf data")
	default:
		return nil, errors.New("Unexpected wavelet configuration in protobuf data")
	}
}
