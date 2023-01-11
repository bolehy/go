package main

import (
	"errors"
	"fmt"
)

func main() {

	circle_1 := -12.45

	area, err := Calc(circle_1)
	if err != nil {
		var areaError *AreaError
		if errors.As(err, &areaError) {
			fmt.Printf("Area calculation failed, radius %0.2f is less than zero and error because of %s\n", areaError.radius, areaError.err)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("area is %0.4f\n", area)
}
