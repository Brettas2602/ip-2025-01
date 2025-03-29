package main

import ("fmt")

func main()  {
	var numero float64

	fmt.Println("Digite um numero")
	fmt.Scan(&numero)

	if (numero > 0) {
		fmt.Println("O numero e positivo")
	} else if (numero < 0) {
		fmt.Println("O numero e negativo")
	} else {
		fmt.Println("O numero e nulo")
	}
}