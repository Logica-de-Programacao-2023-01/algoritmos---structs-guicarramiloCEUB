package main

import "fmt"

type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

func main() {
	f := Funcionario{nome: "luiz", salario: 1000.0, idade: 55}
	sal(f)
	tempo(f)
}

func sal(f Funcionario) {
	var p float64
	fmt.Println("porcentagem desejada: ")
	fmt.Scanln(&p)
	fmt.Println(f.salario * (1.0 + (p / 100)))
}

func tempo(f Funcionario) {
	fmt.Println(f.idade - 18)
}
