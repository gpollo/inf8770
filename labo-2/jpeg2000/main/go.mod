module jpeg2000/main

require (
	github.com/akamensky/argparse v0.0.0-20191006154803-1427fe674291
	github.com/golang/protobuf v1.3.2
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8
	jpeg2000/data v0.0.0
	jpeg2000/helper v0.0.0
)

replace jpeg2000/data => ../data

replace jpeg2000/helper => ../helper

go 1.13
