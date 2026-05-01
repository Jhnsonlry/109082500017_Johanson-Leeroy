# <h1 align="center">Laporan Praktikum Modul 9 - ARRAY </h1>
<p align="center">[Johanson Leeroy] - [109082500017]</p>

## Unguided 

### 1. [Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦)berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.]
#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type lingkaran struct {
	P1, P2, r float64
}

type KoorTitik struct {
	t1, t2 float64
}

func jarak(a, b, c, d float64) float64 {
	var jarak float64
	jarak = math.Sqrt(((a - c) * (a - c)) + ((b - d) * (b - d)))
	return jarak
}

func didalam(p1, p2, t1, t2, r1 float64) bool {
	var didalam bool
	didalam = false
	if r1 >= jarak(p1, p2, t1, t2) {
		didalam = true
	}
	return didalam
}

func main() {
	var lingkaran1, lingkaran2 lingkaran
	var titik KoorTitik

	fmt.Scan(&lingkaran1.P1, &lingkaran1.P2, &lingkaran1.r, &lingkaran2.P1, &lingkaran2.P2, &lingkaran2.r)
	fmt.Scan(&titik.t1, &titik.t2)

	if didalam(lingkaran1.P1, lingkaran1.P2, titik.t1, titik.t2, lingkaran1.r) == true &&
		didalam(lingkaran2.P1, lingkaran2.P2, titik.t1, titik.t2, lingkaran2.r) == true {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(lingkaran1.P1, lingkaran1.P2, titik.t1, titik.t2, lingkaran1.r) == true {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam(lingkaran2.P1, lingkaran2.P2, titik.t1, titik.t2, lingkaran2.r) == true {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_9](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul9/output/output-soal1.png)
[Program di atas dibuat untuk menentukan posisi suatu titik terhadap dua buah lingkaran. Paling pertama merupakan deklarasi struct lingkaran yang berisi titik pusat (P1, P2) dan jari-jari (r), serta struct KoorTitik yang digunakan untuk menyimpan koordinat sebuah titik (t1, t2). Kedua adalah function jarak yang berfungsi untuk menghitung jarak antara dua titik pada bidang menggunakan rumus jarak Euclidean Ketiga adalah function didalam yang digunakan untuk mengecek apakah suatu titik berada di dalam lingkaran atau tidak. Logikanya adalah membandingkan jarak antara titik dengan pusat lingkaran terhadap jari-jari lingkaran. Jika jarak tersebut lebih kecil atau sama dengan jari-jari, maka titik berada di dalam lingkaran, sehingga function mengembalikan nilai true, jika tidak maka false. Terakhir merupakan function main yang berisi deklarasi variabel untuk dua lingkaran dan satu titik, proses input data dari user menggunakan fmt.Scan, serta struktur kontrol if-else untuk menentukan posisi titik: apakah berada di dalam kedua lingkaran, hanya di lingkaran pertama, hanya di lingkaran kedua, atau di luar keduanya.]

### 2. [Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah programyang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:]

#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_9](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul9/output/output-soal2.png)
[Program di atas dibuat untuk mengolah dan menampilkan berbagai informasi dari sebuah array berisi 10 bilangan. Paling pertama merupakan deklarasi beberapa function untuk mencetak isi array, yaitu printArray untuk mencetak seluruh elemen array, printArrayGanjil untuk mencetak elemen dengan indeks ganjil, dan printArrayGenap untuk mencetak elemen dengan indeks genap. Semua function tersebut menggunakan struktur kontrol perulangan for dan seleksi if untuk menentukan elemen mana yang ditampilkan. Kedua adalah function printArrayKelipatanX yang digunakan untuk mencetak elemen array pada indeks yang merupakan kelipatan dari suatu bilangan X yang diinput oleh user. Function ini menggunakan perulangan dan kondisi if untuk mengecek apakah indeks habis dibagi X. Ketiga adalah function simpanganBaku yang digunakan untuk menghitung simpangan baku dari data dalam array. Perhitungannya dilakukan dengan mencari selisih setiap elemen terhadap rata-rata, kemudian dikuadratkan, dijumlahkan, dan dibagi (n-1) untuk mendapatkan nilai simpangan baku (meskipun dalam kode ini belum diakar, sehingga sebenarnya masih berupa varians). Terakhir merupakan function main yang berisi deklarasi array dengan nilai awal, slice untuk penghapusan, serta beberapa variabel bantu. Program kemudian menampilkan isi array dengan berbagai variasi menggunakan function yang telah dibuat, menerima input nilai X dari user, menghitung rata-rata array, menghitung simpangan baku, menghitung frekuensi kemunculan angka 4, serta melakukan simulasi penghapusan elemen indeks ke-0 dengan cara mencetak elemen array mulai dari indeks ke-1 menggunakan struktur kontrol perulangan.]

### 3. [Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.]

#### soal3.go

```go
package main

import "fmt"

func Pemenang(klub1, klub2 string, skor [2]int) []string {
	var menang []string
	var pertandingan int
	var i int = 0
	for {
		pertandingan++
		fmt.Print("Pertangdingan ", pertandingan, ": ")
		fmt.Scan(&skor[0], &skor[1])
		if skor[0] < 0 || skor[1] < 0 {
			break
		}
		if skor[0] > skor[1] {
			menang = append(menang, klub1)
			i++
		} else if skor[1] > skor[0] {
			menang = append(menang, klub2)
			i++
		} else {
			menang = append(menang, "Draw")
		}

	}
	return menang
}

func main() {
	var klub1, klub2 string
	var skor [2]int

	fmt.Print("Klub A : ")
	fmt.Scan(&klub1, "\n")
	fmt.Print("Klub B : ")
	fmt.Scan(&klub2, "\n")

	hasil := Pemenang(klub1, klub2, skor)
	banyakPertandingan := len(hasil)

	for i := 0; i < banyakPertandingan; i++ {
		fmt.Println("hasil", i+1, ":", hasil[i])
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_9](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul9/output/output-soal3.png)
[Program di atas dibuat untuk menentukan pemenang dari beberapa pertandingan antara dua klub berdasarkan skor yang diinput oleh user. Paling pertama merupakan deklarasi function Pemenang yang digunakan untuk memproses hasil pertandingan. Function ini menerima nama dua klub dan array skor, kemudian menggunakan struktur kontrol perulangan tak hingga (for) untuk membaca skor setiap pertandingan. Perulangan akan berhenti jika terdapat skor bernilai negatif. Di dalam function ini juga terdapat struktur kontrol if-else untuk membandingkan skor kedua klub, lalu menyimpan hasilnya ke dalam slice menang (berisi nama klub yang menang atau “Draw”). Kedua adalah proses penambahan data ke dalam slice menggunakan append, yang memungkinkan penyimpanan hasil pertandingan secara dinamis sesuai jumlah pertandingan yang diinput oleh user. Terakhir merupakan function main yang berisi deklarasi variabel untuk nama kedua klub dan skor, proses input nama klub menggunakan fmt.Scan, serta pemanggilan function Pemenang untuk mendapatkan seluruh hasil pertandingan. Setelah itu, program menggunakan perulangan for untuk menampilkan hasil setiap pertandingan satu per satu berdasarkan data yang tersimpan di dalam slice hasil.]

### 4. [Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.]

#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var huruf rune
	fmt.Print("Teks :")
	for *n < NMAX {
		fmt.Scanf(" %c", &huruf)
		if huruf == '.' {
			break
		}
		t[*n] = huruf
		*n++
	}

}
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}

}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var tAwal tabel = t
	var cek bool = true
	balikanArray(&t, n)
	for i := 0; i < n; i++ {
		if tAwal[i] != t[i] {
			cek = false
			break
		}
	}
	return cek
}
func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Reverse Teks : ")
	balikanArray(&tab, m)

	cetakArray(tab, m)

	Cekpalindrom := palindrom(tab, m)
	fmt.Println()
	fmt.Println("Palindrom ?", Cekpalindrom)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_9](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul9/output/output-soal4.png)
[Program di atas dibuat untuk mengolah teks dengan cara membalik urutan karakter dan memeriksa apakah teks tersebut merupakan palindrom. Paling pertama merupakan deklarasi konstanta NMAX sebagai batas maksimum jumlah karakter, serta type tabel yang berupa array bertipe rune untuk menyimpan karakter-karakter teks yang diinput oleh user. Kedua adalah function isiArray yang digunakan untuk membaca input karakter satu per satu dari user dan menyimpannya ke dalam array. Proses input dilakukan menggunakan struktur kontrol perulangan for, dan akan berhenti jika jumlah karakter telah mencapai batas maksimum atau jika user memasukkan tanda titik (.) sebagai penanda akhir input. Ketiga adalah function cetakArray yang berfungsi untuk menampilkan isi array karakter ke layar menggunakan perulangan, sehingga seluruh teks dapat dicetak kembali sesuai isi array. Keempat adalah function balikanArray yang digunakan untuk membalik urutan karakter dalam array. Function ini bekerja dengan cara menukar elemen pertama dengan elemen terakhir, elemen kedua dengan elemen kedua dari belakang, dan seterusnya menggunakan perulangan hingga mencapai setengah panjang array. Kelima adalah function palindrom yang digunakan untuk mengecek apakah teks merupakan palindrom atau tidak. Function ini bekerja dengan menyalin array awal, kemudian membalik array tersebut menggunakan pemanggilan function balikanArray, lalu membandingkan setiap karakter pada array asli dan array hasil pembalikan menggunakan struktur kontrol if. Jika seluruh karakter sama, maka function mengembalikan nilai true, sedangkan jika berbeda maka mengembalikan nilai false. Terakhir merupakan function main yang berisi deklarasi variabel array dan penghitung jumlah karakter, pemanggilan function untuk mengisi array, membalik teks, menampilkan hasil pembalikan, serta melakukan pengecekan palindrom melalui pemanggilan function palindrom, kemudian menampilkan hasilnya ke layar.]