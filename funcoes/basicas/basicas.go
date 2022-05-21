package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func f2(p1 string) {
	fmt.Println("f2", p1)
}

func f3(p1, p2 string) string {
	return p1 + p2
}

func f4(p1, p2 string) (string, string) {
	return p1, p2
}

func main() {

	f1()
	f2("parametro")

	r := f3("parametro1", "parametro2")
	fmt.Println(r)

	r1, r2 := f4("parametro1", "parametro2")
	fmt.Println(r1, r2)

}
