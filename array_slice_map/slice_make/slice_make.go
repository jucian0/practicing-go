package main

import "fmt"

func main() {
	s := make([]int, 35)
	s[30] = 12

	fmt.Println(s)

	s = make([]int, 35, 50)
	s[30] = 12

	fmt.Println(s)

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5)

	fmt.Println(s, len(s), cap(s))
}
