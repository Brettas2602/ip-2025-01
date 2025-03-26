package main

import ("fmt")

func main()  {
	var numeroCasosDeTeste, numeroDePessoas int
	var porcenPopular, porcenGeral, porcenArquibancada, porcenCadeiras float64
	var rendaTotal, rendaPopular, rendaGeral, rendaArquibancada, rendaCadeiras float64
	rendaDosJogos := []float64{}

	fmt.Scan(&numeroCasosDeTeste)

	for i := 0; i < numeroCasosDeTeste; i++ {
		fmt.Scan(&numeroDePessoas, &porcenPopular, &porcenGeral, &porcenArquibancada, &porcenCadeiras)

		rendaPopular = float64(numeroDePessoas) * (porcenPopular/100)
		rendaGeral = (float64(numeroDePessoas) * (porcenGeral/100)) * 5
		rendaArquibancada = (float64(numeroDePessoas) * (porcenArquibancada/100)) * 10
		rendaCadeiras = (float64(numeroDePessoas) * (porcenCadeiras/100)) * 20
		rendaTotal = rendaPopular +  rendaGeral + rendaArquibancada + rendaCadeiras

		rendaDosJogos = append(rendaDosJogos, rendaTotal)
	}

	for index, value := range rendaDosJogos {
		fmt.Printf("A RENDA DO JOGO N. %v E = %.2f\n", index+1, value)
	}

}