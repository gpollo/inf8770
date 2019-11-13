package main

import (
	"decompose/sequence"
	"errors"
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	"os"

	"github.com/akamensky/argparse"
	_ "golang.org/x/image/bmp"
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

	_, err = sequence.FromDirectory(*directory, *format)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
