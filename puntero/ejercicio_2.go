package main
import "fmt"

func intercambiar(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 5
	b := 10
	
	fmt.Println("a antes: ", a)
	fmt.Println("b antes: ", b)
	
	intercambiar(&a, &b)
	
	fmt.Println("a despues: ", a)
	fmt.Println("b despues: ", b)
}