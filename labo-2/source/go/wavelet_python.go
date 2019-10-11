package main

import (
	proto "github.com/golang/protobuf/proto"
)

type PyWavelet struct {
	mode string
}

func (w *PyWavelet) WaveletTransform(d ImageData) ImageData {
	image := ImageDataToProtobuf(d)
	dwt := ProtoDWT{Mode: w.mode, Data: &image}
	data, err := proto.Marshal(&dwt)
	if err != nil {
		panic(err.Error())
	}

	args := []string{"python3", "../python/dwt.py"}
	result, err := CallProcess(args, data)
	if err != nil {
		panic(err.Error())
	}

	wavelet := ProtoImageData{}
	err = proto.Unmarshal(result, &wavelet)
	if err != nil {
		panic(err.Error())
	}

	return ImageDataFromProtobuf(wavelet)
}

func (w *PyWavelet) WaveletInverse(d ImageData) ImageData {
	image := ImageDataToProtobuf(d)
	dwt := ProtoDWT{Mode: w.mode, Data: &image}
	data, err := proto.Marshal(&dwt)
	if err != nil {
		panic(err.Error())
	}

	args := []string{"python3", "../python/idwt.py"}
	result, err := CallProcess(args, data)
	if err != nil {
		panic(err.Error())
	}

	wavelet := ProtoImageData{}
	err = proto.Unmarshal(result, &wavelet)
	if err != nil {
		panic(err.Error())
	}

	return ImageDataFromProtobuf(wavelet)
}
