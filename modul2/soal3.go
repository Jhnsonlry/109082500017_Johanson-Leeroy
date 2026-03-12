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
