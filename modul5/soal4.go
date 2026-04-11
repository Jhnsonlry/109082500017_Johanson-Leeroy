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
