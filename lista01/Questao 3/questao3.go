package main

import (
	"fmt"
	"math"
)

func main()  {
	var numero, resultado, quadrado int
	var digitoInvalido = false

	for i := 2; i >= 0; i-- {
		fmt.Scan(&numero)
		if numero > 9 {
			fmt.Print("DIGITO INVALIDO")
			digitoInvalido = true
			break
		}
		resultado += numero * int(math.Pow10(i)) 
	}

	if !digitoInvalido {
		quadrado = resultado * resultado
		fmt.Print(resultado, " , ", quadrado)
	}
}