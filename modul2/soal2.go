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
