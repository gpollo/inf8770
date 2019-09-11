package main

import (
	"fmt"
	"os"
)

func debugHeaderCount(c uint64) {
	if !*beVerbose {
		return
	}

	fmt.Fprintf(os.Stderr, "header: count=%d\n", c)
}

func debugHeaderEntry(s byte) {
	if !*beVerbose {
		return
	}

	fmt.Fprintf(os.Stderr, "header: symbol=0x%02x (%c)\n", s, s)
}

func debugNewEntry(p []byte, v uint64) {
	if !*beVerbose {
		return
	}

	fmt.Fprintf(os.Stderr, "new entry: value=%d, path=", v)
	for _, b := range p {
		fmt.Fprintf(os.Stderr, "%02x ", b)
	}
	fmt.Fprint(os.Stderr, "(")
	for _, b := range p {
		fmt.Fprintf(os.Stderr, "%c", b)
	}
	fmt.Fprint(os.Stderr, ")\n")
}

func debugOutput(v uint64) {
	if !*beVerbose {
		return
	}

	fmt.Fprintf(os.Stderr, "data: %d\n", v)
}
