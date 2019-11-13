package helper

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/bmp"
)

func SaveImage(i image.Image, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	splited := strings.Split(filename, ".")
	if len(splited) <= 1 {
		return errors.New("Extension not found in output filename")
	}

	extension := splited[len(splited)-1]

	switch extension {
	case "bmp":
		if err := bmp.Encode(f, i); err != nil {
			return err
		}
	case "jpg":
		fallthrough
	case "jpeg":
		if err := jpeg.Encode(f, i, nil); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(f, i); err != nil {
			return err
		}
	}

	return nil
}
