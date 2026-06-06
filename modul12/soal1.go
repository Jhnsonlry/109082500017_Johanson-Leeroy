package main

import "fmt"

const NMAX = 1000

type TabInt [NMAX]int

func sequentialSearch(A TabInt, n, x int) int {
	var i int
	i = 0
	for i < n && A[i] != x {
		i++
	}

	if i < n {
		return i
	}
	return -1
}

func hitungSuara(daftar *TabInt, nDaftar *int, suara *[21]int, total *int, sah *int) {
	var x, idx int

	for {
		fmt.Scan(&x)

		if x == 0 {
			*total++
			break
		}

		*total++

		if x >= 1 && x <= 20 {
			*sah++
			suara[x]++

			idx = sequentialSearch(*daftar, *nDaftar, x)
			if idx == -1 {
				daftar[*nDaftar] = x
				*nDaftar++
			}
		}
	}
}

func tampilkanHasil(daftar TabInt, nDaftar int, suara [21]int, total, sah int) {
	var i, j, temp int

	for i = 0; i < nDaftar-1; i++ {
		for j = i + 1; j < nDaftar; j++ {
			if daftar[i] > daftar[j] {
				temp = daftar[i]
				daftar[i] = daftar[j]
				daftar[j] = temp
			}
		}
	}

	fmt.Println("Suara masuk:", total-1)
	fmt.Println("Suara sah:", sah)

	for i = 0; i < nDaftar; i++ {
		fmt.Printf("%d: %d\n", daftar[i], suara[daftar[i]])
	}
}

func main() {
	var daftar TabInt
	var suara [21]int
	var nDaftar, total, sah int

	hitungSuara(&daftar, &nDaftar, &suara, &total, &sah)
	tampilkanHasil(daftar, nDaftar, suara, total, sah)
}
