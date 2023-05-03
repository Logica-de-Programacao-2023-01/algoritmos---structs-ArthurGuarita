package main

import "fmt"

type Animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

func main() {
	c := Animal{
		nome:    "Morcego",
		especie: "Desmodus rotundus",
		idade:   19,
		som:     "ti ti ti",
	}
	imprimirInfo(c)
}
func modificarSom(c Animal) string {
	var sound string
	fmt.Print("Digite o novo som: ")
	fmt.Scan(&sound)
	c.som = sound
	return sound
}
func imprimirInfo(c Animal) {
	fmt.Println("Nome: ", c.nome)
	fmt.Println("Espécie: ", c.especie)
	fmt.Println("Idade: ", c.idade)
	fmt.Println("Som: ", modificarSom(c))
}
