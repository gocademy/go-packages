package interest

import "fmt"

func init() {
	fmt.Println("interest Paket initialisiert.")
}

// calculate berechnet und gibt die einfachen Zinsen für einen
// Kapitalbetrag p, den Zinssatz r für die Zeitdauer t Jahre zurück
func calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}

func ExportedCalulate(p float64, r float64, t float64) float64 {
	return calculate(p, r, t)
}
