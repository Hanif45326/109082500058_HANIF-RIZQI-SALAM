package main

import "fmt"

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIdx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			hasil := (tumbuh[i] + 1) * float64(pop[i])
			fmt.Printf("%s %.0f\n", prov[i], hasil)
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var provinsi NamaProv
	var populasi PopProv
	var pertumbuhan TumbuhProv
	var cariNama string

	InputData(&provinsi, &populasi, &pertumbuhan)
	fmt.Scan(&cariNama)

	idxTercepat := ProvinsiTercepat(pertumbuhan)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", provinsi[idxTercepat])

	idxCari := IndeksProvinsi(provinsi, cariNama)
	if idxCari != -1 {
		fmt.Printf("\nData provinsi yang dicari : %s\n", provinsi[idxCari])
	} else {
		fmt.Println("\nProvinsi tidak ditemukan.")
	}

	Prediksi(provinsi, populasi, pertumbuhan)
}
