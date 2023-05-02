package main

import "fmt"

type endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}
type Pessoa struct {
	nome     string
	idade    int
	endereco endereco
}

func main() {
	c := Pessoa{
		nome:  "Alan",
		idade: 17,
		endereco: endereco{
			rua:    "Rua A",
			numero: 2,
			cidade: "Planaltina",
			estado: "DF",
		},
	}
	imprimEndereco(c)
}
func imprimEndereco(c Pessoa) {
	fmt.Println(c.endereco.rua)
	fmt.Println(c.endereco.numero)
	fmt.Println(c.endereco.cidade)
	fmt.Println(c.endereco.estado)
}
