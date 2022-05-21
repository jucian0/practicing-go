package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José":  11325.45,
		"Maria": 23456.78,
		"Pedro": 9876.54,
		"Ana":   54321.21,
	}

	funcsESalarios["Anna"] = 1234.56

	delete(funcsESalarios, "Maria")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
