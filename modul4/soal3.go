package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			fmt.Print(n, " ")
		} else {
			n = 3*n + 1
			fmt.Print(n, " ")
		}
	}
}

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	cetakDeret(bilangan)
}
