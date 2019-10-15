module source/main

require (
	github.com/akamensky/argparse v0.0.0-20191006154803-1427fe674291
	github.com/golang/protobuf v1.3.2
	source/data v0.0.0
)

replace source/data => ../data

go 1.13
