package main

import "fmt"

func hitungSkor(soal, skor *int) {
	var nomor, waktu int
	nomor = 0
	waktu = 0
	*soal = 0
	*skor = 0
	for nomor = 8; nomor > 0; nomor-- {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*skor += waktu
			*soal++
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor, jumlahSoal, nilai int
	jumlahSoal = 0
	nilai = 0

	for {
		fmt.Scan(&nama)
		if nama == "selesai" || nama == "Selesai" {
			break
		}
		hitungSkor(&soal, &skor)

		if jumlahSoal < soal {
			jumlahSoal = soal
			nilai = skor
			pemenang = nama
		} else if jumlahSoal == soal {
			if nilai > skor {
				jumlahSoal = soal
				nilai = skor
				pemenang = nama
			}
		}
	}
	fmt.Println(pemenang, jumlahSoal, nilai)
}
