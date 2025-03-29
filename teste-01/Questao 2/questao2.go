package main

import ("fmt")

func main()  {
	var numero float64

	fmt.Println("Digite um numero")
	fmt.Scan(&numero)

	if (numero > 20 && numero < 90) {
		fmt.Println("O numero esta entre 20 e 90")
	} else {
		fmt.Println("O numero nao esta entre 20 e 90")
	}
}