package main

import "fmt"

func main() {
	i := 1
	var p *int = nil

	p = &i // p recebe o endereço de i
	*p++   //remove referência e atribui o valor de i++
	i++

	// GO não tem aritmética de ponteiros
	//p++

	fmt.Println(i, *p)
}
