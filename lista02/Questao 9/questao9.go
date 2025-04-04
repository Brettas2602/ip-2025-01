package main

import "fmt"

func main()  {
	var compra, venda float64

	fmt.Println("Digite o valor da compra")
	fmt.Scan(&compra)

	if compra > 0 {
		if compra < 10 {
			venda = compra * 1.7
		}else if compra >= 10 && compra < 30 {
			venda = compra * 1.5
		}else if compra >= 30 && compra < 50 {
			venda = compra * 1.4
		}else if compra >= 50 {
			venda = compra * 1.3
		}

		fmt.Printf("O valor da venda é igual a = %.2f\n", venda)
	}else {
		fmt.Println("Valor de compra inválido!")
	}
}