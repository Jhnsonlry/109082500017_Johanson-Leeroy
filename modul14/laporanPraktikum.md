# <h1 align="center">Laporan Praktikum Modul 14 - SORTING </h1>
<p align="center">[Johanson Leeroy] - [109082500017]</p>

## Unguided 

### 1. [Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.]

#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_14](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul14/output/output-soal1.png)
[Program di atas dibuat untuk mengurutkan beberapa kumpulan data bilangan menggunakan metode Selection Sort. Paling pertama merupakan deklarasi konstanta NMAX dan DMAX yang digunakan untuk menentukan kapasitas maksimum array yang akan menyimpan data bilangan dan banyaknya kumpulan data yang dapat diproses oleh program. Kedua adalah function selectionSort yang digunakan untuk mengurutkan elemen-elemen array secara menaik (ascending) menggunakan algoritma Selection Sort. Function ini bekerja dengan mencari elemen terkecil pada bagian array yang belum terurut, kemudian menukarnya dengan elemen pada posisi saat ini. Proses tersebut dilakukan berulang menggunakan struktur kontrol perulangan bersarang (for) hingga seluruh data berada dalam urutan yang benar. Terakhir merupakan function main yang berisi deklarasi variabel untuk menyimpan banyaknya kumpulan data, ukuran masing-masing kumpulan data, array penyimpanan data, serta array sementara. Program terlebih dahulu menerima input jumlah kumpulan data (n), kemudian menggunakan perulangan for untuk membaca ukuran setiap kumpulan data (m) beserta elemen-elemennya. Setelah data dibaca, program melakukan pemanggilan function selectionSort untuk mengurutkan data tersebut, lalu menyimpannya ke dalam array dua dimensi data. Setelah seluruh kumpulan data selesai diproses, program menggunakan perulangan for untuk menampilkan kembali setiap kumpulan data yang telah terurut ke layar dengan format yang rapi.]

### 2. [Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.]

#### soal2.go

```go
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
	var n, m, x int
	var i, j, k int

	var hasil [DMAX][NMAX]int
	var ukuran [DMAX]int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		var ganjil, genap [NMAX]int
		var nGanjil, nGenap int

		fmt.Scan(&m)

		for j = 0; j < m; j++ {
			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil[nGanjil] = x
				nGanjil++
			} else {
				genap[nGenap] = x
				nGenap++
			}
		}

		selectionSort(&ganjil, nGanjil)
		selectionSort(&genap, nGenap)

		k = 0

		for j = 0; j < nGanjil; j++ {
			hasil[i][k] = ganjil[j]
			k++
		}

		for j = nGenap - 1; j >= 0; j-- {
			hasil[i][k] = genap[j]
			k++
		}

		ukuran[i] = k
	}

	for i = 0; i < n; i++ {
		for j = 0; j < ukuran[i]; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(hasil[i][j])
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_14](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul14/output/output-soal2.png)
[Program di atas dibuat untuk mengelompokkan dan mengurutkan sekumpulan bilangan berdasarkan paritasnya (ganjil dan genap). Paling pertama merupakan deklarasi konstanta NMAX dan DMAX yang digunakan untuk menentukan kapasitas maksimum array yang menyimpan data bilangan dan jumlah kumpulan data yang dapat diproses oleh program. Kedua adalah function selectionSort yang digunakan untuk mengurutkan elemen-elemen array menggunakan algoritma Selection Sort secara menaik (ascending). Function ini bekerja dengan mencari elemen terkecil pada bagian array yang belum terurut, kemudian menukarnya dengan elemen pada posisi saat ini. Proses tersebut dilakukan berulang menggunakan struktur kontrol perulangan bersarang hingga seluruh elemen array tersusun secara terurut. Terakhir merupakan function main yang berisi deklarasi variabel untuk menyimpan jumlah kumpulan data, ukuran data, elemen yang dibaca, serta array hasil. Program terlebih dahulu menerima input banyaknya kumpulan data (n), kemudian menggunakan perulangan for untuk membaca setiap kumpulan data. Setiap bilangan yang diinput akan dipisahkan ke dalam dua array berbeda, yaitu array ganjil untuk bilangan ganjil dan array genap untuk bilangan genap menggunakan struktur kontrol if-else. Setelah proses pemisahan selesai, program melakukan pemanggilan function selectionSort untuk mengurutkan bilangan ganjil dan genap secara menaik. Selanjutnya, bilangan ganjil yang telah terurut disalin ke array hasil dalam urutan menaik, sedangkan bilangan genap disalin dalam urutan terbalik sehingga menghasilkan urutan menurun (descending). Dengan demikian, hasil akhir yang disimpan terdiri dari seluruh bilangan ganjil yang terurut menaik diikuti oleh seluruh bilangan genap yang terurut menurun. Pada bagian akhir, program menggunakan struktur kontrol perulangan for untuk menampilkan setiap kumpulan data yang telah diproses, sehingga output yang dihasilkan berupa deretan bilangan ganjil terurut menaik diikuti bilangan genap terurut menurun untuk setiap kumpulan data yang diberikan.]

### 3. [Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.]

#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_14](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul14/output/output-soal3.png)
[Program di atas dibuat untuk mengurutkan sekumpulan bilangan dan menentukan apakah jarak antar data yang telah diurutkan bersifat tetap atau tidak. Paling pertama merupakan deklarasi konstanta NMAX yang digunakan untuk menentukan kapasitas maksimum array yang dapat menyimpan data bilangan. Kedua adalah function insertionSort yang digunakan untuk mengurutkan elemen-elemen array secara menaik (ascending) menggunakan algoritma Insertion Sort. Function ini bekerja dengan mengambil satu elemen pada setiap iterasi, kemudian membandingkannya dengan elemen-elemen sebelumnya dan menggesernya ke posisi yang sesuai hingga seluruh data tersusun secara terurut. Proses ini dilakukan menggunakan struktur kontrol perulangan bersarang (for). Terakhir merupakan function main yang berisi deklarasi array penyimpanan data, variabel penghitung jumlah data, variabel input, serta beberapa variabel bantu. Program menerima input bilangan dari user menggunakan perulangan for selama nilai yang dimasukkan bernilai non-negatif. Jika ditemukan bilangan negatif, proses input dihentikan. Setelah seluruh data berhasil dibaca, program melakukan pemanggilan function insertionSort untuk mengurutkan data dari yang terkecil hingga terbesar. Selanjutnya, program menampilkan seluruh data yang telah terurut menggunakan struktur kontrol perulangan for. Kemudian program melakukan pemeriksaan terhadap selisih antar elemen yang berurutan. Jika jumlah data kurang dari atau sama dengan satu, program langsung menampilkan bahwa data memiliki jarak 0. Jika jumlah data lebih dari satu, program menghitung selisih antara dua elemen pertama sebagai jarak acuan, lalu menggunakan perulangan dan struktur kontrol if untuk membandingkan selisih setiap pasangan elemen berikutnya. Jika seluruh selisih bernilai sama, maka program menyimpulkan bahwa data memiliki jarak tetap dan menampilkan nilai jaraknya. Sebaliknya, jika ditemukan selisih yang berbeda, program menyatakan bahwa data memiliki jarak yang tidak tetap. Dengan demikian, program dapat menentukan apakah data yang telah diurutkan membentuk pola seperti barisan aritmetika atau tidak.]

### 4. [Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan.]

#### soal4.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_14]()
[Program di atas dibuat untuk mengelola data buku dalam sebuah perpustakaan. Paling pertama merupakan deklarasi struct Buku yang digunakan untuk menyimpan informasi setiap buku, seperti ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Selain itu, dideklarasikan type DaftarBuku sebagai array yang berisi kumpulan data buku dengan kapasitas maksimum sebanyak nMax buku. Kedua adalah prosedur DaftarkanBuku yang digunakan untuk mengisi data buku ke dalam array pustaka. Prosedur ini menggunakan struktur kontrol perulangan for untuk membaca data setiap buku dari input dan menyimpannya ke dalam elemen array yang sesuai. Ketiga adalah prosedur CetakTerfavorit yang digunakan untuk mencari dan menampilkan buku dengan rating tertinggi. Prosedur ini melakukan penelusuran seluruh data buku menggunakan perulangan for, kemudian membandingkan nilai rating setiap buku dengan rating maksimum yang telah ditemukan sebelumnya. Setelah proses selesai, informasi buku dengan rating tertinggi akan dicetak ke layar. Keempat adalah prosedur UrutBuku yang digunakan untuk mengurutkan data buku berdasarkan rating secara menurun (descending) menggunakan algoritma Insertion Sort. Proses pengurutan dilakukan dengan membandingkan rating setiap buku dan memindahkan posisi data hingga seluruh buku tersusun dari rating tertinggi ke rating terendah. Kelima adalah prosedur Cetak5Terbaru yang digunakan untuk menampilkan maksimal lima buku teratas dari data yang telah diurutkan. Jika jumlah buku kurang dari lima, maka seluruh buku akan ditampilkan. Prosedur ini menggunakan struktur kontrol if-else untuk menentukan batas pencetakan dan perulangan for untuk menampilkan judul buku satu per satu. Terakhir adalah prosedur CariBuku yang digunakan untuk mencari buku berdasarkan nilai rating tertentu. Prosedur ini menerapkan algoritma Binary Search pada data buku yang telah terurut berdasarkan rating. Pencarian dilakukan dengan membandingkan rating yang dicari dengan rating pada posisi tengah array, kemudian mempersempit ruang pencarian hingga data ditemukan atau seluruh kemungkinan telah diperiksa. Jika ditemukan, prosedur akan menampilkan seluruh informasi buku tersebut. Sebaliknya, jika tidak ditemukan, program akan menampilkan pesan bahwa tidak ada buku dengan rating yang dicari. Dengan demikian, kumpulan prosedur tersebut dapat digunakan untuk melakukan pengisian data, pengurutan, pencarian, dan penampilan informasi buku dalam sistem perpustakaan.]
