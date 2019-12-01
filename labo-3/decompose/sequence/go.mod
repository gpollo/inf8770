module decompose/sequence

require (
	decompose/edges v0.0.0
	decompose/expected v0.0.0
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

replace decompose/expected => ../expected

go 1.13
