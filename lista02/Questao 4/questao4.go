package main

import (
	"fmt"
	"math"
)

func main()  {
	var num float64

	fmt.Println("Digite um número")
	fmt.Scan(&num)

	if num >= 0 {
		fmt.Printf("A raiz quadrada de %.2f é %.2f\n", num, math.Sqrt(num))
	} else {
		fmt.Printf("O quadrado de %.2f é %.2f\n", num, math.Pow(num, 2))
	}
}