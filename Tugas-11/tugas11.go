package main

import "fmt"
import (
	"math"
	"sort"
	"sync"
	"time"
)

// Soal 1
func printPhones(data []string, wg *sync.WaitGroup) {
	sort.Strings(data)
	for i, phone := range data {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}

// Soal 2
func getMovies(ch chan string, movies []string) {
	for i, movie := range movies {
		ch <- movie
		fmt.Printf("%d. %s\n", i+1, movie)
	}
	close(ch)
}

// Soal 3
func luasLingkaran(r float64, ch chan float64) {
	ch <- math.Pi * math.Pow(r, 2)
}

func kelLingkaran(r float64, ch chan float64) {
	ch <- math.Pi * 2 * r
}

func volTabung(r float64, ch chan float64) {
	ch <- math.Pi * 2 * r * 10
}

// Soal 4

func luasPP(p int, l int, ch chan int) {
	ch <- p * l
}

func kelPP(p int, l int, ch chan int) {
	ch <- 2 * (p + l)
}

func volBalok(p int, l int, t int, ch chan int) {
	ch <- p * l * t
}

func main() {
	//Soal 1
	var wg sync.WaitGroup
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	wg.Add(1)
	go printPhones(phones, &wg)

	wg.Wait()

	fmt.Println()

	// Soal 2

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies)

	for {
		_, ok := <-moviesChannel
		if ok == false {
			break
		}
	}
	fmt.Println()

	// Soal 3

	getValue := make(chan float64)
	var radius = []float64{8, 14, 20}
	for _, rad := range radius {
		go luasLingkaran(rad, getValue)
		nilaiLuas := <-getValue
		fmt.Printf("Luas Lingkaran dengan jari-jari %f adalah %f\n", rad, nilaiLuas)

		go kelLingkaran(rad, getValue)
		nilaiKel := <-getValue
		fmt.Printf("Keliling Lingkaran dengan jari-jari %f adalah %f\n", rad, nilaiKel)

		go volTabung(rad, getValue)
		nilaiVol := <-getValue
		fmt.Printf("Volume tabung dengan jari-jari alas %f dan tinggi 10 adalah %f\n", rad, nilaiVol)

		fmt.Println()
	}
	fmt.Println()

	// Soal 4
	var hasilLuas = make(chan int)
	go luasPP(2, 4, hasilLuas)

	var hasilKel = make(chan int)
	go kelPP(2, 4, hasilKel)

	var hasilVol = make(chan int)
	go volBalok(2, 4, 6, hasilVol)

	for i := 0; i < 3; i++ {
		select {
		case luas := <-hasilLuas:
			fmt.Printf("Luas persegi panjang adalah %d\t\n", luas)
		case keliling := <-hasilKel:
			fmt.Printf("Keliling persegi panjang adalah %d\t\n", keliling)
		case volume := <-hasilVol:
			fmt.Printf("Volume balok adalah %d\t\n", volume)
		}
	}
}
