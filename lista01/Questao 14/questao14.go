package main

import (
	"fmt"
	"math"
)

func main()  {
	var altura, aresta, areaBase, volume float64

	fmt.Scan(&altura)
	fmt.Scan(&aresta)

	areaBase = (3 * math.Pow(aresta, 2) * math.Sqrt(3)) / 2
	volume = (areaBase * altura) / 3

	fmt.Printf("O VOLUME DA PIRAMIDE E %.2f METROS CUBICOS", volume)
}