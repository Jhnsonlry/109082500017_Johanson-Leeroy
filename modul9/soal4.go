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
