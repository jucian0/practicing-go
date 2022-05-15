package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3))

	var x int = 3
	fmt.Println("O tipo de x é", reflect.TypeOf(x))

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Números inteiros são sempre 64 bits
	var x64 int64 = 123123123123213123
	fmt.Println("O tipo de x64 é", reflect.TypeOf(x64))

	// Números reais são sempre 64 bits
	var xf float64 = 123.12
	fmt.Println("O tipo de xf é", reflect.TypeOf(xf))

	// Números complexos são sempre 128 bits
	var xc complex128 = 3 + 12i
	fmt.Println("O tipo de xc é", reflect.TypeOf(xc))

	// Números de ponto flutuante
	var x32 float32 = 12.12
	fmt.Println("O tipo de x32 é", reflect.TypeOf(x32))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é "
	s2 := "Ricardo"
	fmt.Println(s1 + s2)

	// string com concatenação
	fmt.Println(s1, s2)
	fmt.Println("Olá meu nome é", s2)

	// string com multiplas linhas
	str := `Olá
	meu
	nome
	é
	Ricardo`
	fmt.Println(str)

	// char
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}
