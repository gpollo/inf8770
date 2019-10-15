package main

import "fmt"

type DaubechiesWavelet struct {
	level uint
}

func NewDaubechiesWavelet() DaubechiesWavelet {
	return DaubechiesWavelet{level: 2}
}

func (w *DaubechiesWavelet) SetLevel(level uint) {
	if level < 2 {
		return
	}

	w.level = level
}

func (w *DaubechiesWavelet) WaveletTransform(d ImageData) ImageData {
	pyWavelet := PyWavelet{mode: fmt.Sprintf("db%d", w.level)}
	return pyWavelet.WaveletTransform(d)
}

func (w *DaubechiesWavelet) WaveletInverse(d ImageData) ImageData {
	pyWavelet := PyWavelet{mode: fmt.Sprintf("db%d", w.level)}
	return pyWavelet.WaveletInverse(d)
}
