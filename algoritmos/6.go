package main

import "fmt"

type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

func main() {
	c := Funcionario{
		nome:    "Mike Ross",
		salario: 25000,
		idade:   25,
	}
	fmt.Println("Salário aumentado em 10%: ", aumentarSalario(c))
	fmt.Println("Salário diminuído em 10%: ", diminuirSalario(c))
	fmt.Println("Tempo de serviço: ", tempoServico(c))
}

func aumentarSalario(c Funcionario) float64 {
	return c.salario * 1.1
}
func diminuirSalario(c Funcionario) float64 {
	return c.salario * 0.9
}
func tempoServico(c Funcionario) int {
	return c.idade - 18
}
