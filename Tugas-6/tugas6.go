package main

import (
	"fmt"
	"math"
	"strings"
)

func ubahNilai(luasOriginal *float64, kelilingOriginal *float64, radius float64) {
	*luasOriginal = math.Pi * math.Pow(radius, 2)
	*kelilingOriginal = 2* math.Pi * radius
}

func introduce(sentence *string, nama string, gender string, pekerjaan string, umur string) {
	if strings.ToLower(gender) == "laki-laki" {
		*sentence = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	} else if strings.ToLower(gender)  == "perempuan" {
		*sentence = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}
}

func tambahData(buah []string, buahBuahan *[]string) []string {
	buah = append(buah, *buahBuahan...)
	return buah
}

func main() {
	// Soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64

	var radius float64 = 5
	var newRadius *float64 = &radius	

	luasLingkaran = math.Pi * math.Pow(radius, 2)
  	kelilingLingkaran = 2 * math.Pi * radius
  	
  	radius = 10

  	fmt.Printf("Luas awal: %f\n", luasLingkaran)
  	fmt.Printf("Keliling awal: %f\n", kelilingLingkaran)
  
  	ubahNilai(&luasLingkaran, &kelilingLingkaran, *newRadius)
  	fmt.Printf("\nLuas diganti: %f\n", luasLingkaran)
  	fmt.Printf("Keliling diganti: %f\n", kelilingLingkaran)
  	fmt.Println()

  	// Soal 2
  	var sentence string

  	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	fmt.Println()

	// Soal 3		
	var buah = []string{}
	var buahBuahan = &[]string{"Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat"}

	buah = tambahData(buah, buahBuahan)

	for i, buah := range buah {
		fmt.Printf("%d. %s\n", (i+1), buah)
	}
	fmt.Println()

	// Soal 4
	var dataFilm = []map[string]string{}

	var tambahDataFilm = func (nama string, durasi string, genre string, tahun string, dataFilm *[]map[string]string) {
		var film =  map[string]string {
			"nama": nama,
			"durasi": durasi,
			"genre": genre,
			"tahun": tahun,
		}		
		*dataFilm = append(*dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	keys := make([]string, 0, len(dataFilm))	

	for i, item := range dataFilm {
		// fmt.Println(item)
		for key := range item {
			keys = append(keys, key)
		}

		fmt.Printf("%d. ", i+1)
		count := 0
		for j, nestedItem := range item {			
			if count == 0 {
				fmt.Printf("%s : %s\n", j, nestedItem)
			} else {
				fmt.Printf("   %s : %s\n", j, nestedItem)
			}
			count++
		}
		fmt.Println()
	}
}