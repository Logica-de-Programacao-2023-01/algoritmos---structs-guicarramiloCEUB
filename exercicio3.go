package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func main() {
	t := Triangulo{base: 5.4, altura: 4.4}
	area_t(t)

}
func area_t(t Triangulo) {
	fmt.Println((t.altura * t.base) / 2)
}
