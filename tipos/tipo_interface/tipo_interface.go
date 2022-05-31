package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{} = curso{"Go"}
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = 10
	fmt.Println(coisa2)
}
