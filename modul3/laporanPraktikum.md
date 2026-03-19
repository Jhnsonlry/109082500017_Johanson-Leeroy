# <h1 align="center">Laporan Praktikum Modul 3 - Fungsi </h1>
<p align="center">[Johanson Leeroy] - [109082500017]</p>

## Unguided 

### 1. [Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas?]
#### soal1.go

```go
package main

import "fmt"

func faktorial(f int) int {
	var faktorial, i int
	faktorial = 1
	for i = 1; i <= f; i++ {
		faktorial = faktorial * i
	}
	return faktorial
}

func permutasi(p, q int) int {
	var permutasi int
	permutasi = faktorial(p) / faktorial(p-q)
	return permutasi
}

func kombinasi(x, y int) int {
	var kombinasi int
	kombinasi = faktorial(x) / (faktorial(y) * faktorial((x - y)))
	return kombinasi
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c {
		fmt.Println(permutasi(a, c), kombinasi(a, c))
	} else {
		fmt.Println(permutasi(c, a), kombinasi(c, a))
	}

	if b >= d {
		fmt.Println(permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Println(permutasi(d, b), kombinasi(d, b))
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_3](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul3/output/output-soal1.png)
[Program di atas dibuat untuk menghitung Permutasi dan Kombinasi. Paling pertama merupakan deklarasi Function untuk menghitung faktorial dari sebuah bilangan yang menggunakan logika struktur kontrol perulangan. Kedua adalah Function untuk menghitung permutasi dengan rumus p!/q! dan dengan cara melakukan pemanggilan function faktorial. Ketiga adalah Function untuk menghitung Kombinasi dengan rumus x!/ (y! * (x - y)!) dan dengan cara melakukan pemanggilan function faktorial. Terakhir merupakan function utama yang berisi deklarasi variabel, scan input dari user, dan Struktur kontrol if-else untuk memastikan output tidak salah karena urutan angka yang tidak tepat serta melakukan pemanggilan function permutasi dan kombinasi]

### 2. [Diberikan tiga buah fungsi matematika yaitu 𝑓 (𝑥) = 𝑥^2, 𝑔 (𝑥) = 𝑥 − 2 dan ℎ (𝑥) = 𝑥 +1. Fungsi komposisi (𝑓𝑜𝑔𝑜ℎ)(𝑥) artinya adalah 𝑓(𝑔(ℎ(𝑥))). Tuliskan 𝑓(𝑥), 𝑔(𝑥) dan ℎ(𝑥)dalam bentuk function.]

#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	var f int
	f = x * x
	return f
}
func g(x int) int {
	var g int
	g = x - 2
	return g
}
func h(x int) int {
	var h int
	h = x + 1
	return h
}

func main() {
	var a, b, c, fogoh, gohof, hofog int
	fmt.Scan(&a, &b, &c)

	fogoh = f(g(h(a)))
	gohof = g(h(f(b)))
	hofog = h(f(g(c)))

	fmt.Println(fogoh)
	fmt.Println(gohof)
	fmt.Println(hofog)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_3](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul3/output/output-soal2.png)
[Program di atas dibuat untuk melakukan opearasi fungsi komposisi. Pertama diawali dengan 3 function, yaitu f, g, dan h setiap fungsi memiliki operasi matematikanya sendiri. Kedua ada function utama yang merupakan deklarasi variabel, scan input dari user, pemanggilan function f,g,h untuk menghitung komposisi fungsi ( contoh : fungsi komposisi f o g o h, dihitung dengan pemanggilan fungsi f(g(h(a))) ). Terakhir merupakan output hasil dari masing masing fungsi komposisi]

### 3. [ Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦)berdasarkan dua lingkaran tersebut.]

#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

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
	var p1, p2, r1, pp1, pp2, r2, t1, t2 float64

	fmt.Scan(&p1, &p2, &r1, &pp1, &pp2, &r2, &t1, &t2)

	if didalam(p1, p2, t1, t2, r1) == true && didalam(pp1, pp2, t1, t2, r2) == true {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(p1, p2, t1, t2, r1) == true {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam(pp1, pp2, t1, t2, r2) == true {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_3](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul3/output/output-soal3.png)
[Program di atas dibuat untuk mengetahui apakah sebuat titik sembarang berada dalam lingkaran 1 atau 2 atau tidak sama sekali. Pertama merupakan deklarasi function jarak untuk menghitung jarak dari titik pusat lingkaran 1 dan 2 ke titik sembarang dengan menggunakan rumus 𝑗𝑎𝑟𝑎𝑘 = √(𝑎 − 𝑐)^2 + (𝑏 − 𝑑)^2. Kedua merupakan deklarasi function didalam untuk memeriksa apakah titik sembarang berada dalam lingkaran atau tidak dengan membandingkan panjang jari-jari dan jarak titik sembarang, jika jari-jari >= jarak maka titik berada di dalam lingkaran. Ketiga merupakan function utama yang berisi deklarasi variabel (p1,p2,r1 merupakan variabel untuk lingkaran 1, pp1,pp2,r2 merupakan variabel untuk lingkaran 2, t1,t2 untuk variabel koordinat titik sembarang) dan scan input dari user. Terakhir adalah struktur kontrol if-else untuk menentukan output, apakah titik berada di dalam lingkaran 1, lingkaran 2, kedua lingkaran atau tidak sama sekali]

