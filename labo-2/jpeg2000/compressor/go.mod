module jpeg2000/compressor

require (
	github.com/golang/protobuf v1.3.2 // indirect
	jpeg2000/data v0.0.0
	jpeg2000/helper v0.0.0
)

replace jpeg2000/data => ../data

replace jpeg2000/helper => ../helper

go 1.13
