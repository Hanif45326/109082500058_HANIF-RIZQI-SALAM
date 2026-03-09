package main

import "fmt"

func BiayaPos(beratGram int) {
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000
	biayaKg := beratKg * 10000

	var tambahanBiaya int
	if beratKg >= 10 {
		tambahanBiaya = 0
	} else if sisaGram > 500 {
		tambahanBiaya = sisaGram * 5
	} else {
		tambahanBiaya = sisaGram * 15
	}
	totalBiaya := biayaKg + tambahanBiaya

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Printf("Detail berat  : %d kg + %d  gram\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya  : Rp. %d  + Rp. %d\n", biayaKg, tambahanBiaya)
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}
func main() {
	var beratGram int
	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&beratGram)
	BiayaPos(beratGram)
}
