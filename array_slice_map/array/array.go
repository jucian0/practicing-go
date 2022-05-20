package main

import "fmt"

func main() {
	var notas [3]float64

	fmt.Println(notas)

	notas[0] = 9.5

	fmt.Println(notas)

	for i, v := range notas {
		fmt.Printf("Nota %d: %.2f\n", i, v)
	}
}
