package main

import "fmt"

type Viagem struct {
	origem  string
	destino string
	data    string
	preco   float64
}

func main() {
	c := []Viagem{
		{
			origem:  "Brasilia",
			destino: "São Paulo",
			data:    "18/02/2006",
			preco:   500.75,
		},
		{
			origem:  "São Paulo",
			destino: "Flórida, EUA",
			data:    "17/07/2024",
			preco:   1500.00,
		},
		{
			origem:  "Brasilia",
			destino: "Acre",
			data:    "15/03/2023",
			preco:   450.99,
		},
	}
	fmt.Print(precoCaro(c))
}
func precoCaro(c []Viagem) []Viagem {
	max := c[0].preco
	var r []Viagem
	for i := 1; i < len(c); i++ {
		if c[i].preco > max {
			max = c[i].preco
		}
	}
	for _, v := range c {
		if v.preco == max {
			r = append(r, v)
		}
	}
	return r
}
