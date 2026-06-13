package main

import "fmt"

const NMAX = 100000
const DMAX = 1000

func selectionSort(A *[NMAX]int, n int) {
	var i, j, idxMin, temp int

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if A[j] < A[idxMin] {
				idxMin = j
			}
		}

		temp = A[i]
		A[i] = A[idxMin]
		A[idxMin] = temp
	}
}

func main() {
	var n, m, x int
	var i, j, k int

	var hasil [DMAX][NMAX]int
	var ukuran [DMAX]int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		var ganjil, genap [NMAX]int
		var nGanjil, nGenap int

		fmt.Scan(&m)

		for j = 0; j < m; j++ {
			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil[nGanjil] = x
				nGanjil++
			} else {
				genap[nGenap] = x
				nGenap++
			}
		}

		selectionSort(&ganjil, nGanjil)
		selectionSort(&genap, nGenap)

		k = 0

		for j = 0; j < nGanjil; j++ {
			hasil[i][k] = ganjil[j]
			k++
		}

		for j = nGenap - 1; j >= 0; j-- {
			hasil[i][k] = genap[j]
			k++
		}

		ukuran[i] = k
	}

	for i = 0; i < n; i++ {
		for j = 0; j < ukuran[i]; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(hasil[i][j])
		}
		fmt.Println()
	}
}
