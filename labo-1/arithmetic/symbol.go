package main

import "math/big"

// This type is used for analysing symbols in data.
type Symbol struct {
	number    byte
	count     int
	frequency big.Rat
}

func (s *Symbol) computeFrequency(total int) {
	fcount := big.NewRat(int64(s.count), 1)
	ftotal := big.NewRat(int64(total), 1)
	s.frequency.Quo(fcount, ftotal)
}

// This type is used for sorting symbols by frequencies.
type Frequencies []Symbol

func (f Frequencies) Len() int {
	return len(f)
}

func (f Frequencies) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Frequencies) Less(i, j int) bool {
	return f[i].frequency.Cmp(&f[j].frequency) > 0
}
