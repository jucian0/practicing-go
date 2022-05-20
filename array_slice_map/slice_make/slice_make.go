package main

import "fmt"

func main() {
	s := make([]int, 35)
	s[30] = 12

	fmt.Println(s)

	s = make([]int, 35, 50)
	s[30] = 12

	fmt.Println(s)
}
