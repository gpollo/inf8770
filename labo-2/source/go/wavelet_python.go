package main

import (
	"source/data"

	proto "github.com/golang/protobuf/proto"
)

type PyWavelet struct {
	mode string
}

func (w *PyWavelet) WaveletTransform(d ImageData) ImageData {
	image := ImageDataToProtobuf(d)
	dwt := data.ProtoDWT{Mode: w.mode, Data: &image}
	pdata, err := proto.Marshal(&dwt)
	if err != nil {
		panic(err.Error())
	}

	args := []string{"python3", "../python/dwt.py"}
	result, err := CallProcess(args, pdata)
	if err != nil {
		panic(err.Error())
	}

	wavelet := data.ProtoImageData{}
	err = proto.Unmarshal(result, &wavelet)
	if err != nil {
		panic(err.Error())
	}

	decoded := ImageDataFromProtobuf(wavelet)
	decoded.Times(0.5)

	return decoded
}

func (w *PyWavelet) WaveletInverse(d ImageData) ImageData {
	image := ImageDataToProtobuf(d)
	dwt := data.ProtoDWT{Mode: w.mode, Data: &image}
	pdata, err := proto.Marshal(&dwt)
	if err != nil {
		panic(err.Error())
	}

	args := []string{"python3", "../python/idwt.py"}
	result, err := CallProcess(args, pdata)
	if err != nil {
		panic(err.Error())
	}

	wavelet := data.ProtoImageData{}
	err = proto.Unmarshal(result, &wavelet)
	if err != nil {
		panic(err.Error())
	}

	decoded := ImageDataFromProtobuf(wavelet)
	decoded.Times(2.0)

	return decoded
}
