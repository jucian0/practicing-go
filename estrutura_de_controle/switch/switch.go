package main

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

	default:
		return "C"

	}

}

func main() {

}
