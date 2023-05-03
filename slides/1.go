package main

import "fmt"

type Retangulo struct {
	largura float64
	altura  float64
}

func main() {
	c := Retangulo{
		largura: 20,
		altura:  10,
	}
	fmt.Print(area(c))
}

func area(c Retangulo) float64 {
	areaR := c.largura * c.altura
	return areaR
}
