package main

import "fmt"

func main() {

	warna := [4]string{"merah", "kuning", "hijau", "ungu"}
	berhasil := true

	for percobaan := 1; percobaan <= 5; percobaan++ {
		fmt.Printf("Percobaan %d : ", percobaan)

		var warna1, warna2, warna3, warna4 string
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		if warna1 != warna[0] || warna2 != warna[1] ||
			warna3 != warna[2] || warna4 != warna[3] {
			berhasil = false
		}
	}
	fmt.Printf("BERHASIL : %t\n", berhasil)
}
