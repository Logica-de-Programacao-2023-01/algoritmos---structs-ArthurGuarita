package main

import "fmt"

type Filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []int
}

func main() {
	c := Filme{
		titulo:     "Harry Potter e a Pedra Filosofal",
		diretor:    "Chris Columbus",
		ano:        2001,
		avaliacoes: []int{8, 9, 10},
	}
	adicionar := func(c Filme) []int {
		var add int

		fmt.Print("Digite uma nova avaliação: ")
		fmt.Scan(&add)
		c.avaliacoes = append(c.avaliacoes, add)
		return c.avaliacoes
	}
	remover := func(c Filme) []int {
		var remov int
		fmt.Print("Digite a avaliação a ser removida: ")
		fmt.Scan(&remov)
		for i := 0; i <= len(c.avaliacoes); i++ {
			if c.avaliacoes[i] == remov {
				c.avaliacoes = append(c.avaliacoes[:i], c.avaliacoes[i+1:]...)
				break
			}
		}
		return c.avaliacoes
	}
	media := func(c Filme) float64 {
		var soma float64 = 0
		for _, n := range c.avaliacoes {
			soma += float64(n)
		}
		return soma / float64(len(c.avaliacoes))
	}
	imprimirFilme := func(c Filme) {
		fmt.Println("Titulo: ", c.titulo)
		fmt.Println("Diretor: ", c.diretor)
		fmt.Println("Ano: ", c.ano)
		fmt.Println("Avaliações: ", c.avaliacoes)
		fmt.Println("Média das avaliações: ", media(c))
	}
	adicionar(c)
	remover(c)
	imprimirFilme(c)
}
