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
	parser := argparse.NewParser("jpeg2000", "JPEG2000 encoder and decoder")
	directory := parser.String("d", "directory", &argparse.Options{
		Required: true,
		Help:     "The directory containing the image",
	})
	format := parser.String("f", "format", &argparse.Options{
		Required: false,
		Help:     "The filename format",
		Default:  "%04d.jpg",
	})
	expectedFilename := parser.String("e", "expected", &argparse.Options{
		Required: false,
		Help:     "The filename of the expected results",
		Default:  "anni005.txt",
	})
	save := parser.Flag("s", "save", &argparse.Options{
		Required: false,
		Help:     "Save intermediary images",
		Default:  false,
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

	sequence.FromDirectory(*directory, *format).Run(expected, *save)

	return nil
}

func main() {
	if err := execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
