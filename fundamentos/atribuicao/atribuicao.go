package main

import "fmt"

func main() {
	var b = 3
	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i *= 3 // i = i * 3
	i /= 3 // i = i / 3
	i %= 3 // i = i % 3

	fmt.Println(i)

	x, y := 1, 2

	fmt.Println(x, y)
}
