package main

import (
	"fmt"
	"math"
)

func main()  {
	var numero int
	var pares = []int{}

	fmt.Scan(&numero)

	if numero > 5 && numero < 2000 {
		for i := 1; i <= numero; i++ {
			if i % 2 == 0 {
				pares = append(pares, i)
			}
		}
	} else {
		fmt.Println("VALOR INVALIDO")
	}

	for _, valor := range pares {
		fmt.Printf("%v^2 = %.0f\n", valor, math.Pow(float64(valor), 2))
	}
}