package main

import "fmt"

const NMAX = 1000

func insertionSort(A *[NMAX]int, n int) {
	var pass, i, temp int

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass - 1

		for i >= 0 && A[i] > temp {
			A[i+1] = A[i]
			i--
		}

		A[i+1] = temp
	}
}

func main() {
	var A [NMAX]int
	var n, x int
	var i, jarak int
	var tetap bool

	fmt.Scan(&x)
	for x >= 0 {
		A[n] = x
		n++
		fmt.Scan(&x)
	}

	insertionSort(&A, n)

	for i = 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(A[i])
	}
	fmt.Println()

	if n <= 1 {
		fmt.Println("Data berjarak 0")
	} else {
		jarak = A[1] - A[0]
		tetap = true

		for i = 2; i < n; i++ {
			if A[i]-A[i-1] != jarak {
				tetap = false
			}
		}

		if tetap {
			fmt.Println("Data berjarak", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}
