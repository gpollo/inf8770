package main

import (
	"errors"
	"jpeg2000/data"
	"strconv"
	"strings"
)

type Quantifier interface {
	QuantifierTransform(d ImageData) ImageData
	QuantifierInverse(d ImageData) ImageData
	ToProtobuf() *data.QuantifierConfig
}

func QuantifierFromCommandLine(arg string) (Quantifier, error) {
	splited := strings.Split(arg, ":")
	if len(splited) <= 1 {
		return nil, errors.New("Invalid number of argument for parsing quantifier")
	}

	switch splited[0] {
	case "deadzone":
		if len(splited) != 4 {
			return nil, errors.New("Invalid number of argument for parsing deadzone quantifier")
		}

		width, err := strconv.ParseInt(splited[1], 0, 32)
		if err != nil {
			return nil, err
		}

		delta, err := strconv.ParseInt(splited[2], 0, 32)
		if err != nil {
			return nil, err
		}

		offset, err := strconv.ParseFloat(splited[3], 32)
		if err != nil {
			return nil, err
		}

		return NewDeadZoneQuantifier(width, delta, offset)
	default:
		return nil, errors.New("Unrecognized quantifier type")
	}
}
