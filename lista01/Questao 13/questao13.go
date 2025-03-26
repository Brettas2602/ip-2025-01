package main

import("fmt")

func main()  {
	var nota float64
	var conceito string

	fmt.Scan(&nota)

	if nota < 6 {
		conceito = "D"
	} else if nota >= 6 && nota < 7.5 {
		conceito = "C"
	} else if nota >= 7.5 && nota < 9 {
		conceito = "B"
	} else if nota >= 9 && nota <= 10 {
		conceito = "A"
	} else {
		fmt.Println("NOTA INVALIDA")
	}

	fmt.Printf("NOTA = %.1f CONCEITO = %v\n", nota, conceito)
}