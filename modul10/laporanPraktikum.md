# <h1 align="center">Laporan Praktikum Modul 10 - PENCARIAN NILAI MAX/MIN </h1>
<p align="center">[Johanson Leeroy] - [109082500017]</p>

## Unguided 

### 1. [Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.]
#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_10](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul10/output/output-soal1.png)
[Program di atas dibuat untuk mencari berat terbesar dan berat terkecil dari sekumpulan data bilangan. Paling pertama merupakan deklarasi tipe data arr yang berupa array dengan kapasitas maksimal 1000 elemen bertipe integer. Kedua adalah Function isiData yang digunakan untuk mengisi data ke dalam array dengan menggunakan struktur kontrol perulangan for dan proses input menggunakan fmt.Scan. Ketiga adalah Function terbesar yang digunakan untuk mencari nilai terbesar pada array dengan cara membandingkan setiap elemen array menggunakan struktur kontrol if. Nilai terbesar awal diambil dari elemen pertama array, kemudian akan diperbarui jika ditemukan nilai yang lebih besar. Keempat adalah Function terkecil yang digunakan untuk mencari nilai terkecil pada array dengan cara membandingkan setiap elemen array menggunakan struktur kontrol if. Nilai terkecil awal juga diambil dari elemen pertama array, lalu diperbarui jika ditemukan nilai yang lebih kecil. Terakhir merupakan function utama main yang berisi deklarasi variabel, scan jumlah data dari user, pemanggilan function isiData untuk menginput data array, serta pemanggilan function terbesar dan terkecil untuk menampilkan hasil berat terbesar dan berat terkecil.]

### 2. [Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.]

#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_10](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul10/output/output-soal2.png)
[Program di atas dibuat untuk menghitung total berat ikan pada setiap wadah serta mencari rata-rata berat seluruh wadah. Paling pertama merupakan deklarasi tipe data ARR yang berupa array dengan kapasitas maksimal 1000 elemen bertipe float64. Kedua adalah Function isiIkan yang digunakan untuk mengisi data berat ikan ke dalam array dengan menggunakan struktur kontrol perulangan for dan proses input menggunakan fmt.Scan. Ketiga adalah Function rerata yang digunakan untuk menghitung total berat ikan pada setiap wadah dan rata-rata berat seluruh wadah. Pada function ini digunakan variabel jumlah untuk menampung total sementara setiap wadah, kemudian menggunakan struktur kontrol if untuk memeriksa apakah jumlah data ikan sudah memenuhi kapasitas satu wadah berdasarkan nilai y. Jika sudah memenuhi, maka total berat wadah akan ditampilkan, disimpan ke variabel Total, lalu jumlah direset menjadi 0 untuk perhitungan wadah berikutnya. Setelah seluruh data selesai diproses, program menghitung rata-rata berat wadah dengan rumus total seluruh berat dibagi jumlah wadah. Terakhir merupakan function utama main yang berisi deklarasi variabel, scan input jumlah ikan dan kapasitas tiap wadah dari user, proses perhitungan jumlah wadah, pemanggilan function isiIkan untuk menginput data berat ikan, serta pemanggilan function rerata untuk menampilkan total berat setiap wadah dan hasil rata-rata berat seluruh wadah.]

### 3. [Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.]

#### soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {

	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 0; i < n; i++ {

		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {

	var total float64

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {

	var data arrBalita
	var n int
	var bMin, bMax, rata float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &bMin, &bMax)

	rata = rerata(data, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_10](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul10/output/output-soal3.png)
[Program di atas dibuat untuk menghitung berat balita minimum, berat balita maksimum, dan rerata berat balita dari sejumlah data yang diinputkan. Paling pertama merupakan deklarasi tipe data arrBalita yang berupa array dengan kapasitas maksimal 100 elemen bertipe float64. Kedua adalah Function hitungMinMax yang digunakan untuk mencari berat minimum dan maksimum dari data balita. Pada function ini, nilai awal minimum dan maksimum diambil dari elemen pertama array, kemudian program menggunakan struktur kontrol perulangan for dan struktur kontrol if untuk membandingkan setiap data berat balita. Jika ditemukan nilai yang lebih kecil maka nilai minimum diperbarui, dan jika ditemukan nilai yang lebih besar maka nilai maksimum diperbarui. Function ini menggunakan parameter pointer bMin dan bMax agar hasil perhitungan dapat langsung disimpan ke variabel asli. Ketiga adalah Function rerata yang digunakan untuk menghitung rata-rata berat balita dengan cara menjumlahkan seluruh data berat menggunakan struktur kontrol perulangan for, kemudian total berat dibagi dengan jumlah data balita. Function ini mengembalikan hasil berupa nilai float64. Terakhir merupakan function utama main yang berisi deklarasi variabel, scan jumlah data balita dari user, proses input data berat balita menggunakan perulangan, pemanggilan function hitungMinMax untuk mencari berat minimum dan maksimum, serta pemanggilan function rerata untuk menghitung rata-rata berat balita dan menampilkan seluruh hasil perhitungan ke layar.]

