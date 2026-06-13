package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var i, idxMax int

	idxMax = 0

	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}

	fmt.Println(
		pustaka[idxMax].judul,
		pustaka[idxMax].penulis,
		pustaka[idxMax].penerbit,
		pustaka[idxMax].tahun,
	)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var pass, i int
	var temp Buku

	for pass = 1; pass < n; pass++ {
		temp = pustaka[pass]
		i = pass - 1

		for i >= 0 && pustaka[i].rating < temp.rating {
			pustaka[i+1] = pustaka[i]
			i--
		}

		pustaka[i+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int

	if n < 5 {
		batas = n
	} else {
		batas = 5
	}

	for i = 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	var left, right, mid int
	var ketemu bool

	left = 0
	right = n - 1

	for left <= right && !ketemu {
		mid = (left + right) / 2

		if pustaka[mid].rating == r {
			ketemu = true
		} else if r > pustaka[mid].rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if ketemu {
		fmt.Println(
			pustaka[mid].judul,
			pustaka[mid].penulis,
			pustaka[mid].penerbit,
			pustaka[mid].tahun,
			pustaka[mid].eksemplar,
			pustaka[mid].rating,
		)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}
