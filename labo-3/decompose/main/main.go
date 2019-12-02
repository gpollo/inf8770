package main

import (
	"decompose/expected"
	"decompose/sequence"
	"errors"
	"fmt"

	"os"

	"github.com/akamensky/argparse"
)

func execute() error {
	parser := argparse.NewParser("decompose", "Sequence decomposer")
	directory := parser.String("d", "directory", &argparse.Options{
		Required: false,
		Help:     "The directory containing the image",
		Default:  "images",
	})
	format := parser.String("f", "format", &argparse.Options{
		Required: false,
		Help:     "The filename format",
		Default:  "%04d.png",
	})
	expectedFilename := parser.String("e", "expected", &argparse.Options{
		Required: false,
		Help:     "The filename of the expected results",
		Default:  "anni005.txt",
	})
	skip := parser.Int("s", "skip", &argparse.Options{
		Required: false,
		Help:     "The number of frame to skip",
		Default:  1,
	})

	cmdSobel := parser.NewCommand("do-sobel", "Use sobel filter method")
	save := cmdSobel.Flag("s", "save", &argparse.Options{
		Required: false,
		Help:     "Save intermediary images",
		Default:  false,
	})

	cmdHistogram := parser.NewCommand("do-histogram", "Use histogram method")
	bins := cmdHistogram.Int("b", "bins", &argparse.Options{
		Required: false,
		Help:     "Number of bins in the histograms",
		Default:  20,
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	info, err := os.Stat(*directory)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return errors.New("The specified path is not a directory")
	}

	expected, err := expected.FromFile(*expectedFilename)
	if err != nil {
		return err
	}

	if cmdSobel.Happened() {
		sequence.FromDirectory(*directory, *format).RunSobel(expected, *save, *skip)
	}

	if cmdHistogram.Happened() {
		sequence.FromDirectory(*directory, *format).RunHistogram(expected, uint(*bins), *skip)
	}

	return nil
}

func main() {
	if err := execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
