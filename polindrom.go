package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if (x < 0 && (func() bool {
		return ((x * -1) == (func() int {
			r, t := 0, x*-1
			for t > 0 {
				r = r*10 + t%10
				t /= 10
			}
			return r
		})())
	})()) || (x >= 0 && (func() bool {
		if x < 10 {
			return true
		}

		digitCount := 0
		for t := x; t > 0; digitCount++ {
			t /= 10
		}

		for i := 0; i < (digitCount+1)/2; i++ {
			leftPos := digitCount - 1 - i
			rightPos := i

			leftDigit := (x / pow(10, leftPos)) % 10
			rightDigit := (x / pow(10, rightPos)) % 10

			if leftDigit != rightDigit {
				return false
			}
		}
		return true
	})()) {
		fmt.Println("polindrom")
	} else {
		fmt.Println("bukan polindrom")
	}
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
