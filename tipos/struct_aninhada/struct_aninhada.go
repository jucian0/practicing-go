package main

import "fmt"

type item struct {
	productID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{1, 2, 12.10},
			{2, 1, 23.50},
			{11, 100, 3.50},
		},
	}

	fmt.Printf("Valor total, é de R$ %.2f", pedido.valorTotal())
}
