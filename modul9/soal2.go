package main

import (
	"fmt"
)

func printArray(Array [10]int) {
	for i := 0; i < 10; i++ {
		fmt.Print(Array[i], " ")
	}
	fmt.Println()
}

func printArrayGanjil(Array [10]int) {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Print(Array[i], " ")
		}
	}
	fmt.Println()
}

func printArrayGenap(Array [10]int) {
	for i := 0; i < 10; i++ {
		if i == 0 || i%2 == 0 {
			fmt.Print(Array[i], " ")
		}
	}
	fmt.Println()
}

func printArrayKelipatanX(Array [10]int, X int) {
	for i := 0; i < 10; i++ {
		if i == 0 {

		} else if i%X == 0 {
			fmt.Print(Array[i], " ")
		}
	}
	fmt.Println()
}

func simpanganBaku(Array [10]int, Rerata int) {
	var selisih, selisihKuadrat, total int
	var simpanganBaku float64

	for i := 0; i < 10; i++ {
		selisih = Array[i] - Rerata
		selisihKuadrat = selisih * selisih
		total += selisihKuadrat
		simpanganBaku = float64(total) / 9
	}
	fmt.Printf("%.2f", simpanganBaku)
	fmt.Println()
}

func main() {
	var Array = [10]int{1, 3, 4, 4, 7, 8, 10, 4, 9, 10}
	var hapus = []int{}
	var X, jumlah, Rerata, frekuensi int

	fmt.Println("Cetak Array:")
	printArray(Array)
	fmt.Println("Cetak Array dengan index ganjil:")
	printArrayGanjil(Array)
	fmt.Println("Cetak Array dengan index genap:")
	printArrayGenap(Array)
	fmt.Println("Masukan Ankag kelipatan X: ")
	fmt.Scan(&X, "\n")
	printArrayKelipatanX(Array, X)

	for i := 0; i < 10; i++ {
		jumlah += Array[i]
	}
	Rerata = jumlah / 10
	fmt.Println("Rata-rata:", Rerata)

	fmt.Print("Simpangan baku: ")
	simpanganBaku(Array, Rerata)

	for i := 0; i < 10; i++ {
		if Array[i] == 4 {
			frekuensi++
		}
	}
	fmt.Println("Frekuensi bilangan 4:", frekuensi)

	fmt.Println("Penghapusan indeks 0:")
	for i := 1; i < 10; i++ {
		hapus = Array[:]
		fmt.Print(hapus[i], " ")
	}

}
