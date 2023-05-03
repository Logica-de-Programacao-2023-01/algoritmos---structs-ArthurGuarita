package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func main() {
	c := Triangulo{
		base:   4,
		altura: 3,
	}
	fmt.Print(areaTriang(c))
}
func areaTriang(c Triangulo) float64 {
	return (c.base * c.altura) / 2
}
