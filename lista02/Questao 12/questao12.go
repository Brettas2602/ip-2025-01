package main

import "fmt"

func main()  {
	var idade int
	var classificacao string

	fmt.Println("Informe sua idade")
	fmt.Scan(&idade)

	if idade <= 2 {
		classificacao = "Recém nascido"
	}else if idade >= 3 && idade <= 11 {
		classificacao = "Criança"
	}else if idade >= 12 && idade <= 19 {
		classificacao = "Adolescente"
	}else if idade >= 20 && idade <= 55 {
		classificacao = "Adulto"
	}else if idade > 55 {
		classificacao = "Idoso"
	}

	fmt.Printf("Sua classificação é: %s", classificacao)
}