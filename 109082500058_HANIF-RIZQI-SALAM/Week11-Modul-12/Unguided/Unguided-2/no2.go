package main

import (
	"fmt"
	"sort"
)

func main() {
	var input int
	suaraMasuk := 0
	suaraSah := 0

	rekapSuara := make(map[int]int)

	for {
		_, err := fmt.Scan(&input)
		if err != nil || input == 0 {
			break
		}

		suaraMasuk++

		if input >= 1 && input <= 20 {
			suaraSah++
			rekapSuara[input]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	type Calon struct {
		Nomor int
		Suara int
	}

	var daftarPemenang []Calon
	for nomor, jumlah := range rekapSuara {
		daftarPemenang = append(daftarPemenang, Calon{nomor, jumlah})
	}

	sort.Slice(daftarPemenang, func(i, j int) bool {
		if daftarPemenang[i].Suara != daftarPemenang[j].Suara {
			return daftarPemenang[i].Suara > daftarPemenang[j].Suara
		}
		return daftarPemenang[i].Nomor < daftarPemenang[j].Nomor
	})

	if len(daftarPemenang) >= 1 {
		fmt.Printf("Ketua RT: %d\n", daftarPemenang[0].Nomor)
	}
	if len(daftarPemenang) >= 2 {
		fmt.Printf("Wakil ketua: %d\n", daftarPemenang[1].Nomor)
	}
}
