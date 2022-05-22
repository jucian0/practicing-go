package main

func multiplicar(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, a, b int) int {
	return funcao(a, b)
}

func main() {
	resultado := exec(multiplicar, 2, 3)
	println(resultado)
}
