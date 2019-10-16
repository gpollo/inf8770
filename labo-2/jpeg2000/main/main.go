package main

// TODO: images must be multiple of 4 horizontally
// TODO: images must be multiple of 2 vertically

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"io/ioutil"
	"jpeg2000/compressor"
	"jpeg2000/helper"
	"jpeg2000/quantifier"
	"jpeg2000/subsampler"
	"jpeg2000/wavelet"
	"os"

	"github.com/akamensky/argparse"
	_ "golang.org/x/image/bmp"
)

var verbose *bool

func execute() error {
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

	subsamplerConfig := cmdEncode.String("s", "subsampling",
		&argparse.Options{
			Required: false,
			Help:     "The subsampling format",
			Default:  "420",
		})
	waveletConfig := cmdEncode.String("w", "wavelet",
		&argparse.Options{
			Required: false,
			Help:     "The wavelet configuration",
			Default:  "haar:2",
		})
	quantifierConfig := cmdEncode.String("q", "quantifier",
		&argparse.Options{
			Required: false,
			Help:     "The quantifier configuration",
			Default:  "deadzone:5:2:0.0",
		})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	defer inputFile.Close()
	defer outputFile.Close()

	if cmdEncode.Happened() {
		inputImage, _, err := image.Decode(inputFile)
		if err != nil {
			return err
		}

		subsampler, err := subsampler.FromCommandLine(*subsamplerConfig)
		if err != nil {
			return err
		}

		wavelet, err := wavelet.FromCommandLine(*waveletConfig)
		if err != nil {
			return err
		}

		quantifier, err := quantifier.FromCommandLine(*quantifierConfig)
		if err != nil {
			return err
		}

		pipeline := Pipeline{
			subsampler: subsampler,
			wavelet:    wavelet,
			quantifier: quantifier,
			compressor: &compressor.LZWCompressor{},
		}

		encoded, err := pipeline.EncodeImage(inputImage)
		if err != nil {
			return err
		}

		_, err = outputFile.Write(encoded)
		if err != nil {

			return err
		}
	}

	if cmdDecode.Happened() {
		inputImage, err := ioutil.ReadAll(inputFile)
		if err != nil {
			return err
		}

		pipeline := Pipeline{}
		decoded, err := pipeline.DecodeImage(inputImage)
		if err != nil {
			return err
		}

		err = helper.SaveImage(decoded, outputFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
