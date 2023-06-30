package main

import "fmt"

type Viagem struct {
	origem  string
	destino string
	data    string
	preco   float64
}

func main() {
	s := []Viagem{{
		origem:  "bsb",
		destino: "rio",
		data:    "25/06",
		preco:   875.99,
	},
		{
			origem:  "slz",
			destino: "bsb",
			data:    "27/08",
			preco:   1004.99,
		},
	}
	caro(s)
}

func caro(s []Viagem) {
	var mais_barato Viagem
	for i := 0; i < len(s); i++ {
		if mais_barato.preco < s[i].preco {
			mais_barato = s[i]
		}
	}
	fmt.Println(mais_barato)
}
