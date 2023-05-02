package main

import "fmt"

type musica struct {
	titulo  string
	artista string
	duracao float64
}
type Playlist struct {
	nome    string
	musicas []musica
}

func main() {
	c := Playlist{
		nome: "Random Musics",
		musicas: []musica{
			{
				titulo:  "Without your love",
				artista: "Conro",
				duracao: 3.45,
			},
			{
				titulo:  "Eyes closed",
				artista: "Ed Sherran",
				duracao: 2.67,
			},
		},
	}
	imprimirPlaylist(c)
}
func imprimirPlaylist(c Playlist) {
	fmt.Println("Nome da playlist: ", c.nome)
	var soma float64 = 0
	for _, n := range c.musicas {
		soma += n.duracao
	}
	fmt.Println("Tempo total de duração: ", soma)
	for _, m := range c.musicas {
		fmt.Println("\nMúsica: ", m.titulo)
		fmt.Println("Artista: ", m.artista)
		fmt.Println("Duração: ", m.duracao)
		fmt.Println("")
	}
}
