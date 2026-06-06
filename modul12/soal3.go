package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k, idx int

	fmt.Scan(&n, &k)

	isiArray(n)

	idx = posisi(n, k)

	if idx == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(idx)
	}
}

func isiArray(n int) {
	/* I.S. terdefinisi integer n, dan sejumlah n data sudah siap pada piranti masukan.
	   F.S. Array data berisi n (<=NMAX) bilangan */
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	/* mengembalikan posisi k dalam array data dengan n elemen.Posisi dimulai dari posisi 0.
	Jika tidak ada kembalikan -1 */

	var kiri, kanan, tengah int

	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if data[tengah] == k {
			return tengah
		} else if data[tengah] < k {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}
