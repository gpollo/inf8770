module decompose/main

require (
	decompose/average v0.0.0
	decompose/edges v0.0.0
	decompose/expected v0.0.0
	decompose/helper v0.0.0
	decompose/histogram v0.0.0
	decompose/layer v0.0.0
	decompose/sequence v0.0.0
	decompose/sobel v0.0.0
	github.com/akamensky/argparse v0.0.0-20191006154803-1427fe674291
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8
)

replace decompose/sequence => ../sequence

replace decompose/histogram => ../histogram

replace decompose/layer => ../layer

replace decompose/helper => ../helper

replace decompose/sobel => ../sobel

replace decompose/edges => ../edges

replace decompose/expected => ../expected

replace decompose/average => ../average

go 1.13
