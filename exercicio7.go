package main

import "fmt"

type Animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

func main() {
	a := Animal{
		nome:    "gui",
		especie: "humano",
		idade:   20,
		som:     "fala",
	}
	som(&a)
	infos(a)
}
func som(ptr *Animal) {
	var n_som string
	fmt.Println("novo som: ")
	fmt.Scanln(&n_som)
	(*ptr).som = n_som
}
func infos(a Animal) {
	fmt.Println(a)
}
