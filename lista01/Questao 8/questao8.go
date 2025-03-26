package main

import (
	"fmt"
	"math"
)

func main()  {
	const PI float64 = 3.14159
	var raio, altura, areaCilindro, custoLata float64

	fmt.Scan(&raio)
	fmt.Scan(&altura)

	areaCilindro = (2 * PI * math.Pow(raio, 2)) + (2 * PI * raio * altura)
	custoLata = areaCilindro * 100

	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custoLata)
}