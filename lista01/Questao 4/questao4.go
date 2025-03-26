package main

import ("fmt")

func main()  {
	var salarioMinimo, consumoEmKw, custoPorKw, custoDoConsumo, custoComDesconto float64

	fmt.Scan(&salarioMinimo)
	fmt.Scan(&consumoEmKw)

	custoPorKw = salarioMinimo * 0.007
	custoDoConsumo = custoPorKw * consumoEmKw
	custoComDesconto = custoDoConsumo * 0.9

	fmt.Printf("Custo por kw: R$ %.2f\n", custoPorKw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoDoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoComDesconto)
}