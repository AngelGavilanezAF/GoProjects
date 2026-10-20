package main

import "fmt"

// Definición de la estructura CuentaBancaria con el campo saldo
type CuentaBancaria struct {
	saldo float64
}

// Método Depositar utilizando un receptor de puntero (*CuentaBancaria)
func (c *CuentaBancaria) Depositar(monto float64) {
	c.saldo += monto
	fmt.Printf("Depósito exitoso. Saldo actual: %.2f\n", c.saldo)
}

// Método Retirar utilizando un receptor de puntero con validación de fondos
func (c *CuentaBancaria) Retirar(monto float64) {
	if monto > c.saldo {
		fmt.Println("Error: Fondos insuficientes para realizar el retiro.")
		return
	}
	c.saldo -= monto
	fmt.Printf("Retiro exitoso. Saldo actual: %.2f\n", c.saldo)
}

// Función principal para probar la cuenta, depósitos y retiros
func main() {
	// Crear una instancia de CuentaBancaria
	miCuenta := CuentaBancaria{saldo: 500.0}

	fmt.Printf("Saldo inicial en la cuenta: %.2f\n", miCuenta.saldo)

	// Realizar un depósito
	miCuenta.Depositar(200.0)

	// Realizar un retiro válido
	miCuenta.Retirar(150.0)

	// Intentar un retiro mayor al saldo actual
	miCuenta.Retirar(700.0)

	// Mostrar el saldo final
	fmt.Printf("Saldo final de la cuenta: %.2f\n", miCuenta.saldo)
}