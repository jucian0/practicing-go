package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	b := 3

	fmt.Println("Soma:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Módulo:", a%b)

	//bitwise
	fmt.Println("E:", a&b)
	fmt.Println("OU:", a|b)
	fmt.Println("XOR:", a^b)
	fmt.Println("Complemento:", ^a)
	fmt.Println("Left Shift:", a<<2)
	fmt.Println("Right Shift:", a>>2)

	c := 3.0
	d := 2.0

	fmt.Println("Maior:", math.Max(float64(a), float64(b)))
	fmt.Println("Menor:", math.Min(c, d))
	fmt.Println("Exponenciação:", math.Pow(c, d))
}
