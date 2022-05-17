package main

import "fmt"

func oberterResultado(nota float64) string {
	if nota >= 7 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(oberterResultado(7.9))
}
