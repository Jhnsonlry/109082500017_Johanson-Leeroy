package main

import "fmt"

func ganjil(n, f int) {
	if f <= n {
		if f%2 != 0 {
			fmt.Print(f, " ")
		}
		ganjil(n, f+1)
	}
}

func main() {
	var n, f int
	f = 1
	fmt.Scan(&n)
	ganjil(n, f)

}
