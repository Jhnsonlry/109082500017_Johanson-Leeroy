# <h1 align="center">Laporan Praktikum Modul 12 - SEARCHING </h1>
<p align="center">[Johanson Leeroy] - [109082500017]</p>

## Unguided 

### 1. [Pada pemilihan ketua RT yang baru saja berlangsung, terdapat 20 calon ketua yang bertanding memperebutkan suara warga. Perhitungan suara dapat segera dilakukan karena warga cukup mengisi formulir dengan nomor dari calon ketua RT yang dipilihnya. Seperti biasa, selalu ada pengisian yang tidak tepat atau dengan nomor pilihan di luar yang tersedia, sehingga data juga harus divalidasi. Tugas Anda untuk membuat program mencari siapa yang memenangkan pemilihan ketua RT. Buatlah program pilkart yang akan membaca, memvalidasi, dan menghitung suara yang diberikan dalam pemilihan ketua RT tersebut.]

#### soal1.go

```go
package main

import "fmt"

const NMAX = 1000

type TabInt [NMAX]int

func sequentialSearch(A TabInt, n, x int) int {
	var i int
	i = 0
	for i < n && A[i] != x {
		i++
	}

	if i < n {
		return i
	}
	return -1
}

func hitungSuara(daftar *TabInt, nDaftar *int, suara *[21]int, total *int, sah *int) {
	var x, idx int

	for {
		fmt.Scan(&x)

		if x == 0 {
			*total++
			break
		}

		*total++

		if x >= 1 && x <= 20 {
			*sah++
			suara[x]++

			idx = sequentialSearch(*daftar, *nDaftar, x)
			if idx == -1 {
				daftar[*nDaftar] = x
				*nDaftar++
			}
		}
	}
}

func tampilkanHasil(daftar TabInt, nDaftar int, suara [21]int, total, sah int) {
	var i, j, temp int

	for i = 0; i < nDaftar-1; i++ {
		for j = i + 1; j < nDaftar; j++ {
			if daftar[i] > daftar[j] {
				temp = daftar[i]
				daftar[i] = daftar[j]
				daftar[j] = temp
			}
		}
	}

	fmt.Println("Suara masuk:", total-1)
	fmt.Println("Suara sah:", sah)

	for i = 0; i < nDaftar; i++ {
		fmt.Printf("%d: %d\n", daftar[i], suara[daftar[i]])
	}
}

func main() {
	var daftar TabInt
	var suara [21]int
	var nDaftar, total, sah int

	hitungSuara(&daftar, &nDaftar, &suara, &total, &sah)
	tampilkanHasil(daftar, nDaftar, suara, total, sah)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_12](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul12/output/output-soal1.png)
[Program di atas dibuat untuk menghitung hasil pemilihan ketua RT berdasarkan suara yang masuk. Pertama terdapat fungsi sequentialSearch yang berfungsi untuk mencari apakah nomor calon sudah terdapat dalam daftar calon yang memperoleh suara. Jika ditemukan fungsi akan mengembalikan indeks data, sedangkan jika tidak ditemukan akan mengembalikan nilai -1. Kedua terdapat prosedur hitungSuara yang berfungsi untuk membaca dan menghitung seluruh suara yang masuk hingga pengguna memasukkan angka 0 sebagai tanda selesai. Setiap suara yang valid (1–20) akan dihitung sebagai suara sah, jumlah suara calon akan ditambah, dan nomor calon akan dicatat ke dalam daftar jika belum pernah muncul sebelumnya. Ketiga terdapat prosedur tampilkanHasil yang berfungsi untuk mengurutkan daftar nomor calon yang memperoleh suara, kemudian menampilkan jumlah suara masuk, jumlah suara sah, serta jumlah suara yang diperoleh masing-masing calon. Terakhir terdapat fungsi main yang berisi deklarasi variabel dan pemanggilan prosedur hitungSuara serta tampilkanHasil. Output akhir berupa jumlah suara masuk, jumlah suara sah, dan perolehan suara setiap calon yang mendapatkan suara.]

### 2. [Berdasarkan program sebelumnya, buat program pilkart yang mencari siapa pemenang pemilihan ketua RT. Sekaligus juga ditentukan bahwa wakil ketua RT adalah calon yang mendapatkan suara terbanyak kedua. Jika beberapa calon mendapatkan suara terbanyak yang sama, ketua terpilih adalah dengan nomor peserta yang paling kecil dan wakilnya dengan nomor peserta terkecil berikutnya.]

#### soal2.go

```go
package main

import "fmt"

const NMAX = 1000

type TabInt [NMAX]int

func sequentialSearch(A TabInt, n, x int) int {
	var i int
	i = 0

	for i < n && A[i] != x {
		i++
	}

	if i < n {
		return i
	}
	return -1
}

func bacaSuara(daftar *TabInt, nDaftar *int, suara *[21]int, total *int, sah *int) {
	var x, idx int

	for {
		fmt.Scan(&x)

		*total++

		if x == 0 {
			break
		}

		if x >= 1 && x <= 20 {
			*sah++
			suara[x]++

			idx = sequentialSearch(*daftar, *nDaftar, x)
			if idx == -1 {
				daftar[*nDaftar] = x
				*nDaftar++
			}
		}
	}
}

func cariKetuaNWakil(suara [21]int, ketua, wakil *int) {
	var maks1, maks2 int
	var i int

	for i = 1; i <= 20; i++ {
		if suara[i] > maks1 {
			maks1 = suara[i]
		}
	}

	for i = 1; i <= 20; i++ {
		if suara[i] == maks1 {
			*ketua = i
			break
		}
	}

	for i = *ketua + 1; i <= 20; i++ {
		if suara[i] == maks1 {
			*wakil = i
			return
		}
	}

	for i = 1; i <= 20; i++ {
		if suara[i] > maks2 && suara[i] < maks1 {
			maks2 = suara[i]
		}
	}

	for i = 1; i <= 20; i++ {
		if suara[i] == maks2 {
			*wakil = i
			break
		}
	}
}

func main() {
	var daftar TabInt
	var suara [21]int
	var nDaftar, total, sah int
	var ketua, wakil int

	bacaSuara(&daftar, &nDaftar, &suara, &total, &sah)
	cariKetuaNWakil(suara, &ketua, &wakil)

	fmt.Println("Suara masuk:", total-1)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_12](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul12/output/output-soal2.png)
[Program di atas dibuat untuk menentukan ketua dan wakil ketua RT berdasarkan hasil pemungutan suara. Pertama terdapat fungsi sequentialSearch yang berfungsi untuk mencari apakah nomor calon sudah terdapat dalam daftar calon yang memperoleh suara. Jika ditemukan fungsi akan mengembalikan indeks data, sedangkan jika tidak ditemukan akan mengembalikan nilai -1. Kedua terdapat prosedur bacaSuara yang berfungsi untuk membaca dan menghitung seluruh suara yang masuk hingga pengguna memasukkan angka 0 sebagai tanda selesai. Setiap suara yang valid (1–20) akan dihitung sebagai suara sah, jumlah suara calon akan ditambah, dan nomor calon akan dicatat ke dalam daftar jika belum pernah muncul sebelumnya. Ketiga terdapat prosedur cariKetuaNWakil yang berfungsi untuk menentukan ketua dan wakil ketua RT. Program terlebih dahulu mencari jumlah suara terbanyak, kemudian menetapkan calon pertama yang memiliki jumlah suara tersebut sebagai ketua. Jika terdapat calon lain dengan jumlah suara yang sama, maka calon tersebut akan menjadi wakil ketua. Jika tidak ada, wakil ketua ditentukan dari calon dengan jumlah suara terbanyak kedua. Terakhir terdapat fungsi main yang berisi deklarasi variabel dan pemanggilan prosedur bacaSuara serta cariKetuaNWakil. Output akhir berupa jumlah suara masuk, jumlah suara sah, nomor calon yang terpilih sebagai ketua RT, dan nomor calon yang terpilih sebagai wakil ketua RT.]

### 3. [Diberikan n data integer positif dalam keadaan terurut membesar dan sebuah integer lain k, apakah bilangan k tersebut ada dalam daftar bilangan yang diberikan? Jika ya, berikan indeksnya, jika tidak sebutkan "TIDAK ADA".]

#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_12](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul12/output/output-soal3.png)
[Program di atas dibuat untuk mencari posisi suatu bilangan pada kumpulan data yang telah terurut membesar. Pertama terdapat prosedur isiArray yang berfungsi untuk membaca dan menyimpan sejumlah n bilangan ke dalam array data. Data yang dimasukkan akan digunakan sebagai tempat pencarian bilangan yang dicari. Kedua terdapat fungsi posisi yang berfungsi untuk mencari letak bilangan k dalam array menggunakan metode binary search. Pencarian dilakukan dengan membandingkan nilai tengah array dengan nilai yang dicari. Jika nilai tengah sama dengan k, maka fungsi mengembalikan indeksnya. Jika nilai tengah lebih kecil dari k, pencarian dilanjutkan ke bagian kanan array, sedangkan jika lebih besar maka pencarian dilanjutkan ke bagian kiri array. Proses ini berlangsung hingga data ditemukan atau ruang pencarian habis. Terakhir terdapat fungsi main yang berisi deklarasi variabel dan pemanggilan prosedur isiArray serta fungsi posisi. Setelah pencarian selesai, program menggunakan struktur if-else untuk menentukan output. Jika fungsi posisi mengembalikan -1, program akan menampilkan "TIDAK ADA", sedangkan jika data ditemukan program akan menampilkan indeks posisi bilangan tersebut dalam array. Output akhir berupa indeks bilangan yang dicari atau tulisan "TIDAK ADA" jika bilangan tersebut tidak terdapat pada data.]

