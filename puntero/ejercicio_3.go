package main

import "fmt"

type Libro struct {
	titulo string
	autor  string
	precio float64
}


func aplicarDescuento(l *Libro, porcentaje float64) {
	
	l.precio = l.precio - (l.precio * porcentaje)
}

func main() {
	
	miLibro := Libro{
		titulo: "El mata burras",
		autor:  "Jose Jose",
		precio: 30000.0,
	}
	
	fmt.Println("Precio base: ", miLibro.precio)
	
	aplicarDescuento(&miLibro, 0.2)
	
	fmt.Println("Precio final con descuento: ", miLibro.precio)
}