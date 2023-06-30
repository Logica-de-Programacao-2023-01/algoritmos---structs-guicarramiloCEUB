package main

import "fmt"

type Musica struct {
	titulo  string
	artista string
	duracao int
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func main() {
	p := Playlist{nome: "musicas show", musicas: []Musica{{titulo: "b d", artista: "gui", duracao: 3}, {titulo: "as", artista: "ju", duracao: 4}}}
	play(p)
}

func play(p Playlist) {
	total := 0
	for i := 0; i < len(p.musicas); i++ {
		total += p.musicas[i].duracao
		fmt.Println(p.musicas[i].titulo)
		fmt.Println(p.musicas[i].artista)
		fmt.Println(p.musicas[i].duracao)
	}
	fmt.Println("duração total: ", total)
}
