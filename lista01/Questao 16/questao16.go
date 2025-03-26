package main

import ("fmt")

func main()  {
	var salario, salarioReajustado float64

	fmt.Scan(&salario)

	if salario <= 300 {
		salarioReajustado = salario * 1.5
	} else {
		salarioReajustado = salario * 1.3
	}

	fmt.Printf("SALARIO COM REAJUSTE %.2f\n", salarioReajustado)
}