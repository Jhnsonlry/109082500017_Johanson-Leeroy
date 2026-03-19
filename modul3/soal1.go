package main

import "fmt"

func faktorial(f int) int {
	var faktorial, i int
	faktorial = 1
	for i = 1; i <= f; i++ {
		faktorial = faktorial * i
	}
	return faktorial
}

func permutasi(p, q int) int {
	var permutasi int
	permutasi = faktorial(p) / faktorial(p-q)
	return permutasi
}

func kombinasi(x, y int) int {
	var kombinasi int
	kombinasi = faktorial(x) / (faktorial(y) * faktorial((x - y)))
	return kombinasi
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c {
		fmt.Println(permutasi(a, c), kombinasi(a, c))
	} else {
		fmt.Println(permutasi(c, a), kombinasi(c, a))
	}

	if b >= d {
		fmt.Println(permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Println(permutasi(d, b), kombinasi(d, b))
	}
}
