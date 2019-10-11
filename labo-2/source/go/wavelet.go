package main

type Wavelet interface {
	WaveletTransform(d ImageData) ImageData
	WaveletInverse(d ImageData) ImageData
}
