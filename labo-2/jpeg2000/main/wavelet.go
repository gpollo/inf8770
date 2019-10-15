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
