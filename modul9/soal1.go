package main

import (
	"fmt"
	"math"
)

type lingkaran struct {
	P1, P2, r float64
}

type KoorTitik struct {
	t1, t2 float64
}

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
	var lingkaran1, lingkaran2 lingkaran
	var titik KoorTitik

	fmt.Scan(&lingkaran1.P1, &lingkaran1.P2, &lingkaran1.r, &lingkaran2.P1, &lingkaran2.P2, &lingkaran2.r)
	fmt.Scan(&titik.t1, &titik.t2)

	if didalam(lingkaran1.P1, lingkaran1.P2, titik.t1, titik.t2, lingkaran1.r) == true &&
		didalam(lingkaran2.P1, lingkaran2.P2, titik.t1, titik.t2, lingkaran2.r) == true {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(lingkaran1.P1, lingkaran1.P2, titik.t1, titik.t2, lingkaran1.r) == true {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam(lingkaran2.P1, lingkaran2.P2, titik.t1, titik.t2, lingkaran2.r) == true {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
