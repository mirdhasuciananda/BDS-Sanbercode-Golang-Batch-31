package main

import (
	"fmt"
)

func luasPersegiPanjang (panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang (panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok (panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

func introduce (nama string, gender string, pekerjaan string, umur string) (kalimatPerkenalan string){
	if gender == "laki-laki" {
		kalimatPerkenalan = "Pak "
	} else if gender == "perempuan" {
		kalimatPerkenalan = "Bu "
	}
	
	kalimatPerkenalan = kalimatPerkenalan + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"

	return 
}

func buahFavorit (nama string, buahBuahan ...string) string {
	var kalimat string
	kalimat = "Halo nama saya " + nama + " dan buah favorit saya adalah"
	for idx, buah := range buahBuahan {		
		kalimat = kalimat + " " + `"` + buah + `"`
		if idx != len(buahBuahan) -1 {
			kalimat = kalimat + ","
		}
	}
	return kalimat
}

func main() {

	// Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
	fmt.Println()

	// Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun
	fmt.Println()

	// Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
	fmt.Println()

	// Soal 4
	var dataFilm = []map[string]string{} // Slice dengan tipe data map
	
	// Yang ini map dengan value slice
	// var dataFilm = map[string][]string{
	// 	"nama": []string{"LOTR", "avenger", "spiderman", "juon"},
	// 	"durasi": []string{"2 jam", "2 jam", "2 jam", "2 jam"},
	// 	"genre": []string{"action", "action", "action", "horror"},
	// 	"tahun": []string{"1999", "2019", "2004", "2004"},
	// }
	
	// Buatlah closure function disini
	var tambahDataFilm = func (nama string, durasi string, genre string, tahun string) {
		var film =  map[string]string {
			"nama": nama,
			"durasi": durasi,
			"genre": genre,
			"tahun": tahun,
		}		
		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}
