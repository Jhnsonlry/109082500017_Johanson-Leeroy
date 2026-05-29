package main

import (
	"fmt"
)

type ARR [1000]float64

func isiIkan(X *ARR, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&X[i])
	}
}

func rerata(X ARR, x, y int, totalWadah float64) {
	var jumlah, Total, rerata float64
	var no int
	for i := 0; i < x; i++ {
		jumlah = jumlah + X[i]
		if (i+1)%y == 0 {
			no++
			fmt.Println("Total berat Wadah", no, ":", jumlah)
			Total += jumlah
			jumlah = 0
		}

	}
	rerata = Total / totalWadah
	fmt.Println("Rata-rata berat ke", no, "wadah:", rerata)
}

func main() {
	var x, y int
	var totalWadah float64
	var dataIkan ARR
	fmt.Scan(&x, &y)
	totalWadah = float64(x) / float64(y)
	isiIkan(&dataIkan, x)
	rerata(dataIkan, x, y, totalWadah)
}
