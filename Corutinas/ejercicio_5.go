package main

import (
	"fmt"
	"sync"
)

func tarea(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("La tarea", id, "ha terminado")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go tarea(i, &wg)
	}

	wg.Wait()

	fmt.Println("Todas las tareas terminaron")
}
