package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)
	fmt.Println("&&", true && true)

	fmt.Println("||", true || false)
	fmt.Println("!", !true)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas", d1 == d2)
	fmt.Println("Datas", d1 != d2)
	fmt.Println("Datas", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Nome 1"}
	p2 := Pessoa{"Nome 2"}

	fmt.Println("Pessoas:", p1 == p2)
}
