package main

import "fmt"

func main() {
	a1 := make([]int, 35, 100)
	a2 := append(a1, 1, 2, 3, 4, 5)
	fmt.Println(a1)

	a1[30] = 12

	fmt.Println(a1, a2)
}
