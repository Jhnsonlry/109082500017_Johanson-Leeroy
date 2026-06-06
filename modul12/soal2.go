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

func bacaSuara(daftar *TabInt, nDaftar *int, suara *[21]int, total *int, sah *int) {
	var x, idx int

	for {
		fmt.Scan(&x)

		*total++

		if x == 0 {
			break
		}

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

func cariKetuaNWakil(suara [21]int, ketua, wakil *int) {
	var maks1, maks2 int
	var i int

	for i = 1; i <= 20; i++ {
		if suara[i] > maks1 {
			maks1 = suara[i]
		}
	}

	for i = 1; i <= 20; i++ {
		if suara[i] == maks1 {
			*ketua = i
			break
		}
	}

	for i = *ketua + 1; i <= 20; i++ {
		if suara[i] == maks1 {
			*wakil = i
			return
		}
	}

	for i = 1; i <= 20; i++ {
		if suara[i] > maks2 && suara[i] < maks1 {
			maks2 = suara[i]
		}
	}

	for i = 1; i <= 20; i++ {
		if suara[i] == maks2 {
			*wakil = i
			break
		}
	}
}

func main() {
	var daftar TabInt
	var suara [21]int
	var nDaftar, total, sah int
	var ketua, wakil int

	bacaSuara(&daftar, &nDaftar, &suara, &total, &sah)
	cariKetuaNWakil(suara, &ketua, &wakil)

	fmt.Println("Suara masuk:", total-1)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
