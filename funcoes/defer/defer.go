package main

import "fmt"

func obrtValorAprovado(valor int) int {
	defer fmt.Println("Valor aprovado:", valor)
	if valor > 5000 {
		fmt.Println("Valor muito alto...")
		return 5000
	}
	fmt.Println("Valor aprovado:", valor)

	return valor

}

func main() {

	fmt.Println("Início")
	obrtValorAprovado(5000)
	fmt.Println("Fim")

}
