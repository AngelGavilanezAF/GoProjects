package main

import (
	"fmt"
	"time"
)

func mensaje(string) {
	fmt.Println("Hola desde un go rutine")
}

func main() {
	go mensaje("Hola desde un go rutin")
	time.Sleep(10 * time.Second)
}
