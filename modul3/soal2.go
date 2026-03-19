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
