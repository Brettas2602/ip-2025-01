package main

import ("fmt")

func main()  {
	var matriz = [2][2]float64{}
	var determinante float64

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	determinante = matriz[0][0] * matriz[1][1] - matriz[0][1] * matriz[1][0]

	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f", determinante)
}