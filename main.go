package main

import (
	"errors"
	"fmt"
)

type ErrorArea struct {
	err    string
	length float64
	width  float64
}

func (e *ErrorArea) Error() string {
	return e.err
}
func (e *ErrorArea) lengthNegative() bool {
	return e.length < 0
}
func (e *ErrorArea) widthNegative() bool {
	return e.width < 0
}

func RactArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += " , width is less than zero"
		}
	}
	if err != "" {
		return 0, &ErrorArea{
			err:    err,
			length: length,
			width:  width,
		}
	}
	return length * width, nil
}

func main() {

	raclength, racWidth := 12.5, 14.78

	racArea, err := RactArea(raclength, racWidth)

	if err != nil {
		var errorArea *ErrorArea
		if errors.As(err, &errorArea) {
			if errorArea.lengthNegative() {
				fmt.Printf("error, length % 0.4f is less than zero\n", raclength)
			}
			if errorArea.widthNegative() {
				fmt.Printf("error, width %0.4f is less than zero\n", racWidth)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf(" area of the ractungle is %0.4f qmm\n", racArea)
}
