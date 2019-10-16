package main

import (
	"errors"
	"jpeg2000/data"
	"strconv"
	"strings"
)

type Quantifier interface {
	QuantifierTransform(l data.Layer) data.Layer
	QuantifierInverse(l data.Layer) data.Layer
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

func QuantifierFromProtobuf(d *data.QuantifierConfig) (Quantifier, error) {
	switch d.Data.(type) {
	case *data.QuantifierConfig_DeadZone:
		deadzone := DeadZoneQuantifier{}
		if err := deadzone.FromProtobuf(d); err != nil {
			return nil, err
		} else {
			return &deadzone, nil
		}
	case nil:
		return nil, errors.New("Quantifier configuration not found in protobuf data")
	default:
		return nil, errors.New("Unexpected quantifier configuration in protobuf data")
	}
}
