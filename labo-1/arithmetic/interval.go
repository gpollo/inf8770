package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"sync"
)

// This type defines an interval with a rational start and end. It allows
// the interval to be mapped and resize using only rational operation and
// keeping all precision.
type Interval struct {
	number byte
	start  big.Rat
	end    big.Rat
}

func (i *Interval) GetStart(x *big.Rat) *big.Rat {
	x.Set(&i.start)
	return x
}

func (i *Interval) GetEnd(x *big.Rat) *big.Rat {
	x.Set(&i.end)
	return x
}

func (i *Interval) GetSize(x *big.Rat) *big.Rat {
	var start, end big.Rat
	i.GetStart(&start)
	i.GetEnd(&end)
	x.Sub(&end, &start)
	return x
}

func (i *Interval) Contains(x *big.Rat) bool {
	if i.start.Cmp(x) > 0 {
		return false
	}

	if i.end.Cmp(x) < 0 {
		return false
	}

	return true
}

func (i *Interval) Resize(start1, start2, factor *big.Rat) {
	i.start.Sub(&i.start, start1)
	i.start.Mul(&i.start, factor)
	i.start.Add(&i.start, start2)

	i.end.Sub(&i.end, start1)
	i.end.Mul(&i.end, factor)
	i.end.Add(&i.end, start2)
}

// This type defines an interval composed of multiple increasing sub-interval
// that doesn't overlap with each other.
type Intervals []Interval

func (i *Intervals) GetStart(x *big.Rat) *big.Rat {
	if len(*i) == 0 {
		x.Set(big.NewRat(0, 1))
	} else {
		x.Set(&(*i)[0].start)
	}

	return x
}

func (i *Intervals) GetEnd(x *big.Rat) *big.Rat {
	if len(*i) == 0 {
		x.Set(big.NewRat(0, 1))
	} else {
		x.Set(&(*i)[len(*i)-1].end)
	}

	return x
}

func (i *Intervals) GetSize(x *big.Rat) *big.Rat {
	var start, end big.Rat
	i.GetStart(&start)
	i.GetEnd(&end)
	x.Sub(&end, &start)
	return x
}

func (i *Intervals) Resize(interval Interval) {
	var start1, start2 big.Rat
	i.GetStart(&start1)
	interval.GetStart(&start2)

	var size1, size2 big.Rat
	i.GetSize(&size1)
	interval.GetSize(&size2)

	var factor big.Rat
	factor.Quo(&size2, &size1)

	if !*parallel {
		for j := 0; j < len(*i); j++ {
			(*i)[j].Resize(&start1, &start2, &factor)
		}
	} else {
		var group sync.WaitGroup
		group.Add(*workers)

		var intervals chan *Interval = make(chan *Interval)
		for i := 0; i < *workers; i++ {
			go func() {
				defer group.Done()

				for {
					received := <-intervals
					if received == nil {
						break
					}
					received.Resize(&start1, &start2, &factor)
				}
			}()
		}

		for j := 0; j < len(*i); j++ {
			intervals <- &(*i)[j]
		}

		close(intervals)
		group.Wait()
	}
}

func (i *Intervals) Append(number byte, x *big.Rat) {
	interval := Interval{number: number}
	i.GetEnd(&interval.start)
	i.GetEnd(&interval.end)
	interval.end.Add(&interval.end, x)
	*i = append(*i, interval)
}

func (i *Intervals) Find(number byte) (Interval, error) {
	for j := 0; j < len(*i); j++ {
		if (*i)[j].number == number {
			return (*i)[j], nil
		}
	}

	return Interval{}, errors.New("Symbol not found")
}

func (i Intervals) FindByRat(x *big.Rat) (Interval, error) {
	for j := 0; j < len(i); j++ {
		if i[j].Contains(x) {
			return i[j], nil
		}
	}

	return Interval{}, errors.New("Interval not found")
}

func (i *Intervals) LowerThan(x *big.Rat) bool {
	var end big.Rat
	i.GetEnd(&end)

	return end.Cmp(x) < 0
}

func (i *Intervals) GreaterThan(x *big.Rat) bool {
	var start big.Rat
	i.GetStart(&start)

	return start.Cmp(x) > 0
}

func (i Intervals) Print() {
	for j := 0; j < len(i); j++ {
		fmt.Fprintf(os.Stderr, "%d (%c): [%s, %s]\n",
			i[j].number, i[j].number,
			i[j].start.FloatString(50),
			i[j].end.FloatString(50),
		)
	}
}
