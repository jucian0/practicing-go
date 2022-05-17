package main

import "fmt"

func imprimirResultado(nota float64) string {
	if nota >= 7 {
		return "Aprovado com nota " + fmt.Sprintf("%.2f", nota)
	} else {
		return "Reprovado com nota " + fmt.Sprintf("%.2f", nota)
	}
}

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 7 && nota < 9 {
		return "B"
	} else if nota >= 4 && nota < 7 {
		return "C"
	} else if nota >= 0 && nota < 4 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(imprimirResultado(7.9))
	fmt.Println(imprimirResultado(4.9))

	fmt.Println(notaParaConceito(7.9))
}
