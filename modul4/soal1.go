package main

import "fmt"

func faktorial(f int, faktorial *int) {
	var i int
	*faktorial = 1
	for i = 1; i <= f; i++ {
		*faktorial = *faktorial * i
	}
}

func permutasi(p, q int, permuatsi *int) {
	var r, s int
	faktorial(p, &r)
	faktorial(p-q, &s)
	*permuatsi = r / s
}

func kombinasi(x, y int, kombinasi *int) {
	var h, i, j int
	faktorial(x, &h)
	faktorial(y, &i)
	faktorial(x-y, &j)
	*kombinasi = h / (i * j)
}

func main() {
	var a, b, c, d, prmt, comb int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c {
		permutasi(a, c, &prmt)
		kombinasi(a, c, &comb)
		fmt.Println(prmt, comb)
	} else {
		permutasi(c, a, &prmt)
		kombinasi(c, a, &comb)
		fmt.Println(prmt, comb)
	}

	if b >= d {
		permutasi(b, d, &prmt)
		kombinasi(b, d, &comb)
		fmt.Println(prmt, comb)
	} else {
		permutasi(d, b, &prmt)
		kombinasi(d, b, &comb)
		fmt.Println(prmt, comb)
	}
}
