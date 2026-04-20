package main

import "fmt"

type angka int
type kata string
type buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunterbit   angka
	jumlahHalaman angka
}

func main() {
	var b buku

	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan judul buku: ")
	fmt.Scan(&b.judul)
	fmt.Print("Masukkan nama penulis: ")
	fmt.Scan(&b.penulis)
	fmt.Print("Masukkan penerbit: ")
	fmt.Scan(&b.penerbit)
	fmt.Print("Masukkan tahun terbit: ")
	fmt.Scan(&b.tahunterbit)
	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scan(&b.jumlahHalaman)

	fmt.Println("\n=== BIODATA BUKU ===")
	fmt.Println("Judul Buku : ", b.judul)
	fmt.Println("Penulis: ", b.penulis)
	fmt.Println("Tahun Terbit: ", b.tahunterbit)
	fmt.Println("Jumlah Halaman: ", b.jumlahHalaman)
}
