package main

import (
	"fmt"
	"strconv"
)

func hitungIndeks(nilai int) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 60 {
		return "C"
	} else if nilai >= 50 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	// Soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjangPersegiPanjangInt, _ := strconv.Atoi(panjangPersegiPanjang)
	lebarPersegiPanjangInt, _ := strconv.Atoi(lebarPersegiPanjang)

	alasSegitigaInt, _ := strconv.Atoi(alasSegitiga)
	tinggiSegitigaInt, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjangPersegiPanjangInt * lebarPersegiPanjangInt
	var kelilingPersegiPanjang int = 2 * (panjangPersegiPanjangInt + lebarPersegiPanjangInt)
	var luasSegitiga int = alasSegitigaInt * tinggiSegitigaInt / 2

	fmt.Println("Luas persegi panjang = " + strconv.Itoa(luasPersegiPanjang))
	fmt.Println("Keliling persegi panjang = " + strconv.Itoa(kelilingPersegiPanjang))
	fmt.Println("Luas segitiga = " + strconv.Itoa(luasSegitiga))

	// Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	indeksJohn := hitungIndeks(nilaiJohn)
	indeksDoe := hitungIndeks(nilaiDoe)

	fmt.Printf("\nIndeks John = %s", indeksJohn)
	fmt.Printf("\nIndeks Doe = %s\n", indeksDoe)

	// Soal 3
	var tanggal = 7
	var bulan = 8
	var tahun = 1999

	var bulanStr string
	switch bulan {
	case 1:
		bulanStr = "Januari"
	case 2:
		bulanStr = "Februari"
	case 3:
		bulanStr = "Maret"
	case 4:
		bulanStr = "April"
	case 5:
		bulanStr = "Mei"
	case 6:
		bulanStr = "Juni"
	case 7:
		bulanStr = "Juli"
	case 8:
		bulanStr = "Agustus"
	case 9:
		bulanStr = "September"
	case 10:
		bulanStr = "Oktober"
	case 11:
		bulanStr = "November"
	case 12:
		bulanStr = "Desember"
	}

	fmt.Printf("\n%d %s %d\n", tanggal, bulanStr, tahun)

	// Soal 4
	var generasi string
	if tahun >= 1994 && tahun <= 1964 {
		generasi = "Baby boomer"
	} else if tahun >= 1965 && tahun <= 1979 {
		generasi = "Generasi X"
	} else if tahun >= 1980 && tahun <= 1994 {
		generasi = "Generasi Y (Millenials)"
	} else if tahun >= 1995 && tahun <= 2015 {
		generasi = "Generasi Z"
	}

	fmt.Printf("\nGenerasi: %s", generasi)
}
