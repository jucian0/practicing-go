package main

import "fmt"

func main() {
	//var aprovadors map[int]string

	aprovados := make(map[int]string)

	aprovados[1] = "Maria"
	aprovados[2] = "João"
	aprovados[3] = "Pedro"
	aprovados[4] = "Ana"
	aprovados[5] = "José"
	fmt.Println(aprovados)

	for id, name := range aprovados {
		fmt.Printf("%s (ID; %d)\n", name, id)
	}

	fmt.Println(aprovados[2])

	delete(aprovados, 3)
	fmt.Println(aprovados[3])
}
