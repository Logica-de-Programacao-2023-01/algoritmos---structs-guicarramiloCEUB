package main

import "fmt"

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func main() {
	a := Aluno{nome: "gui", idade: 20, notas: []float64{}}
	alt(&a)
}

func alt(a *Aluno) {
	for i := 1; i != 0; i++ {
		var op int
		var nota float64
		fmt.Println("Oq deseja fazer: ")
		fmt.Println("0) Sair")
		fmt.Println("1) Adicionar nota")
		fmt.Println("2) Remover nota")
		fmt.Println("3) Calcular media")
		fmt.Println("4) Imprimir aluno")
		fmt.Scanln(&op)
		if op == 0 {
			break
		} else if op == 1 {
			fmt.Println("insira nota: ")
			fmt.Scanln(&nota)
			(*a).notas = append((*a).notas, nota)
		} else if op == 2 {
			index := 0
			fmt.Println("insira nota: ")
			fmt.Scanln(&nota)
			for i := 0; i < len((*a).notas); i++ {
				if nota == (*a).notas[i] {
					index = i
				}
			}
			(*a).notas = append((*a).notas[:index], (*a).notas[index+1:]...)

		} else if op == 3 {
			soma := 0.0
			count := 0.0
			for i := 0; i < len((*a).notas); i++ {
				soma += (*a).notas[i]
				count += 1.0
			}
			fmt.Printf("media: %f \n", soma/count)
		} else if op == 4 {
			fmt.Println(*a)
		}
	}
}
