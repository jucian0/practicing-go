package main

import "fmt"

func mainSlice() {
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, v := range numeros {
		fmt.Printf("%d: %d\n", i, v)
	}
}
