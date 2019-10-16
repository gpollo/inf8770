package main

import (
	"encoding/binary"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/image/bmp"
)

func ReadAll(r io.Reader) ([]byte, error) {
	data := []byte{}

	for {
		buffer := make([]byte, 1024)

		n, err := r.Read(buffer)
		if err != nil && err != io.EOF {
			return []byte{}, err
		}

		data = append(data, buffer[:n]...)
		if n != len(buffer) {
			break
		}
	}

	return data, nil
}

func CallProcess(args []string, data []byte) ([]byte, error) {
	command := exec.Command(args[0], args[1:]...)
	command.Stderr = os.Stderr

	inputW, err := command.StdinPipe()
	if err != nil {
		return []byte{}, err
	}

	outputR, err := command.StdoutPipe()
	if err != nil {
		return []byte{}, err
	}

	if err := command.Start(); err != nil {
		return []byte{}, err
	}

	_, err = inputW.Write(data)
	if err != nil {
		return []byte{}, err
	}
	inputW.Close()

	result, err := ReadAll(outputR)
	if err != nil {
		return []byte{}, err
	}
	outputR.Close()

	command.Wait()
	return result, nil
}

func WriteVarint(w io.Writer, x int64) error {
	buffer := make([]byte, 64)
	count := binary.PutVarint(buffer, x)

	if _, err := w.Write(buffer[0:count]); err != nil {
		return err
	}

	return nil
}

func SaveImage(i image.Image, f *os.File) error {
	splited := strings.Split(f.Name(), ".")
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
