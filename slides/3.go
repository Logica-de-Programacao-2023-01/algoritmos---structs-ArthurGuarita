package main

import "fmt"

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func main() {
	c := Aluno{
		nome:  "Arthur",
		idade: 17,
		notas: []float64{7.6, 9.2, 6.2},
	}
	mediaNotas(c)
}
func mediaNotas(c Aluno) {
	fmt.Println("Aluno: ", c.nome)
	fmt.Println("Idade: ", c.idade)
	var soma float64 = 0
	for _, n := range c.notas {
		soma += n
	}
	fmt.Println("Media: ", soma/float64(len(c.notas)))
}
