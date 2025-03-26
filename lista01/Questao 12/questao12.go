package main

import ("fmt")

func main()  {
	var horas int
	var valorConta float64

	fmt.Scan(&horas)

	valorConta = float64((horas / 3 * 10) + (horas % 3 * 5))

	fmt.Printf("O VALOR A PAGAR E %.2f\n", valorConta)
}