package library

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//----------------------------------
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2*p.Panjang + 2*p.Lebar
}

type Tabung struct {
	JariJari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * t.Tinggi
}

type Balok struct {
	Panjang, Lebar, Tinggi float64
}

func (b Balok) Volume() float64 {
	return b.Panjang * b.Lebar * b.Tinggi
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * (b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi)
}

//-------------------------------------------------
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

func (p Phone) TampilPhone() string {
	var a string = "\nname: " + p.Name +
		"\nbrand: " + p.Brand +
		"\nyear: " + strconv.Itoa(p.Year) +
		"\ncolors: " + strings.Join(p.Colors, ", ")
	return a
}

//-------------------------------------------------
func LuasPersegi(sisi int, kondisi bool) interface{} {
	var luas int = sisi * sisi
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
