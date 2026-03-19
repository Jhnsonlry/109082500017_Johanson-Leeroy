package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	var jarak float64
	jarak = math.Sqrt(((a - c) * (a - c)) + ((b - d) * (b - d)))
	return jarak
}

func didalam(p1, p2, t1, t2, r1 float64) bool {
	var didalam bool
	didalam = false
	if r1 >= jarak(p1, p2, t1, t2) {
		didalam = true
	}
	return didalam
}

func main() {
	var p1, p2, r1, pp1, pp2, r2, t1, t2 float64

	fmt.Scan(&p1, &p2, &r1, &pp1, &pp2, &r2, &t1, &t2)

	if didalam(p1, p2, t1, t2, r1) == true && didalam(pp1, pp2, t1, t2, r2) == true {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(p1, p2, t1, t2, r1) == true {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam(pp1, pp2, t1, t2, r2) == true {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
