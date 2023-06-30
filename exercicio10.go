package main

import "fmt"

type Filme struct {
	titulo  string
	diretor string
	ano     int
	notas   []int
	media   float64
}

func main() {
	f := Filme{
		titulo:  "aus",
		diretor: "gui",
		ano:     2023,
		notas:   []int{},
		media:   0,
	}
	filme(&f)
}

func filme(f *Filme) {
	for i := 1; i != 0; i++ {
		var op int
		var nota int
		fmt.Println("Oq deseja fazer: ")
		fmt.Println("0) Sair")
		fmt.Println("1) Adicionar nota")
		fmt.Println("2) Remover nota")
		fmt.Println("3) Calcular media")
		fmt.Println("4) Imprimir filme")
		fmt.Scanln(&op)
		if op == 0 {
			break
		} else if op == 1 {
			fmt.Println("insira nota: ")
			fmt.Scanln(&nota)
			(*f).notas = append((*f).notas, nota)
		} else if op == 2 {
			index := 0
			fmt.Println("insira nota: ")
			fmt.Scanln(&nota)
			for i := 0; i < len((*f).notas); i++ {
				if nota == (*f).notas[i] {
					index = i
				}
			}
			(*f).notas = append((*f).notas[:index], (*f).notas[index+1:]...)

		} else if op == 3 {
			soma := 0
			count := 0.0
			for i := 0; i < len((*f).notas); i++ {
				soma += (*f).notas[i]
				count += 1.0
			}
			(*f).media = float64(soma) / count
		} else if op == 4 {
			fmt.Println(*f)
		}
	}
}
