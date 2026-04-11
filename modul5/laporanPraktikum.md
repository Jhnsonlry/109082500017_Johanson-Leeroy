# <h1 align="center">Laporan Praktikum Modul 5 - Rekursif </h1>
<p align="center">[Johanson Leeroy] - [109082500017]</p>

## Unguided 

### 1. [Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan 𝑆𝑛 = 𝑆𝑛−1 + 𝑆𝑛−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.]
#### soal1.go

```go
package main

import "fmt"

func Fibonanci(n, x1, x2, x3 int) {

	if n > 0 {
		x1 = x2
		x2 = x3
		x3 = x2 + x1
		fmt.Print(x2, " ")
		Fibonanci(n-1, x1, x2, x3)
	}
}

func main() {
	var x1, x2, x3 int
	x1 = 0
	x2 = 0
	x3 = 1
	fmt.Print(x2, " ")
	Fibonanci(10, x1, x2, x3)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_5](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul5/output/output-soal1.png)
[Program di atas dibuat untuk mencetak deret bilangan Fibonacci. Pertama merupakan deklarasi Prosedur Fibonanci untuk menghitung dan menampilkan angka-angka dalam deret yang menggunakan logika struktur kontrol percabangan if-else sebagai base case rekursif agar proses berhenti saat jumlah deret yang diinginkan sudah terpenuhi. Kedua adalah proses pergeseran nilai variabel di mana nilai terdahulu disimpan ke dalam variabel pendukung agar dapat dijumlahkan untuk menghasilkan angka berikutnya sesuai dengan deret matematika Fibonacci. Ketiga adalah pemanggilan kembali Prosedur Fibonanci di dalam dirinya sendiri (rekursif) dengan memperbarui nilai pada variabel x1,x2,x3 yang telah dihitung serta mengurangi nilai pada variabel n agar perulangan terus berjalan hingga mencapai batas. Terakhir merupakan Fungsi utama yang berisi deklarasi variabel awal, inisialisasi nilai pertama dari variabel x1,x2,x3, serta pemanggilan Prosedur Fibonanci untuk memulai proses pengeluaran output angka sebanyak jumlah yang ditentukan.]

### 2. [Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.]

#### soal2.go

```go
package main

import "fmt"

func SegiTiga(n int, bintang string) {

	if n > 0 {
		bintang = bintang + "*"
		fmt.Println(bintang)

		SegiTiga(n-1, bintang)
	}
}

func main() {
	var bintang string
	var n int
	bintang = ""
	fmt.Scan(&n)
	SegiTiga(n, bintang)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_5](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul5/output/output-soal2.png)
[Program di atas dibuat untuk membuat pola segitiga siku-siku. Pertama merupakan deklarasi Prosedur SegiTiga untuk mengeluarkan output bintang ke layar yang menggunakan logika struktur kontrol percabangan if-else sebagai kondisi penghenti agar tidak terjadi infinite loop. Kedua adalah proses penambahan string "*" pada variabel bintang untuk membentuk pola baris yang semakin panjang secara bertahap. Ketiga adalah pemanggilan kembali Prosedur SegiTiga di dalam dirinya sendiri dengan mengirimkan nilai baris yang telah dikurangi dan string bintang yang telah diperbarui agar dapat melanjutkan pencetakan ke baris berikutnya. Terakhir merupakan Fungsi utama yang berisi deklarasi variabel untuk menyimpan input pengguna berupa jumlah baris yang diinginkan, serta pemanggilan Prosedur SegiTiga sebagai awal berjalannya proses rekursi tersebut.]

### 3. [Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.]

#### soal3.go

```go
package main

import "fmt"

func faktor(n, f int) {
	if f <= n {
		if n%f == 0 {
			fmt.Print(f, " ")
		}
		faktor(n, f+1)
	}
}

func main() {
	var n, f int
	f = 1
	fmt.Scan(&n)
	faktor(n, f)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_5](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul5/output/output-soal3.png)
[Program di atas dibuat untuk mencari dan menampilkan semua faktor pembagi dari sebuah bilangan. Paling pertama merupakan deklarasi Prosedur faktor untuk mengevaluasi setiap angka pembagi potensial yang menggunakan logika struktur kontrol percabangan if-else pertama sebagai pembatas agar proses rekursif berhenti ketika pembagi sudah melebihi bilangan input. Kedua adalah proses pengecekan kondisi dengan operator modulo untuk menentukan apakah angka tersebut merupakan pembagi habis dari bilangan input atau bukan, sehingga hanya angka yang valid yang akan dikeluarkan. Ketiga adalah pemanggilan kembali Prosedur faktor secara rekursif dengan meningkatkan nilai pembagi secara bertahap agar program dapat memeriksa angka selanjutnya hingga seluruh faktor didapatkan. Terakhir merupakan Fungsi utama yang berisi deklarasi variabel untuk menampung input dan nilai awal pembagi, proses pengambilan input dari pengguna, serta pemanggilan Prosedur faktor.]

### 4. [Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N]

#### soal4.go

```go
package main

import "fmt"

func kurang(n int) {
	if n >= 1 {
		fmt.Print(n, " ")
		kurang(n - 1)
	}
}

func tambah(n, m int) {
	if m <= n {
		fmt.Print(m, " ")
		tambah(n, m+1)

	}
}

func main() {
	var n, m int
	fmt.Scan(&n)
	kurang(n)
	m = 2
	tambah(n, m)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_5](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul5/output/output-soal4.png)
[Program di atas dibuat untuk menampilkan deret angka yang menurun kemudian naik kembali. Pertama merupakan deklarasi Prosedur kurang untuk mencetak angka secara mundur dari nilai input hingga angka satu yang menggunakan logika struktur kontrol percabangan if-else sebagai pembatas rekursif. Kedua adalah Prosedur tambah yang berfungsi untuk mencetak angka secara menaik mulai dari nilai 1 hingga kembali ke nilai input awal. Ketiga adalah pemanggilan kembali masing-masing Prosedur di dalam dirinya sendiri secara rekursif dengan memperbarui nilai parameter, baik dengan pengurangan maupun penambahan, agar urutan angka terbentuk dengan benar. Terakhir merupakan Fungsi utama yang berisi deklarasi variabel, proses pengambilan input dari pengguna, serta pengaturan urutan pemanggilan Prosedur kurang dan tambah untuk memastikan pola angka tercetak dengan benar sesuai alur yang diinginkan.]

### 5. [Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.]

#### soal5.go

```go
package main

import "fmt"

func ganjil(n, f int) {
	if f <= n {
		if f%2 != 0 {
			fmt.Print(f, " ")
		}
		ganjil(n, f+1)
	}
}

func main() {
	var n, f int
	f = 1
	fmt.Scan(&n)
	ganjil(n, f)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 5_5](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul5/output/output-soal5.png)
[Program di atas dibuat untuk menampilkan deret bilangan ganjil dari satu hingga N. Pertama merupakan deklarasi Prosedur ganjil untuk menelusuri setiap bilangan dalam rentang yang ditentukan yang menggunakan logika struktur kontrol percabangan if-else pertama sebagai base case agar proses pemanggilan berhenti ketika angka sudah melebihi nilai input. Kedua adalah proses pengecekan kondisi menggunakan operator modulo untuk mengecek apakah bilangan tersebut memiliki sisa bagi tidak sama dengan 0 ketika dibagi dengan 2, yang menandakan bahwa bilangan tersebut adalah ganjil. Ketiga adalah pemanggilan kembali Prosedur ganjil secara rekursif dengan meningkatkan nilai variabel f sebesar satu angka di setiap tahapnya agar seluruh rangkaian bilangan dapat diperiksa secara berurutan. Terakhir merupakan Fungsi utama yang berisi deklarasi variabel untuk menyimpan batas angka dan nilai awal, proses pengambilan input dari pengguna, serta pemanggilan Prosedur ganjil untuk memulai eksekusi program pencetakan deret bilangan tersebut.]

### 6. [Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.]

#### soal6.go

```go
package main

import "fmt"

func pangkat(hasil, x, y int) {
	if y > 0 {
		hasil *= x
		pangkat(hasil, x, y-1)
	} else {
		fmt.Print(hasil)
	}
}

func main() {
	var hasil, x, y int
	hasil = 1
	fmt.Scan(&x, &y)
	pangkat(hasil, x, y)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 6_5](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul5/output/output-soal6.png)
[Program di atas dibuat untuk menghitung hasil pangka.Ppertama merupakan deklarasi Prosedur pangkat untuk melakukan operasi perkalian berulang yang menggunakan logika struktur kontrol percabangan if sebagai penentu apakah proses perkalian masih harus berlanjut berdasarkan nilai variabel y yang tersisa. Kedua adalah proses akumulasi nilai di mana variabel hasil dikalikan dengan bilangan variabel x secara terus-menerus, sementara nilai variabel y dikurangi satu tiap kali pemanggilan agar mencapai kondisi akhir atau base case. Ketiga adalah pemanggilan kembali Prosedur pangkat secara rekursif dengan membawa nilai hasil perkalian terbaru dan variabel y yang telah diperbarui, atau mencetak hasil akhir ke layar jika variabel y telah mencapai angka nol. Terakhir merupakan Fungsi utama yang berisi deklarasi variabel untuk menampung variabel x, variabel y, dan nilai awal hasil, proses pengambilan input dari pengguna, serta pemanggilan Prosedur pangkat untuk memulai perhitungan sesuai dengan angka yang dimasukkan.]
