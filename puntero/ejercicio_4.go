package main

import "fmt"

func duplicar(srpuntero *int) {
    if srpuntero == nil {
        fmt.Println("Error: puntero nulo")
        return
    }

    *srpuntero *= 2
}

func main() {
    numero := 10

    fmt.Println("Valor inicial:", numero)

    duplicar(&numero)
    fmt.Println("Después de duplicar:", numero)

    duplicar(nil)
}