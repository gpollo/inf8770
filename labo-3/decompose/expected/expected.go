package expected

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Expected struct {
	cuts       map[uint]uint
	cutsGood   uint
	cutsBad    uint
	cutsMissed uint
	fades      [][2]uint
	fadesGood  uint
	fadesBad   uint
}

func Empty() *Expected {
	return &Expected{
		cuts:       make(map[uint]uint),
		cutsGood:   0,
		cutsBad:    0,
		cutsMissed: 0,
		fades:      [][2]uint{},
		fadesGood:  0,
		fadesBad:   0,
	}
}

func (e *Expected) AddCut(position int) error {
	if position < 0 {
		return fmt.Errorf("Invalid cut frame (%d)", position)
	}

	e.cuts[uint(position)] = uint(position)

	return nil
}

func (e *Expected) AddFade(start, end int) error {
	if start < 0 {
		return fmt.Errorf("Invalid start frame (%d)", start)
	}

	if end < 0 {
		return fmt.Errorf("Invalid end frame (%d)", end)
	}

	if start >= end {
		return fmt.Errorf("Invalid start (%d) and end (%d) frame", start, end)
	}

	e.fades = append(e.fades, [2]uint{uint(start), uint(end)})

	return nil
}

func FromFile(filename string) (*Expected, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	defer file.Close()

	expected := Empty()
	for {
		line, prefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		if prefix {
			return nil, errors.New("Unsupported prefix line")
		}

		numbers := strings.Split(string(line), " ")
		switch len(numbers) {
		case 1:
			position, err := strconv.Atoi(numbers[0])
			if err != nil {
				return nil, err
			}

			if err = expected.AddCut(position); err != nil {
				return nil, err
			}
		case 2:
			start, err := strconv.Atoi(numbers[0])
			if err != nil {
				return nil, err
			}

			end, err := strconv.Atoi(numbers[1])
			if err != nil {
				return nil, err
			}

			if err = expected.AddFade(start, end); err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("Unsupported line format")
		}
	}

	return expected, nil
}

func (e *Expected) CheckValidCut(position uint) bool {
	if _, ok := e.cuts[position]; ok {
		fmt.Printf("Ok:  %d is a valid cut frame\n", position)
		e.cutsGood++
		return true
	}

	fmt.Printf("Err: %d isn't a valid cut frame\n", position)
	e.cutsBad++
	return false
}

func (e *Expected) CheckInvalidCut(position uint) bool {
	if _, ok := e.cuts[position]; ok {
		fmt.Printf("Err: %d is a valid cut frame\n", position)
		e.cutsMissed++
		return false
	}

	return true
}

func (e *Expected) CheckValidFade(start, end uint) bool {
	for _, r := range e.fades {
		if start < r[0] && end < r[0] {
			continue
		}

		if start > r[1] && end > r[1] {
			continue
		}

		fmt.Printf("Ok:  [%d, %d] is a valid fade\n", start, end)
		e.fadesGood++
		return true
	}

	fmt.Printf("Err:  [%d, %d] isn't a valid fade\n", start, end)
	e.fadesBad++
	return false
}

func (e *Expected) Report() {
	fmt.Printf("Good Cuts:   %3d\n", e.cutsGood)
	fmt.Printf("Bad Cuts:    %3d\n", e.cutsBad)
	fmt.Printf("Missed Cuts: %3d\n", e.cutsMissed)
	fmt.Printf("Good Fades:  %3d\n", e.fadesGood)
	fmt.Printf("Bad Fades:   %3d\n", e.fadesBad)
	fmt.Printf("Total Cuts:  %3d\n", len(e.cuts))
	fmt.Printf("Total Fades: %3d\n", len(e.fades))
}
