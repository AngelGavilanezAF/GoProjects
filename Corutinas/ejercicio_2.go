package main

import (
	"fmt"
	"time"
)

func tareas(id int, txt string) {
	for i := 0; i < 3; i++ {
		fmt.Println(txt)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go tareas(1, "Proceso A")
	go tareas(2, "Proceso B")
	go tareas(3, "Proceso C")

	time.Sleep(2 * time.Second)
}
