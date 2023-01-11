package main

import (
	"fmt"
	"math"
)

type AreaError struct {
	radius float64
	err    string
}

func (e *AreaError) Error() string {
	return fmt.Sprintf("radius %0.4f, %s", e.radius, e.err)
}

func Calc(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &AreaError{
			radius: radius,
			err:    "radius is negative",
		}
	}
	return math.Pi * radius * radius, nil
}
