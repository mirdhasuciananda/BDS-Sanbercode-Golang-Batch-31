package main

import (
  "fmt"
  "math"
  "strconv"
  "strings"
)

//----------------------------------
type segitigaSamaSisi struct {
  alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
  return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
  return s.alas * 3
}

type persegiPanjang struct {
  panjang, lebar int
}

func (p persegiPanjang) luas() int {
  return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
  return 2*p.panjang + 2*p.lebar
}

type tabung struct {
  jariJari, tinggi float64
}

func (t tabung) volume() float64 {
  return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
  return 2 * math.Pi * t.jariJari * t.tinggi
}

type balok struct {
  panjang, lebar, tinggi float64
}

func (b balok) volume() float64 {
  return b.panjang * b.lebar * b.tinggi
}

func (b balok) luasPermukaan() float64 {
  return 2 * (b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi)
}

//-------------------------------------------------
type phone struct {
  name, brand string
  year int
  colors []string
}

func (p phone) tampilPhone() string {
  var a string = "\nname: " + p.name +
    "\nbrand: " + p.brand +
    "\nyear: " + strconv.Itoa(p.year) +
    "\ncolors: " + strings.Join(p.colors, ", ")
  return a
}


//-------------------------------------------------
func luasPersegi(sisi int, kondisi bool) interface{} {
	var luas int = sisi*sisi
	var out interface{}	

	if kondisi == true {
		if sisi != 0 {
			out = fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
		} else {
			out = "Maaf anda belum menginput sisi dari persegi"
		}
	} else {
		if sisi != 0 {
			out = sisi
		} else {
			out = nil
		}
	}

	return out
}

//-------------------------------------------------




func main() {
  //Soal 1
  type hitungBangunDatar interface {
    luas() int
    keliling() int
  }

  type hitungBangunRuang interface {
    volume() float64
    luasPermukaan() float64
  }

  var bangunDatar hitungBangunDatar

  bangunDatar = segitigaSamaSisi{2, 4}
  fmt.Println("\n------------ Segitiga Sama Sisi")
  fmt.Println("luas:", bangunDatar.luas())
  fmt.Println("keliling :", bangunDatar.keliling())

  bangunDatar = persegiPanjang{4, 2}
  fmt.Println("\n------------- Persegi Panjang")
  fmt.Println("luas:", bangunDatar.luas())
  fmt.Println("keliling:", bangunDatar.keliling())

  var bangunRuang hitungBangunRuang

  bangunRuang = tabung{2, 4}
  fmt.Println("\n------------ Tabung")
  fmt.Println("volume:", bangunRuang.volume())
  fmt.Println("Luas Permukaan : ", bangunRuang.luasPermukaan())

  bangunRuang = balok{2, 4, 8}
  fmt.Println("\n------------ Balok")
  fmt.Println("Volume:", bangunRuang.volume())
  fmt.Println("Luas Permukaan:", bangunRuang.luasPermukaan())

  // Soal 2
  type tampil interface {
    tampilPhone() string
  }

  var hp tampil

  hp = phone{
    "Samsung Galaxy Note 20",
    "Samsung Galaxy Note 20",
    2020,
    []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
  }

  fmt.Println(hp.tampilPhone())
  fmt.Println()

  // Soal 3
  fmt.Println(luasPersegi(4, true))
  fmt.Println(luasPersegi(8, false))
  fmt.Println(luasPersegi(0, true))
  fmt.Println(luasPersegi(0, false))
  fmt.Println()

  // Soal 4  
  var prefix interface{} = "hasil penjumlahan dari "

  var kumpulanAngkaPertama interface{} = []int{6, 8}

  var kumpulanAngkaKedua interface{} = []int{12, 14}

  var tampilAngka1 = kumpulanAngkaPertama.([]int)
  var tampilAngka2 = kumpulanAngkaKedua.([]int)
  var sum = tampilAngka1[0] + tampilAngka1[1] + tampilAngka2[0] + tampilAngka2[1]

  fmt.Println(prefix, tampilAngka1[0], "+", tampilAngka1[1], "+", tampilAngka2[0], "+", tampilAngka2[1], "=", sum)



}