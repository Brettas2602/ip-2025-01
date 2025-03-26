package main

import ("fmt")

func main()  {
	var valorInicial, razao, numeroElementos, soma int

	fmt.Scan(&valorInicial, &razao, &numeroElementos)

	for i := 0; i < numeroElementos; i++ {
		soma += valorInicial + razao * i
	}

	fmt.Println(soma)
}