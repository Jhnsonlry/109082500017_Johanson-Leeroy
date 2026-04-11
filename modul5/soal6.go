package main

import "fmt"

func pangkat(hasil, x, y int) {
	if y > 0 {
		hasil *= x
		pangkat(hasil, x, y-1)
	} else {
		fmt.Print(hasil)
	}
}

func main() {
	var hasil, x, y int
	hasil = 1
	fmt.Scan(&x, &y)
	pangkat(hasil, x, y)
}
