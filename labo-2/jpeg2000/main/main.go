package main

// TODO: images must be multiple of 4 horizontally
// TODO: images must be multiple of 2 vertically

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"io/ioutil"
	"os"

	"github.com/akamensky/argparse"
	_ "golang.org/x/image/bmp"
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
	cmdDecode := parser.NewCommand("decode", "Decode from JPEG2000")

	waveletConfig := cmdEncode.String("w", "wavelet",
		&argparse.Options{
			Required: false,
			Help:     "The wavelet to use",
			Default:  "haar:2",
		})
	quantifierConfig := cmdEncode.String("q", "quantifier",
		&argparse.Options{
			Required: false,
			Help:     "The quantifier to use",
			Default:  "deadzone:5:2:0.0",
		})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}

	defer inputFile.Close()
	defer outputFile.Close()

	if cmdEncode.Happened() {
		inputImage, _, err := image.Decode(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, parser.Usage(err))
			os.Exit(1)
		}

		wavelet, err := WaveletFromCommandLine(*waveletConfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, parser.Usage(err))
			os.Exit(1)
		}

		quantifier, err := QuantifierFromCommandLine(*quantifierConfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, parser.Usage(err))
			os.Exit(1)
		}

		pipeline := Pipeline{
			subsampler: &Subsampler444{},
			wavelet:    wavelet,
			quantifier: quantifier,
			compressor: &LZWCompressor{},
		}

		encoded, err := pipeline.EncodeImage(inputImage)
		if err != nil {
			fmt.Fprintf(os.Stderr, parser.Usage(err))
			os.Exit(1)
		}

		_, err = outputFile.Write(encoded)
		if err != nil {
			fmt.Fprintf(os.Stderr, parser.Usage(err))
			os.Exit(1)
		}
	}

	if cmdDecode.Happened() {
		inputImage, err := ioutil.ReadAll(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, parser.Usage(err))
			os.Exit(1)
		}

		pipeline := Pipeline{}
		decoded, err := pipeline.DecodeImage(inputImage)
		if err != nil {
			fmt.Fprintf(os.Stderr, parser.Usage(err))
			os.Exit(1)
		}

		err = SaveImage(decoded, outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, parser.Usage(err))
			os.Exit(1)
		}
	}
}
