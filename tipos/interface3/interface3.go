package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	acelerarAtual()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw struct{}

func (b bmw) ligarTurbo() {
	fmt.Println("Turbo ligado")
}

func (b bmw) acelerarAtual() {
	fmt.Println("Acelerando")
}

func main() {
	var e esportivoLuxuoso = bmw{}
	e.ligarTurbo()
	e.acelerarAtual()

	var l esportivoLuxuoso = bmw{}
	l.acelerarAtual()
}
