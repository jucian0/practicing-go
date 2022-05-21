package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   11325.45,
		},
		"J": {
			"José da Silva": 11325.45,
			"João da Silva": 23456.78,
		},
		"P": {
			"Pedro Junior": 9876.54,
			"Paula Souza":  54321.21,
		},
	}

	delete(funcsPorLetra, "G")

	for letra, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Println(letra, nome, salario)
		}
	}

}
