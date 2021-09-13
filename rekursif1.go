package main

import (
	"fmt"
	"unicode"
)

var jumlah = 0.0
var jumlahI = 0.0

func deret(i float64, j float64) float64 {

	if i <= 1 {
		fmt.Printf("%.2f", j)
		jumlah = jumlah + j
		return j
	}
	y := deret(i-1, j)
	jumlah = jumlah + j/i
	fmt.Printf(" %.2f", j/i)

	return y
}

func deretD(i float64) float64 {
	if i <= 1 {
		jumlahI = jumlahI + 1
		jumlah = jumlah + 2
		fmt.Print("1/2x")
		return i
	}
	y := deretD(i - 1)
	jumlahI = jumlahI + i
	x := i * i
	jumlah = jumlah + x
	fmt.Printf(" %.0f/%.0fx", i, x)

	return y
}

func deretE(i float64) float64 {
	if i <= 1 {
		jumlahI = jumlahI + 2
		jumlah = jumlah + 1
		fmt.Print("2x/1")
		return i
	}
	y := deretE(i - 1)
	x := i * i
	jumlahI = jumlahI + x
	jumlah = jumlah + i
	fmt.Printf(" %.0fx/%.0f", x, i)

	return y
}

func kapital(i int, x string) int {
	if i < 0 {
		return i
	}
	y := rune(x[i])
	if unicode.IsLower(y) {
		jumlah = jumlah + 1
	}

	return kapital(i-1, x)
}

func findAngka(i int, x string) int {
	if i < 0 {
		return i
	}
	y := rune(x[i])
	if unicode.IsNumber(y) {
		jumlah = jumlah + 1
	}
	return findAngka(i-1, x)
}

func main() {
	var pilih int
	var batas float64
	var angka float64
	var kata string

	fmt.Println("===PILIHAN===")
	fmt.Println("1. Deret Pembagian angka (a, b, c)")
	fmt.Println("2. Deret Pembagian angka (d)")
	fmt.Println("3. Deret Pembagian angka (e)")
	fmt.Println("4. Hitung huruf")
	fmt.Println("5. Hitung angka pada suatu kata")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		fmt.Print("Masukkan batasan: ")
		fmt.Scan(&batas)
		fmt.Print("Masukkan angka pertama: ")
		fmt.Scan(&angka)

		deret(batas, angka)
		fmt.Printf(" = %.2f", jumlah)
	case 2:
		fmt.Print("Masukkan batasan: ")
		fmt.Scan(&batas)

		deretD(batas)
		fmt.Printf(" = %.0f/%.0fx", jumlahI, jumlah)
	case 3:
		fmt.Print("Masukkan batasan: ")
		fmt.Scan(&batas)

		deretE(batas)
		fmt.Printf(" = %.0fx/%.0f", jumlahI, jumlah)
	case 4:
		fmt.Print("Masukkan sebuah kata: ")
		fmt.Scan(&kata)

		kapital(len(kata)-1, kata)
		fmt.Printf("jumlah huruf kecil: %.0f", jumlah)
	case 5:
		fmt.Print("Masukkan sebuah kata(JKT48): ")
		fmt.Scan(&kata)

		findAngka(len(kata)-1, kata)
		fmt.Printf("Jumlah angka : %.0f", jumlah)
	}
}
