package main

import (
	"errors"
	"fmt"
	"log"
)

func Divide(a int, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return float64(a) / float64(b), nil
}

func main() {

	a, b := 10, 2
	result, err := Divide(a, b)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Divide(%d, %d) = %.2f\n", a, b, result)

}
