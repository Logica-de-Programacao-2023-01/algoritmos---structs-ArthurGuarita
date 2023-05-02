package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func main() {
	c := Circulo{
		raio: 4,
	}
	fmt.Print(areaPi(c))
}
func areaPi(c Circulo) float64 {
	area := math.Pi * (c.raio * c.raio)
	return area
}
