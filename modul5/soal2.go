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
