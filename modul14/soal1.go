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
	var n, m int
	var data [DMAX][NMAX]int
	var ukuran [DMAX]int
	var temp [NMAX]int
	var i, j int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&m)
		ukuran[i] = m

		for j = 0; j < m; j++ {
			fmt.Scan(&temp[j])
		}

		selectionSort(&temp, m)

		for j = 0; j < m; j++ {
			data[i][j] = temp[j]
		}
	}

	for i = 0; i < n; i++ {
		for j = 0; j < ukuran[i]; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(data[i][j])
		}
		fmt.Println()
	}
}
