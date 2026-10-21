package main

import (
	"fmt"
	"sync"
)

// 1. Definir una estructura Tarea con ID y Precio
type Tarea struct {
	ID     int
	Precio float64
}

func main() {
	// 2. Crear un canal de punteros a tareas con un búfer de 5
	tareas := make(chan *Tarea, 5)

	// Usaremos sync.WaitGroup para esperar a que el worker termine de procesar todas las tareas
	var wg sync.WaitGroup
	wg.Add(1)

	// 3. Iniciar la goroutine (el worker)
	go func() {
		defer wg.Done()

		// Leer continuamente del canal usando for range
		for t := range tareas {
			// 4. Modificar el precio apuntado (simulando un impuesto del 16%)
			t.Precio = t.Precio * 1.16
			fmt.Printf("[Worker] Tarea %d procesada. Nuevo precio: %.2f\n", t.ID, t.Precio)
		}
	}()

	// 5. En el main, crear instancias de Tarea
	t1 := Tarea{ID: 1, Precio: 100.0}
	t2 := Tarea{ID: 2, Precio: 200.0}

	fmt.Println("--- Valores Originales Antes del Worker ---")
	fmt.Printf("Tarea 1: ID=%d, Precio=%.2f\n", t1.ID, t1.Precio)
	fmt.Printf("Tarea 2: ID=%d, Precio=%.2f\n\n", t2.ID, t2.Precio)

	// Enviar las direcciones de las instancias al canal
	tareas <- &t1
	tareas <- &t2

	// Cerrar el canal para indicar al worker que ya no hay más tareas (permite que termine el for range)
	close(tareas)

	// Esperar a que el worker termine de procesar todas las tareas del canal
	wg.Wait()

	// Mostrar los objetos originales modificados desde el main
	fmt.Println("\n--- Resultados en el Main (Objetos Modificados) ---")
	fmt.Printf("Tarea 1: ID=%d, Precio=%.2f\n", t1.ID, t1.Precio)
	fmt.Printf("Tarea 2: ID=%d, Precio=%.2f\n", t2.ID, t2.Precio)
}
