package main

import "fmt"

func Fibonanci(n, x1, x2, x3 int) {

	if n > 0 {
		x1 = x2
		x2 = x3
		x3 = x2 + x1
		fmt.Print(x2, " ")
		Fibonanci(n-1, x1, x2, x3)
	}
}

func main() {
	var x1, x2, x3 int
	x1 = 0
	x2 = 0
	x3 = 1
	fmt.Print(x2, " ")
	Fibonanci(10, x1, x2, x3)
}
