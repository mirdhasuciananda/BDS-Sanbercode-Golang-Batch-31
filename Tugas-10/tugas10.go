package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

// Soal 1
func myFunc(kalimat string, tahun int) {
	fmt.Println(kalimat)
	fmt.Println(tahun)
}
func runApps() {
	defer myFunc("Golang Backend Development", 2022)
	fmt.Println("Run Apps")
}

// Soal 2
func cekKondisi(sisi int, kondisi bool) (string, error) {
	if sisi != 0 && kondisi == true {
		return fmt.Sprintf("Keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm\n", sisi, 3*sisi), nil
	} else if sisi != 0 && kondisi == false {
		return fmt.Sprintln(sisi), nil
	} else if sisi == 0 && kondisi == true {
		return "0", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else {
		panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}
}

func kelilingSegitigaSamaSisi(sisi int, kondisi bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()

	kalimat, err := cekKondisi(sisi, kondisi)
	if err == nil {
		fmt.Print(kalimat)
	} else {
		fmt.Println(err.Error())
	}
}

// Soal 3
func tambahAngka(tambahan int, angka *int) int {
	*angka = *angka + tambahan
	return *angka
}

func cetakAngka(angka *int) {
	fmt.Println(*angka)
}

// Soal 4
func tambahData(data *[]string, tambahan string) *[]string {
	*data = append(*data, tambahan)
	return data
}

// Soal 5
func luasLingkaran(radius float64) float64 {
	return math.Round(math.Pi * math.Pow(radius, 2))
}

func kelilingLingkaran(radius float64) float64 {
	return math.Round(2 * math.Pi * radius)
}

// Soal 6
func luasPersegiPanjang(panjang, lebar float64) float64 {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar float64) float64 {
	return 2 * (panjang + lebar)
}

func main() {
	angka := 1

	// Soal 1
	runApps()
	fmt.Println()

	// Soal 2
	kelilingSegitigaSamaSisi(4, true)
	kelilingSegitigaSamaSisi(8, false)
	kelilingSegitigaSamaSisi(0, true)
	kelilingSegitigaSamaSisi(0, false)
	fmt.Println()

	// Soal 3
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// Soal 4
	var phones = []string{}

	phones = *tambahData(&phones, "Xiaomi")
	phones = *tambahData(&phones, "Asus")
	phones = *tambahData(&phones, "IPhone")
	phones = *tambahData(&phones, "Samsung")
	phones = *tambahData(&phones, "Oppo")
	phones = *tambahData(&phones, "Vivo")

	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second * 1)
	}
	fmt.Println()

	// Soal 5
	var radiusPertama float64 = 7
	var radiusKedua float64 = 10
	var radiusKetiga float64 = 15

	fmt.Printf("\nLingkaran pertama\n")
	fmt.Printf("Luas\t\t: %f\n", luasLingkaran(radiusPertama))
	fmt.Printf("Keliling\t: %f\n", kelilingLingkaran(radiusPertama))

	fmt.Printf("\nLingkaran kedua\n")
	fmt.Printf("Luas\t\t: %f\n", luasLingkaran(radiusKedua))
	fmt.Printf("Keliling\t: %f\n", kelilingLingkaran(radiusKedua))

	fmt.Printf("\nLingkaran ketiga\n")
	fmt.Printf("Luas\t\t: %f\n", luasLingkaran(radiusKetiga))
	fmt.Printf("Keliling\t: %f\n", kelilingLingkaran(radiusKetiga))
	fmt.Println()

	// Soal 6
	var panjang = flag.Float64("panjang", 8, "type the length")
	var lebar = flag.Float64("lebar", 6, "type the width")

	flag.Parse()
	fmt.Printf("Panjang\t\t: %f\n", *panjang)
	fmt.Printf("Lebar\t\t: %f\n", *lebar)
	fmt.Printf("Luas\t\t: %f\n", luasPersegiPanjang(*panjang, *lebar))
	fmt.Printf("Keliling\t: %f\n", kelilingPersegiPanjang(*panjang, *lebar))
	fmt.Println()
}
