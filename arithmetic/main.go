package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

const SYMBOL_COUNT = 256

var debug *bool
var timed *bool
var parallel *bool
var workers *int

func main() {
	parser := argparse.NewParser("encode", "STDIN arithmetic encoder/decoder")
	decode := parser.Flag("d", "decode", &argparse.Options{
		Required: false,
		Help:     "Run the decoder on the input data.",
	})
	encode := parser.Flag("e", "encode", &argparse.Options{
		Required: false,
		Help:     "Run the encoder on the input data. [default]",
	})
	timed = parser.Flag("t", "timed", &argparse.Options{
		Required: false,
		Help:     "Print each symbol encoding time.",
	})
	parallel = parser.Flag("p", "parallel", &argparse.Options{
		Required: false,
		Help:     "Run interval resizing using multiple threads",
	})
	workers = parser.Int("w", "workers", &argparse.Options{
		Required: false,
		Help:     "Number of workers when `--parallel` arguments is used",
		Default:  4,
	})
	debug = parser.Flag("v", "verbose", &argparse.Options{
		Required: false,
		Help:     "Print debug informations while encoding.",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if !*encode && !*decode {
		*encode = true
	}

	if *encode && *decode {
		fmt.Print("Cannot have both `--encode` and `--decode` at the same time")
		os.Exit(1)
	}

	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	defer stdout.Flush()

	if *encode {
		if err := encodeData(stdin, stdout); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	} else {
		if err := decodeData(stdin, stdout); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}
}
