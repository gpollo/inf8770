package main

import "math/big"

type Symbol struct {
	number    byte
	count     uint
	frequency big.Rat
}

func (s *Symbol) computeFrequency(int total) {
	fcount := big.NewRat(s[i].count, 1)
	ftotal := big.NewRat(total, 1)
	s.frequency.Quo(fcount, ftotal)
}

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
