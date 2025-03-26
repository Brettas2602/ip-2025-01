package main

import ("fmt")

func main()  {
	var nota1, nota2, nota3, media float64

	fmt.Println("Escreva os valores das 3 notas:")
	fmt.Scan(&nota1, &nota2, &nota3)
	media = (nota1 + nota2 + nota3) / 3

	if media >= 6 {
		fmt.Printf("MEDIA = %.2f\n", media)
		fmt.Print("APROVADO")
	} else {
		fmt.Printf("MEDIA = %.2f\n", media)
		fmt.Print("REPROVADO")
	}
}