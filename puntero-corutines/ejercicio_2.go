package main

import (
	"fmt"
	"sync"
)

type Resultado struct {
	Suma       int
	Completado bool
}

func calcularSuma(res *Resultado, numeros []int, wg *sync.WaitGroup) {
	defer wg.Done()

	sumaTemporal := 0

	for _, n := range numeros {
		sumaTemporal += n
	}

	res.Suma = sumaTemporal
	res.Completado = true
}

func main() {
	var res Resultado

	numeros := []int{10, 20, 30, 40, 50}

	var wg sync.WaitGroup
	wg.Add(1)

	go calcularSuma(&res, numeros, &wg)

	wg.Wait()

	fmt.Println("--- Resultados desde el main ---")
	fmt.Printf("Suma total: %d\n", res.Suma)
	fmt.Printf("¿Completado?: %v\n", res.Completado)
}
