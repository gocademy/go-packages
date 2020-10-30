package main

import (
	_ "finance/env"
	"finance/interest"
	"fmt"
	"log"
	"os"
)

var p, r, t = 5000.0, 10.0, 1.0

var testVariable = "Hello from testVariable before Init()"

/*
* init function to check if p, r and t are greater than zero
 */
func init() {
	println("Main Paket initialisiert.")
	if p < 0 {
		log.Fatal("Principal weniger als 0")
	}
	if r < 0 {
		log.Fatal("Zinssatz is weniger als 0")
	}
	if t < 0 {
		log.Fatal("Dauer ist weniger als 0")
	}

	fmt.Println("Inhalt der Test Variable:")
	fmt.Println(testVariable)

	testVariable = os.Getenv("TEST_EXPORT")
	fmt.Println(testVariable)
}

func main() {
	fmt.Println("Starte Berechnung:")
	si := interest.ExportedCalulate(p, r, t)

	fmt.Println("Interest:", si)
}
