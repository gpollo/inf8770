module jpeg2000/main

require (
	github.com/akamensky/argparse v0.0.0-20191006154803-1427fe674291
	github.com/golang/protobuf v1.3.2
	jpeg2000/data v0.0.0
)

replace jpeg2000/data => ../data

go 1.13
