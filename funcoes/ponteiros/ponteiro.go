package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é referenciado por um *
func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	//revisão: & usado para obeter a referência de um valor
	inc2(&n) // por referência
	fmt.Println(n)
}
