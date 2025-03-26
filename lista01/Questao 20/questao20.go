package main

import ("fmt")

func main()  {
	var horas, minutos, segundos, tempoTotal int

	fmt.Scan(&horas)
	fmt.Scan(&minutos)
	fmt.Scan(&segundos)

	tempoTotal = horas * 3600 + minutos * 60 + segundos

	fmt.Printf("O TEMPO EM SEGUNDOS E = %v\n", tempoTotal)
}