package main

import (
	"fmt"
	"sync"
)

// 1. Definir la estructura Servidor
type Servidor struct {
	Nombre string
	IP     string
	Activo bool
}

func main() {
	// Lista con los nombres e IP de los servidores
	nombres := []string{
		"Servidor Web",
		"Servidor Base de Datos",
		"Servidor API",
	}

	ips := []string{
		"192.168.1.10",
		"192.168.1.20",
		"192.168.1.30",
	}

	// Crear un canal para enviar punteros a Servidor
	canal := make(chan *Servidor)

	var wg sync.WaitGroup

	// 2. Lanzar una goroutine por cada servidor
	for i := 0; i < len(nombres); i++ {
		wg.Add(1)

		go func(nombre, ip string) {
			defer wg.Done()

			// 3. Configurar una instancia local
			s := &Servidor{
				Nombre: nombre,
				IP:     ip,
				Activo: true,
			}

			// Enviar el puntero al canal
			canal <- s
		}(nombres[i], ips[i])
	}

	// Cerrar el canal cuando todas las goroutines terminen
	go func() {
		wg.Wait()
		close(canal)
	}()

	// 4. Recibir los punteros e imprimir los servidores
	fmt.Println("--- Servidores configurados ---")

	for servidor := range canal {
		fmt.Printf("Nombre: %s\n", servidor.Nombre)
		fmt.Printf("IP: %s\n", servidor.IP)
		fmt.Printf("Activo: %v\n", servidor.Activo)
		fmt.Println("------------------------------")
	}
}
