package main

import "fmt"

type suhu float64

func CelciusToReamur(c suhu) suhu {
	return c * 0.8
}
func CelciusToFahrenheit(c suhu) suhu {
	return (c * 1.8) + 32
}
func CelciusToKelvin(c suhu) suhu {
	return c + 273.15
}

func main() {
	var inputCelcius suhu

	fmt.Println("=== KONVERT CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius): ")
	fmt.Scan(&inputCelcius)

	reamur := CelciusToReamur(inputCelcius)
	fahrenheit := CelciusToFahrenheit(inputCelcius)
	kelvin := CelciusToKelvin(inputCelcius)

	fmt.Println()
	fmt.Printf("%.0f celcius = %.1f reamur\n", inputCelcius, reamur)
	fmt.Printf("%.0f celcius = %.1f fahrenheit\n", inputCelcius, fahrenheit)
	fmt.Printf("%.0f celcius = %.2f kelvin\n", inputCelcius, kelvin)

}
