package main

import "fmt"

func main() {
	// Mencetak angka dalam format decimal (basis 10)
	angka := 42
	fmt.Println("Angka dalam format desimal:", angka)

	// Konversi dan cetak dalam format biner (basis 2)
	fmt.Printf("Angka dalam format biner: %b\n", angka)

	// Konversi dan cetak dalam format heksadesimal (basis 16)
	fmt.Printf("Angka dalam format heksadesimal: %x\n", angka)

	// Konversi dan cetak dalam format floating point
	fmt.Printf("Angka dalam format floating point: %f\n", float64(angka))
}
