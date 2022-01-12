package main

import (
	"fmt"
	"reflect"
)

type buah struct {
	nama string
	warna string
	biji string
	harga int16
}

type segitiga struct{
	alas, tinggi int
}

type persegi struct{
	sisi int
}

type persegiPanjang struct{
	panjang, lebar int
}

func (s segitiga) luasSegitiga() {
	luas := s.alas * s.tinggi / 2
	fmt.Printf("\nLuas segitiga = %d", luas)
}

func (p persegi) luasPersegi() {
	luas:= p.sisi * p.sisi
	fmt.Printf("\nLuas persegi = %d", luas)
}

func (pp persegiPanjang) luasPersegiPanjang() {
	luas:= pp.panjang * pp.lebar
	fmt.Printf("\nLuas persegi panjang = %d", luas)
}

type phone struct {
	name, brand string
	year int
	colors []string
}

func (p phone) tambahColor(tambahanColor *[]string) phone {
	p.colors = append(p.colors, *tambahanColor...)	
	return p
}

type Movie struct {
	Title, Genre string
	Duration, Year int16
}

func tambahDataFilm(judul string, durasi int16, genre string, tahun int16, dataFilm *[]Movie) {
	var film = Movie{
		Title: judul,
		Duration: durasi,
		Genre: genre,
		Year: tahun,
	}
	*dataFilm = append(*dataFilm, film)	
}

func main() {
	// Soal 1
	var buahNanas = buah{
		nama: "Nanas",
		warna: "Kuning",
		biji: "Tidak",
		harga: 9000,
	}
	var buahJeruk = buah{
		nama: "Jeruk",
		warna: "Oranye",
		biji: "Ada",
		harga: 8000,
	}
	var buahSemangka = buah{
		nama: "Semangka",
		warna: "Hijau & Merah",
		biji: "Ada",
		harga: 10000,
	}
	var buahPisang = buah{
		nama: "Pisang",
		warna: "Kuning",
		biji: "Tidak",
		harga: 5000,
	}

	fmt.Println(buahNanas)
	fmt.Println(buahJeruk)
	fmt.Println(buahSemangka)
	fmt.Println(buahPisang)
	fmt.Println()

	// Soal 2
	var segitigaku = segitiga{alas: 4, tinggi: 6}
	var persegiku = persegi{sisi: 10}
	var persegiPanjangku = persegiPanjang{panjang: 8, lebar: 3}

	segitigaku.luasSegitiga()
	persegiku.luasPersegi()
	persegiPanjangku.luasPersegiPanjang()
	fmt.Println()
	fmt.Println()

	// Soal 3
	var myPhone = phone{
		name: "M1",
		brand: "ASUS",
		year: 2019,
		colors: []string{"Biru"},
	}

	var tambahanColor = &[]string{"Hitam", "Coklat", "Putih"}
	myPhone = myPhone.tambahColor(tambahanColor)
	fmt.Println(myPhone)
	fmt.Println()

	// Soal 4
	var dataFilm = []Movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)	

	for no, s := range dataFilm {	// film adalah objek dari struct movie
		v := reflect.ValueOf(s)
		typeOfS := v.Type()

		fmt.Printf("%d. ", no+1)
		count := 0
		for i := 0; i< v.NumField(); i++ {			
			if count == 0 {
				fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
			} else {
				fmt.Printf("   %s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
			}
			count++
    	}
    	fmt.Println()
	}
}
