package main

// TODO: images must be multiple of 4 horizontally
// TODO: images must be multiple of 2 vertically

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"os"

	"github.com/akamensky/argparse"
)

var verbose *bool

func main() {
	parser := argparse.NewParser("jpeg2000", "JPEG2000 encoder and decoder")
	inputFile := parser.File("i", "input", os.O_RDONLY, 0600,
		&argparse.Options{
			Required: true,
			Help:     "The input file",
		})
	outputFile := parser.File("o", "output", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600,
		&argparse.Options{
			Required: true,
			Help:     "The output file",
		})

	cmdEncode := parser.NewCommand("encode", "Encode to JPEG2000")
	cmdDecode := parser.NewCommand("encode", "Decode from JPEG2000")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}

	defer inputFile.Close()
	defer outputFile.Close()

	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}

	if cmdEncode.Happened() {
		_, err = EncodeImage(inputImage)
		if err != nil {
			fmt.Fprintf(os.Stderr, parser.Usage(err))
			os.Exit(1)
		}
	}

	if cmdDecode.Happened() {
		//if err = DecodeImage(inputImage, outputFile); err != nil {
		//	fmt.Fprintf(os.Stderr, parser.Usage(err))
		//	os.Exit(1)
		//}
	}
}