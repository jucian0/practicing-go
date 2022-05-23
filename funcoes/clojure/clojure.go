package main

import "fmt"

func clojure() func() {
	x := 10

	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}

func main() {

	x := 20

	fmt.Println(x)

	imprimeX := clojure()

	imprimeX()

}
