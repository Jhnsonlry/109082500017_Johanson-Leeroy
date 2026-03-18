# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Johanson Leeroy] - [109082500017]</p>

## Unguided 

### 1. [Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?]
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
![Screenshot Output Unguided 1_2](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul2/output/output-soal1.png)
[
	1. Program meminta user untuk memasukan string kedalam variabel satu, dua, tiga. 
	2. Program akan mengeluarkan output string pada variabel satu, dua, tiga secara berurutan.
	3. Program memindahkan string pada variabel satu ke variabel temp, dua ke satu, tiga ke dua dan terakhir temp ke tiga.
	4. Output akhir berupa string terakhir dari variabel satu, dua, tiga. ]

### 2. [Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.]

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
![Screenshot Output Unguided 2_2](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul2/output/output-soal2.png)
[
	1. Program melakukan perulangan percobaan sebanyak 5 kali
	2. Program meminta user untuk memasukan warna dari setiap gelas
	3. Jika warna gelas sesuai dengan yang diminta (gelas1 = merah, gelas2 = kuning, gelas3 = hijau, dst.) maka percobaannya berhasil dan variabel count bertambah 1
	4. Jika variabel count = 5, berarti seluruh percobaan telah berhasil(BERHASIL = true), tetapi jika count != 5, berarti ada percobaan yang gagal (Berhasil = false)]

### 3. [PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut!Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.]

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
![Screenshot Output Unguided 3_2](https://github.com/Jhnsonlry/109082500017_Johanson-Leeroy/blob/main/modul2/output/output-soal3.png)
[
	1. Program meminta user untuk memberikan input pada variabel beratParsel (berat dalam gram)
	2. Program akan melakukan operasi aritmatika untuk menghitung beratSisa (mod 1000), menghitung beratParsel terbaru (beratParsel - beratSisa), dan mengubah beratParsel dari gram ke kilo gram (kg)
	3. Program mengeluarkan output berupa keterangan dari Detail Berat parsel
	4. Program melakukan operasi aritmatika menghitung biaya, yaitu Rp. 10000/kg
	5. Jika berat sisa >= 500, maka biaya tambahannya adalah Rp. 5/g beratSisa. Selain itu biaya tambahannya adalah Rp. 15/g beratSisa
	6. Program mengeluarkan output berupa keterangan dari Detai Biaya
	7. Jika Berat Parsel > 10kg maka tidak akan dikenakan biayaTambahan
	8. Terakhir program mengeluarkan output berupa Biaya Akhir yang harus dibayarkan]

