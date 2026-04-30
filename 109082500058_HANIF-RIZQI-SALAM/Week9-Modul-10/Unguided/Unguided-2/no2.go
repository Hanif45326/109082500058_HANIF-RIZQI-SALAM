package main

import (
	"fmt"
)

func main() {
	var beratIkan [1000]float64
	var x, y int

	fmt.Scan(&x, &y)

	if x > 1000 {
		x = 1000
	}

	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalBeratWadah [1000]float64
	var jumlahWadah int = 0

	for i := 0; i < x; i++ {
		indeksWadah := i / y
		totalBeratWadah[indeksWadah] += beratIkan[i]

		if indeksWadah >= jumlahWadah {
			jumlahWadah = indeksWadah + 1
		}
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f", totalBeratWadah[i])
		if i < jumlahWadah-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	for i := 0; i < jumlahWadah; i++ {
		var ikanDiWadahIni int

		if i == jumlahWadah-1 && x%y != 0 {
			ikanDiWadahIni = x % y
		} else {
			ikanDiWadahIni = y
		}

		rataRata := totalBeratWadah[i] / float64(ikanDiWadahIni)
		fmt.Printf("%.2f", rataRata)

		if i < jumlahWadah-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
