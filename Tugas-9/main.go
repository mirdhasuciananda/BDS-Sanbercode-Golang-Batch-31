package main

import (
	. "Tugas-9/library"
	"fmt"
)

func main() {
	//Soal 1
	type HitungBangunDatar interface {
		Luas() int
		Keliling() int
	}

	type HitungBangunRuang interface {
		Volume() float64
		LuasPermukaan() float64
	}

	var bangunDatar HitungBangunDatar

	bangunDatar = SegitigaSamaSisi{2, 4}
	fmt.Println("\n------------ Segitiga Sama Sisi")
	fmt.Println("luas:", bangunDatar.Luas())
	fmt.Println("keliling :", bangunDatar.Keliling())

	bangunDatar = PersegiPanjang{4, 2}
	fmt.Println("\n------------- Persegi Panjang")
	fmt.Println("luas:", bangunDatar.Luas())
	fmt.Println("keliling:", bangunDatar.Keliling())

	var bangunRuang HitungBangunRuang

	bangunRuang = Tabung{2, 4}
	fmt.Println("\n------------ Tabung")
	fmt.Println("volume:", bangunRuang.Volume())
	fmt.Println("Luas Permukaan : ", bangunRuang.LuasPermukaan())

	bangunRuang = Balok{2, 4, 8}
	fmt.Println("\n------------ Balok")
	fmt.Println("Volume:", bangunRuang.Volume())
	fmt.Println("Luas Permukaan:", bangunRuang.LuasPermukaan())

	// Soal 2
	type Tampil interface {
		TampilPhone() string
	}

	var hp Tampil

	hp = Phone{
		"Samsung Galaxy Note 20",
		"Samsung Galaxy Note 20",
		2020,
		[]string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(hp.TampilPhone())
	fmt.Println()

	// Soal 3
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))
	fmt.Println()

	// Soal 4
	var Prefix interface{} = "hasil penjumlahan dari "

	var KumpulanAngkaPertama interface{} = []int{6, 8}

	var KumpulanAngkaKedua interface{} = []int{12, 14}

	var tampilAngka1 = KumpulanAngkaPertama.([]int)
	var tampilAngka2 = KumpulanAngkaKedua.([]int)
	var sum = tampilAngka1[0] + tampilAngka1[1] + tampilAngka2[0] + tampilAngka2[1]

	fmt.Println(Prefix, tampilAngka1[0], "+", tampilAngka1[1], "+", tampilAngka2[0], "+", tampilAngka2[1], "=", sum)

}
