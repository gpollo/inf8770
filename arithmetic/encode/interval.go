package main

import "math/big"
import "errors"
import "fmt"
import "os"

type Interval struct {
	number byte
	start  big.Rat
	end    big.Rat
}

type Intervals []Interval

func (i *Interval) size(x *big.Rat) *big.Rat {
	x.Sub(&i.end, &i.start)
	return x
}

func (i *Interval) resize(start1, start2, factor *big.Rat) {
	i.start.Sub(&i.start, start1)
	i.start.Mul(&i.start, factor)
	i.start.Add(&i.start, start2)

	i.end.Sub(&i.end, start1)
	i.end.Mul(&i.end, factor)
	i.end.Add(&i.end, start2)
}

func (i *Intervals) find(number byte) (Interval, error) {
	for j := 0; j < len(*i); j++ {
		if (*i)[j].number == number {
			return (*i)[j], nil
		}
	}

	return Interval{}, errors.New("Symbol not found")
}

func (i *Intervals) size(x *big.Rat) *big.Rat {
	x.Sub(&(*i)[len(*i)-1].end, &(*i)[0].start)
	return x
}

func (i *Intervals) resize(interval Interval) {
	start1 := &(*i)[0].start
	start2 := &interval.start
	size1 := &i.size()
	size2 := &interval.size()

	var factor big.Rat
	factor.Quo(size2, size1)

	fmt.Fprintf(os.Stderr, "%s %s %s\n",
		start1.Text('f', 40),
		start2.Text('f', 40),
		factor.Text('f', 40),
	)

	for j := 0; j < len(*i); j++ {
		(*i)[j].resize(&start1, &start2, &factor)
	}
}
