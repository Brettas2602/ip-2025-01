package main

import "fmt"

func main()  {
	var destino, tipo int
	var valor float64

	fmt.Println("De acordo com a tabela:")
	fmt.Println("===============================================================")
	fmt.Println("1 - Região Norte")
	fmt.Println("2 - Região Nordeste")
	fmt.Println("3 - Região Centro-Oeste")
	fmt.Println("4 - Região Sul")
	fmt.Println("===============================================================")
	fmt.Println("Digite o número referente a seu destino")
	fmt.Scan(&destino)

	fmt.Println("Informe o tipo da sua viagem")
	fmt.Println("===============================================================")
	fmt.Println("1 - Só ida")
	fmt.Println("2 - Ida e volta")
	fmt.Println("===============================================================")
	fmt.Scan(&tipo)


	if (destino < 0 || destino > 4) || (tipo < 0 || tipo > 2) {
		fmt.Println("Valores inválidos!")
	}else{
		switch destino {
		case 1:
			switch tipo {
			case 1:
				valor = 500
	
			case 2: 
				valor = 900
			}
	
		case 2: 
			switch tipo {
			case 1:
				valor = 350
	
			case 2: 
				valor = 650
			}
	
		case 3: 
			switch tipo {
			case 1:
				valor = 350
	
			case 2:
				valor = 600
			}
	
		case 4: 
			switch tipo {
			case 1:
				valor = 300
			
			case 2: 
				valor = 550
			}
		}
	
		fmt.Printf("O valor da passagem é igual a = %.2f\n", valor)
	}
}