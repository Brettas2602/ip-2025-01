package main

import (
	"fmt"
	"math"
)

func main()  {
	var conta int
	var consumoAgua, valorDaConta float64
	var tipoConsumidor string

	fmt.Scan(&conta, &consumoAgua, &tipoConsumidor)

	switch tipoConsumidor {
	case "R", "r":
		valorDaConta = 5 + (0.05 * consumoAgua)

	case "C", "c":
		valorDaConta = 500 + (math.Max(0, consumoAgua - 80) * 0.25)

	case "I", "i":
		valorDaConta = 800 + (math.Max(0, consumoAgua - 100) * 0.04)
	}

	fmt.Printf("CONTA = %v\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f", valorDaConta)
}