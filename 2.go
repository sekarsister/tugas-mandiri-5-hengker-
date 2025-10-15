package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n < 0 {
		return
	}

	if n >= 0 {
		fmt.Print(0, " ")
	}
	if n >= 1 {
		fmt.Print(1, " ")
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		next := a + b
		fmt.Print(next, " ")
		a, b = b, next
	}
}
