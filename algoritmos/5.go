package main

import "fmt"

type musica struct {
	titulo  string
	artista string
	duracao float64
}
type Playlists struct {
	nome     string
	musicass []musica
}

func main() {
	c := []Playlists{
		{nome: "Random Musics",
			musicass: []musica{
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
			}},
		{nome: "Another random musics",
			musicass: []musica{
				{
					titulo:  "Photography",
					artista: "Ed sherran",
					duracao: 4.10,
				},
				{
					titulo:  "Without your love",
					artista: "Conro",
					duracao: 3.45,
				},
			},
		},
	}
	fmt.Print(encontrarPlaylist("Without your love", c))

}
func encontrarPlaylist(nomeMusica string, c []Playlists) []string {
	var play []string
	for _, p := range c {
		for _, m := range p.musicass {
			if m.titulo == nomeMusica {
				play = append(play, p.nome)
			}
		}
	}
	return play
}
