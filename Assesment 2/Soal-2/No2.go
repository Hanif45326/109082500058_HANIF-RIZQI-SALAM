package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(tab arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if tab[i].NIM == nim {
			return tab[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(tab arrayMahasiswa, n int, nim string) int {
	max := -1
	found := false
	for i := 0; i < n; i++ {
		if tab[i].NIM == nim {
			if !found || tab[i].nilai > max {
				max = tab[i].nilai
			}
			found = true

		}
	}
	return max
}

func main() {
	var tab arrayMahasiswa
	var n int
	var nimCari string

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	if n > nMax {
		n = nMax
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&tab[i].NIM, &tab[i].nama, &tab[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	pertama := cariNilaiPertama(tab, n, nimCari)
	terbesar := cariNilaiTerbesar(tab, n, nimCari)

	if pertama != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, pertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, terbesar)
	} else {
		fmt.Println("NIM tidak ditemukan.")
	}
}
