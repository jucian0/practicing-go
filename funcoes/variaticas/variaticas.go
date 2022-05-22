package main

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	resultado := media(1.0, 2.0, 3.0, 4.0, 5.0)
	println(resultado)
}
