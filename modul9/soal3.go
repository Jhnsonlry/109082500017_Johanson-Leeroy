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
