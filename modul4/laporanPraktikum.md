# <h1 align="center">Laporan Praktikum Modul 4 - Prosedur </h1>
<p align="center">[Johanson Leeroy] - [109082500017]</p>

## Unguided 

### 1. [Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas?]
#### soal1.go

```go
package main

import "fmt"

func faktorial(f int, faktorial *int) {
	var i int
	*faktorial = 1
	for i = 1; i <= f; i++ {
		*faktorial = *faktorial * i
	}
}

func permutasi(p, q int, permuatsi *int) {
	var r, s int
	faktorial(p, &r)
	faktorial(p-q, &s)
	*permuatsi = r / s
}

func kombinasi(x, y int, kombinasi *int) {
	var h, i, j int
	faktorial(x, &h)
	faktorial(y, &i)
	faktorial(x-y, &j)
	*kombinasi = h / (i * j)
}

func main() {
	var a, b, c, d, prmt, comb int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c {
		permutasi(a, c, &prmt)
		kombinasi(a, c, &comb)
		fmt.Println(prmt, comb)
	} else {
		permutasi(c, a, &prmt)
		kombinasi(c, a, &comb)
		fmt.Println(prmt, comb)
	}

	if b >= d {
		permutasi(b, d, &prmt)
		kombinasi(b, d, &comb)
		fmt.Println(prmt, comb)
	} else {
		permutasi(d, b, &prmt)
		kombinasi(d, b, &comb)
		fmt.Println(prmt, comb)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_4](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul4/soal1.go)
[Program di atas dibuat untuk menghitung Permutasi dan Kombinasi. Paling pertama merupakan deklarasi Prosedur untuk menghitung faktorial dari sebuah bilangan yang menggunakan logika struktur kontrol perulangan. Kedua adalah Prosedur untuk menghitung permutasi dengan rumus p!/q! dan dengan cara melakukan pemanggilan Prosedur faktorial. Ketiga adalah Prosedur untuk menghitung Kombinasi dengan rumus x!/ (y! * (x - y)!) dan dengan cara melakukan pemanggilan Prosedur faktorial. Terakhir merupakan Fungsi utama yang berisi deklarasi variabel, scan input dari user, dan Struktur kontrol if-else untuk memastikan output tidak salah karena urutan angka yang tidak tepat serta melakukan pemanggilan Prosedur permutasi dan kombinasi]

### 2. [Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya.]

#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_4](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul4/soal2.go)
[Program di atas dibuat untuk menentukan pemenang dari kompetisi pemrograman tingkat nasional. Pertama terdapat prosedur hitungSkor yang berfungsi untuk meminta input waktu yang dibutuhkan untuk menyelesaikan setiap soal yang dilakukan oleh peserta, nantinya akan dapat diketahui berapa nomor yang berhasil dikerjakan beserta total waktunya. Kedua ada fungsi utama, berisi deklarasi variabel dan struktur perulangan yang diawali dengan pengimputan nama oleh user, dilanjutkan dengan pemanngilan prosedur hitungSkor dan struktur if-else. Struktur if-else tersebut dibuat untuk menentukan pemengang dari kompetisi. Kondisi pertama yang dicek adalah jumlahSoal < soal, maksudnya adalah jika jumlah soal yang diselesaikan peserta sebelumnya (pemenang sementara) lebih sedikit dari jumlah soal yang diselesaikan peserta saat ini, maka detail jumlah soal, nilai dan pemenang sebelumnya akan diupdate dengan yang saat ini. Jika kondisi yang sebelumnya tidak terpenuhi maka program akan masuk kekondisi kedua, yaitu jumlahSoal == soal, maksudnya adalah jika jumlah soal yang diselesaikan peserta sebelumnya sama (pemenang sementara) dengan peserta saat ini, maka pemenang akan ditentukan dari junlah skor (total waktu pengerjaan) yang lebih sedikit. Perulangan ini akan berhenti ketika nama yang diinputkan adalah "selesai" atau "Selesai". Output akhir berupa detail pemenang dari nama, jumlah soal yang dikerjakan, dan total waktunya.]

### 3. [ Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah ½n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1.]

#### soal3.go

```go
package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			fmt.Print(n, " ")
		} else {
			n = 3*n + 1
			fmt.Print(n, " ")
		}
	}
}

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	cetakDeret(bilangan)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_4](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul4/soal3.go)
[Program di atas dibuat untuk menghasilkan deret dari bilangan yang diinputkan. Pertama terdapat prosedur cetakDeret untuk melakukan, dalam prosedur ini terdapat struktur perulangan dengan kondisi n != 1. Deret akan dibentuk dari 2 kondisi if-else yaitu n == genap atau n == ganjil. Jika n == genap, maka operasi yang dilakukan adalah n = n / 2. Jika n == ganjil maka operasi yang dilakukan adalah n = 3*n + 1, disetiap kondisi if-else terdapat perintah untuk mencetak nilai n terbaru. Kedua terdapat fungsi utama dari program untuk deklarasi variabel, scan input dan pemanggilan prosedur cetakDeret. Output akhir berupa deret dari sebuah bilangan yang diinputkan oleh user]

