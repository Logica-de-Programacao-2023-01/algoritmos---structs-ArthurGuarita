package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
	ano    string
}

func main() {
	c := Livro{
		titulo: "'Bob esponja: uma aventura fora do mar'",
		autor:  "Stephen Hillenburg",
		ano:    "2054",
	}
	imprimir(c)
}
func imprimir(c Livro) {
	fmt.Println("Título: ", c.titulo)
	fmt.Println("Autor: ", c.autor)
	fmt.Println("Ano: ", c.ano)
}
