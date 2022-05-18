package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i, v := range "abc" {
		fmt.Println(i, v)
	}

	for i, v := range []int{1, 2, 3, 4, 5} {
		fmt.Println(i, v)
	}

	i := 0

	for ; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("loop")
		//break
	}
}
