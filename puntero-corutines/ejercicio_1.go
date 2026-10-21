package main

import (
	"fmt"
	"sync"
)

func main() {
	var contador int = 0

	var wg sync.WaitGroup

	const numGoroutines = 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		go func(ptr *int) {
			defer wg.Done()

			*ptr++
		}(&contador)
	}

	wg.Wait()

	fmt.Printf("Valor final del contador: %d\n", contador)

}
