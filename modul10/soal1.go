package main

import "fmt"

type arr [1000]int

func isiData(X *arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&X[i])
	}
}

func terbesar(X arr, n int) {
	var terbesar int = X[0]
	for i := 0; i < n; i++ {
		if X[i] > terbesar {
			terbesar = X[i]
		}
	}
	fmt.Println("Berat Terbesar:", terbesar)
}

func terkecil(X arr, n int) {
	var terkecil int = X[0]
	for i := 0; i < n; i++ {
		if X[i] < terkecil {
			terkecil = X[i]
		}
	}
	fmt.Println("Berat Terkecil:", terkecil)
}

func main() {
	var n int
	var data arr

	fmt.Scan(&n)
	isiData(&data, n)
	terbesar(data, n)
	terkecil(data, n)
}
