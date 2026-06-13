package main

import "fmt"

const NMAX = 1000000

func insertionSort(A *[NMAX]int, n int) {
	var i, j, temp int

	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1

		for j >= 0 && A[j] > temp {
			A[j+1] = A[j]
			j--
		}

		A[j+1] = temp
	}
}

func median(A [NMAX]int, n int) int {
	if n%2 == 1 {
		return A[n/2]
	}
	return (A[n/2-1] + A[n/2]) / 2
}

func main() {
	var data [NMAX]int
	var n int
	var x int

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {
			insertionSort(&data, n)
			fmt.Println(median(data, n))
		} else {
			data[n] = x
			n++
		}
	}
}
