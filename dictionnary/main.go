package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var beVerbose *bool

func main() {
	parser := argparse.NewParser("encode", "LZW encoder and decoder")
	doDecode := parser.Flag("d", "decode", &argparse.Options{
		Required: false,
		Help:     "Run the decoder on the input data.",
	})
	doEncode := parser.Flag("e", "encode", &argparse.Options{
		Required: false,
		Help:     "Run the encoder on the input data. [default]",
	})
	beVerbose = parser.Flag("v", "verbose", &argparse.Options{
		Required: false,
		Help:     "Print debug informations while processing.",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if !*doEncode && !*doDecode {
		*doEncode = true
	}

	if *doEncode && *doDecode {
		fmt.Print("Cannot have both `--encode` and `--decode` at the same time")
		os.Exit(1)
	}

	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	defer stdout.Flush()

	if *doEncode {
		if err := Encode(stdin, stdout); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	} else {
		if err := Decode(stdin, stdout); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}
}
