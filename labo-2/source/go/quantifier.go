package main

type Quantifier interface {
	QuantifierTransform(d ImageData) ImageData
	QuantifierInverse(d ImageData) ImageData
}
