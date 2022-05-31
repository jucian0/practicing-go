package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {

	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	fmt.Println(carro1.turboLigado)

	carro2 := ferrari{"F50", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro2.turboLigado)
}
