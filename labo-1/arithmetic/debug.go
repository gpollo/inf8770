package main

import (
	"fmt"
	"math/big"
	"os"
)

func debugHeaderEntryCount(c uint64) {
	if !*verbose {
		return
	}

	fmt.Fprintf(os.Stderr, "header: count=%d\n", c)
}

func debugHeaderEntry(s byte, c uint64) {
	if !*verbose {
		return
	}

	fmt.Fprintf(os.Stderr, "header: symbol=0x%02x (%c), count=%d\n", s, s, c)
}

func debugIntervals(i Intervals) {
	if !*verbose {
		return
	}

	var start, end big.Rat
	i.GetStart(&start)
	i.GetEnd(&end)
	fmt.Fprintf(os.Stderr, "data: [%s, %s]", start.FloatString(50), end.FloatString(50))
}

func debugRatFloat(n *big.Rat) {
	if !*verbose {
		return
	}

	fmt.Fprintf(os.Stderr, "data: %s\n", n.FloatString(100))
}
