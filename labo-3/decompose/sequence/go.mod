module decompose/sequence

require (
	decompose/edges v0.0.0
	decompose/helper v0.0.0
	decompose/histogram v0.0.0
	decompose/layer v0.0.0
	decompose/sobel v0.0.0
)

replace decompose/histogram => ../histogram

replace decompose/layer => ../layer

replace decompose/helper => ../helper

replace decompose/sobel => ../sobel

replace decompose/edges => ../edges

go 1.13
