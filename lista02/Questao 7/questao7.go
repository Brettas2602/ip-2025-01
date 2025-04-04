package main

import "fmt"

func main()  {
	var numeros = [3]int{}
	var MAIOR, MENOR, INTER, soma int

	fmt.Println("Digite 3 números distintos")
	for i := 0; i < len(numeros); i++ {
		fmt.Scan(&numeros[i])
	}

	for i := 0; i < len(numeros); i++ {
		if i == 0 {
			MAIOR = numeros[i]
			MENOR = numeros[i]
		}

		if numeros[i] > MAIOR {
			MAIOR = numeros[i]
		}

		if numeros[i] < MENOR {
			MENOR = numeros[i]
		}

		soma += numeros[i]
	}

	INTER = soma - MAIOR - MENOR

	fmt.Printf("Os números ordenados em ordem crescente são: %v, %v, %v\n", MENOR, INTER, MAIOR)
}