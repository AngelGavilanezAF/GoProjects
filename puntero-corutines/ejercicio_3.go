package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. Declarar un slice de enteros de tamaño 5
	datos := make([]int, 5)

	var wg sync.WaitGroup

	// 2. Lanzar 5 goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(p *int, idx int) {
			defer wg.Done()

			// 3. Multiplicar el índice por 10
			*p = idx * 10
		}(&datos[i], i)
	}

	// 4. Esperar a que terminen todas las goroutines
	wg.Wait()

	// 5. Imprimir el slice
	fmt.Println("--- Slice resultante ---")
	fmt.Println(datos)
}
