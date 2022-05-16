package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(a int) {
	fmt.Println(a)
}

func main() {

	resultado := somar(2, 3)
	imprimir(resultado)

}
