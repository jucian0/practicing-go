package main

import (
	"fmt"
	"time"
)

func notaParaLetra(nota float64) string {
	switch nota {
	case 10:
		fallthrough
	case 9.5:
		return "A-"
	case 9:
		return "B+"
	case 8.5:
		return "B"
	case 8:
		return "B-"
	default:
		return "C"
	}
}

func switch2() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")

	}
}

func switch3(i interface{}) string {

	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float64"
	default:
		return "outro"
	}
}

func main() {

}
