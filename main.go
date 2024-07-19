package main

import (
	"fmt"
)

func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	years := []int{2000, 1900, 2024, 2019}
	for _, year := range years {
		if IsLeapYear(year) {
			fmt.Printf("%d is a leap year.\n", year)
		} else {
			fmt.Printf("%d is not a leap year.\n", year)
		}
	}
}
