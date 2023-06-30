package main

import (
	"fmt"
)

type Endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	idade    int
	endereco Endereco
}

func main() {
	p := Pessoa{nome: "Gui", idade: 20, endereco: Endereco{rua: "Asa Sul", estado: "DF", numero: 103, cidade: "Brasília"}}
	endereco(p)
}

func endereco(p Pessoa) {
	fmt.Println(p.endereco.estado)
	fmt.Println(p.endereco.cidade)
	fmt.Println(p.endereco.rua)
	fmt.Println(p.endereco.numero)
}
