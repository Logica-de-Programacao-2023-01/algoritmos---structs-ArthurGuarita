package main

import (
	"fmt"
)

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func main() {
	c := Aluno{
		nome:  "Arthur",
		idade: 17,
		notas: []float64{7.8, 6.9, 8.9},
	}
	adicionar := func(c Aluno) []float64 {
		var n int
		fmt.Print("Digite uma nova nota: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Print("erro")
			return nil
		}
		c.notas = append(c.notas, float64(n))
		return c.notas
	}
	remover := func(c Aluno) []float64 {
		var m int
		fmt.Print("Digite a nota que deseja remover: ")
		_, err := fmt.Scan(&m)
		if err != nil {
			fmt.Print("erro")
			return nil
		}
		for i := 0; i < len(c.notas); i++ {
			if c.notas[i] == float64(m) {
				c.notas = append(c.notas[:i], c.notas[i+1:]...)
				break
			}
		}
		return c.notas
	}
	media := func(c Aluno) float64 {
		var soma float64 = 0
		for _, l := range c.notas {
			soma += l
		}
		return soma / float64(len(c.notas))
	}
	imprimir := func(c Aluno) {
		fmt.Println("Nome: ", c.nome)
		fmt.Println("Idade: ", c.idade)
		fmt.Println("Média: ", media(c))
	}
	adicionar(c)
	remover(c)
	imprimir(c)
}
