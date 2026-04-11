package main

import "fmt"

func faktor(n, f int) {
	if f <= n {
		if n%f == 0 {
			fmt.Print(f, " ")
		}
		faktor(n, f+1)
	}
}

func main() {
	var n, f int
	f = 1
	fmt.Scan(&n)
	faktor(n, f)

}
