# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Johanson Leeroy] - [109082500017]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[
	1. Program meminta user untuk memasukan string kedalam variabel satu, dua, tiga. 
	2. Program akan mengeluarkan output string pada variabel satu, dua, tiga secara berurutan.
	3. Program memindahkan string pada variabel satu ke variabel temp, dua ke satu, tiga ke dua dan terakhir temp ke tiga.
	4. Output akhir berupa string terakhir dari variabel satu, dua, tiga. ]

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var gelas1, gelas2, gelas3, gelas4 string
	var percobaan, count int
	var cek bool
	count = 0
	cek = false

	for percobaan = 1; percobaan <= 5; percobaan++ {
		fmt.Print("Perobaan", percobaan, ": ")
		fmt.Scanln(&gelas1, &gelas2, &gelas3, &gelas4)
		if gelas1 == "merah" && gelas2 == "kuning" && gelas3 == "hijau" && gelas4 == "ungu" {
			count++
		}
	}
	if count == 5 {
		cek = true
	}
	fmt.Println("BERHASIL:", cek)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_2](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[
	1. Program melakukan perulangan percobaan sebanyak 5 kali
	2. Program meminta user untuk memasukan warna dari setiap gelas
	3. Jika warna gelas sesuai dengan yang diminta (gelas1 = merah, gelas2 = kuning, gelas3 = hijau, dst.) maka percobaannya berhasil dan variabel count bertambah 1
	4. Jika variabel count = 5, berarti seluruh percobaan telah berhasil(BERHASIL = true), tetapi jika count != 5, berarti ada percobaan yang gagal (Berhasil = false)]

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var beratParsel, beratSisa, beratKG, biaya, biayaTambahan, biayaTotal int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	beratSisa = beratParsel % 1000
	beratParsel = beratParsel - beratSisa
	beratKG = beratParsel / 1000

	fmt.Println("Detail Berat:", beratKG, "kg +", beratSisa, "gr")

	biaya = 10000 * beratKG

	if beratSisa >= 500 {
		biayaTambahan = beratSisa * 5
	} else {
		biayaTambahan = beratSisa * 15
	}

	fmt.Println("Detail Biaya: Rp.", biaya, "+ Rp.", biayaTambahan)

	if beratKG < 10 {
		biayaTotal = biaya + biayaTambahan
	} else {
		biayaTotal = biaya
	}

	fmt.Println("Biaya Total: Rp.", biayaTotal)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_2](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[
	1. Program meminta user untuk memberikan input pada variabel beratParsel (berat dalam gram)
	2. Program akan melakukan operasi aritmatika untuk menghitung beratSisa (mod 1000), menghitung beratParsel terbaru (beratParsel - beratSisa), dan mengubah beratParsel dari gram ke kilo gram (kg)
	3. Program mengeluarkan output berupa keterangan dari Detail Berat parsel
	4. Program melakukan operasi aritmatika menghitung biaya, yaitu Rp. 10000/kg
	5. Jika berat sisa >= 500, maka biaya tambahannya adalah Rp. 5/g beratSisa. Selain itu biaya tambahannya adalah Rp. 15/g beratSisa
	6. Program mengeluarkan output berupa keterangan dari Detai Biaya
	7. Jika Berat Parsel > 10kg maka tidak akan dikenakan biayaTambahan
	8. Terakhir program mengeluarkan output berupa Biaya Akhir yang harus dibayarkan]

