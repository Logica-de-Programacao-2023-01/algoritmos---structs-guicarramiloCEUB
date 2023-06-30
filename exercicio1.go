package main

import "fmt"

type circulo struct {
	raio float64
}

func main() {
	c := circulo{raio: 4}
	area(c)
}
func area(c circulo) {
	fmt.Println(3.14 * c.raio * c.raio)
}
