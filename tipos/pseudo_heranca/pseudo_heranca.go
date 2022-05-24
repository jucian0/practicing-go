package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Println(f)
	fmt.Println("A ferrari", f.nome, "esta com", f.velocidadeAtual, "km/h e o turbo está ligado?", f.turboLigado)

}
