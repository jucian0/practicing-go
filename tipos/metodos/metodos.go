package main

import "strings"

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	p2 := pessoa{"Maria", "Souza"}

	p1.setNomeCompleto("João da Silva")
	p2.setNomeCompleto("Maria Pereira")

	println(p1.getNomeCompleto())
	println(p2.getNomeCompleto())
}
