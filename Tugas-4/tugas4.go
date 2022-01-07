package main

import (
	"fmt"
)


func main() {
	// Soal 1
	var batasAtas int = 20
	for i := 1; i <= batasAtas; i++ {
		if i%2 == 0 { // Jika genap
			fmt.Printf("%d - Berkualitas\n", i)
		} else {
			if i%3 == 0 {
				fmt.Printf("%d - I Love Coding\n", i)
			} else {
				fmt.Printf("%d - Santai\n", i)
			}
		}
	}
	fmt.Println()

	// Soal 2	
	var tinggiSegitiga int = 7
	for i := 1; i <= tinggiSegitiga; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
	fmt.Println()

	// Soal 3	
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var idxDelete = [...]int{0, 1}

	kalimatSlice := kalimat[0:len(kalimat)]
	count := 0
	for _, idx := range idxDelete {
		for i := 0; i < len(kalimatSlice); i++ {
			if (idx - count) == i { // Index maju sesuai count
				kalimatSlice = append(kalimatSlice[:i], kalimatSlice[i+1:]...)
				count ++
            	break
			}
		}
	}

	fmt.Println(kalimatSlice)
	fmt.Println()

	// Soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam")
	sayuran = append(sayuran, "Buncis")
	sayuran = append(sayuran, "Kangkung")
	sayuran = append(sayuran, "Kubis")
	sayuran = append(sayuran, "Seledri")
	sayuran = append(sayuran, "Tauge")
	sayuran = append(sayuran, "Timun")

	for idx, sayur := range sayuran {
		fmt.Printf("%d. %s\n", (idx+1), sayur)
	}
	fmt.Println()

	// Soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar": 4,
		"tinggi": 6,
	}

	var volume int = 1
	for key, val := range satuan {
		volume = volume * val
		fmt.Printf("%s = %d\n", key, val)		
	}
	fmt.Printf("Volume Balok = %d\n", volume)
}
