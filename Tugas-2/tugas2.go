package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	var text1, text2, text3, text4, text5 string
	text1 = "Bootcamp"
	text2 = "Digital"
	text3 = "Skill"
	text4 = "Sanbercode"
	text5 = "Golang"
	fmt.Println(text1 + " " + text2 + " " + text3 + " " + text4 + " " + text5)

	// Soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(halo)

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	var kataGabungan = kataPertama + " " + strings.Replace(kataKedua, "s", "S", 1) + " " + strings.Replace(kataKetiga, "r", "R", 1) + " " + strings.ToUpper(kataKeempat)
	fmt.Println(kataGabungan)

	// Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"
	angkaPertamaInt, _ := strconv.ParseInt(angkaPertama, 2, 8)
	angkaKeduaInt, _ := strconv.ParseInt(angkaKedua, 2, 8)
	angkaKetigaInt, _ := strconv.ParseInt(angkaKetiga, 2, 8)
	angkaKeempatInt, _ := strconv.ParseInt(angkaKeempat, 2, 8)
	angkaJumlah := angkaPertamaInt + angkaKeduaInt + angkaKetigaInt + angkaKeempatInt
	fmt.Println(angkaJumlah)

	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2021
	kalimat = strings.Replace(kalimat, "halo", "Hi", -1)
	angkaString := strconv.FormatInt(int64(angka), 10)
	fmt.Println("\"" + kalimat + "\"" + " - " + angkaString)
}
