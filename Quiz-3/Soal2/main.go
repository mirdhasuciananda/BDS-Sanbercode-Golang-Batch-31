package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

type SegitigaSamaSisi struct {
	Alas   float64 `json:"alas"`
	Tinggi float64 `json:"tinggi"`
	Hitung float64 `json:"hitung"`
}

type Persegi struct {
	Sisi   float64 `json:"sisi"`
	Hitung float64 `json:"hitung"`
}

type PersegiPanjang struct {
	Panjang float64 `json:"panjang"`
	Lebar   float64 `json:"lebar"`
	Hitung  float64 `json:"hitung"`
}

type Lingkaran struct {
	JariJari float64 `json:"jarijari"`
	Hitung   float64 `json:"hitung"`
}

type JajarGenjang struct {
	Sisi   float64 `json:"sisi"`
	Alas   float64 `json:"alas"`
	Tinggi float64 `json:"tinggi"`
	Hitung float64 `json:"hitung"`
}

var nilaiSegitigaSamaSisi = []SegitigaSamaSisi{}
var nilaiPersegi = []Persegi{}
var nilaiPersegiPanjang = []PersegiPanjang{}
var nilaiLingkaran = []Lingkaran{}
var nilaiJajarGenjang = []JajarGenjang{}

func LuasSegitigaSamaSisi(a float64, t float64, ch chan float64) {
	ch <- a * t / 2
}

func KelSegitigaSamaSisi(a float64, ch chan float64) {
	ch <- 3 * a
}

func LuasPersegi(s float64, ch chan float64) {
	ch <- s * s
}

func KelPersegi(s float64, ch chan float64) {
	ch <- 4 * s
}

func LuasPersegiPanjang(p float64, l float64, ch chan float64) {
	ch <- p * l
}

func KelPersegiPanjang(p float64, l float64, ch chan float64) {
	ch <- 2 * (p + l)
}

func LuasLingkaran(r float64, ch chan float64) {
	ch <- math.Pi * math.Pow(r, 2)
}

func KelLingkaran(r float64, ch chan float64) {
	ch <- math.Pi * 2 * r
}

func LuasJajarGenjang(a float64, t float64, ch chan float64) {
	ch <- a * t
}

func KelJajarGenjang(a float64, s float64, ch chan float64) {
	ch <- 2 * (a * s)
}

func getNilaiSegitigaSamaSisi(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilaiSegitigaSamaSisi, err := json.Marshal(nilaiSegitigaSamaSisi)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(dataNilaiSegitigaSamaSisi)
		return
	}
	http.Error(rw, "method not allowed", http.StatusNotFound)
	return
}

func getNilaiPersegi(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilaiPersegi, err := json.Marshal(nilaiPersegi)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(dataNilaiPersegi)
		return
	}
	http.Error(rw, "method not allowed", http.StatusNotFound)
	return
}

func getNilaiPersegiPanjang(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilaiPersegiPanjang, err := json.Marshal(nilaiPersegiPanjang)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(dataNilaiPersegiPanjang)
		return
	}
	http.Error(rw, "method not allowed", http.StatusNotFound)
	return
}

func getNilaiLingkaran(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilaiLingkaran, err := json.Marshal(nilaiLingkaran)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(dataNilaiLingkaran)
		return
	}
	http.Error(rw, "method not allowed", http.StatusNotFound)
	return
}

func getNilaiJajarGenjang(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilaiJajarGenjang, err := json.Marshal(nilaiJajarGenjang)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(dataNilaiJajarGenjang)
		return
	}
	http.Error(rw, "method not allowed", http.StatusNotFound)
	return
}

func postNilaiLuasSegitigaSamaSisi(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var hitungSegitiga = SegitigaSamaSisi{}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&hitungSegitiga); err != nil {
				log.Fatal(err)
			}
		} else {
			alasS := r.FormValue("alas")
			tinggiS := r.FormValue("tinggi")
			getAlas, _ := strconv.Atoi(alasS)
			getTinggi, _ := strconv.Atoi(tinggiS)

			hitungSegitiga.Alas = float64(getAlas)
			hitungSegitiga.Tinggi = float64(getTinggi)

		}
		if hitungSegitiga.Alas < 0 || hitungSegitiga.Tinggi < 0 {
			http.Error(rw, "Nilai tidak boleh bernilai tinggi", http.StatusBadRequest)
			return
		}

		getValue := make(chan float64)
		go LuasSegitigaSamaSisi(hitungSegitiga.Alas, hitungSegitiga.Tinggi, getValue)
		getLuas := <-getValue
		hitungSegitiga.Hitung = getLuas

		dataSegitiga, _ := json.Marshal(hitungSegitiga)
		nilaiSegitigaSamaSisi = append(nilaiSegitigaSamaSisi, hitungSegitiga)
		rw.Write(dataSegitiga)
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

func postNilaiKelSegitigaSamaSisi(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var hitungSegitiga = SegitigaSamaSisi{}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&hitungSegitiga); err != nil {
				log.Fatal(err)
			}
		} else {
			alasS := r.FormValue("alas")
			tinggiS := r.FormValue("tinggi")
			getAlas, _ := strconv.Atoi(alasS)
			getTinggi, _ := strconv.Atoi(tinggiS)

			hitungSegitiga.Alas = float64(getAlas)
			hitungSegitiga.Tinggi = float64(getTinggi)

		}
		if hitungSegitiga.Alas < 0 || hitungSegitiga.Tinggi < 0 {
			http.Error(rw, "Nilai tidak boleh bernilai tinggi", http.StatusBadRequest)
			return
		}

		getValue := make(chan float64)

		go KelSegitigaSamaSisi(hitungSegitiga.Alas, getValue)
		getKel := <-getValue
		hitungSegitiga.Hitung = getKel

		dataSegitiga, _ := json.Marshal(hitungSegitiga)
		nilaiSegitigaSamaSisi = append(nilaiSegitigaSamaSisi, hitungSegitiga)
		rw.Write(dataSegitiga)
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

func postNilaiLuasPersegi(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var hitungPersegi = Persegi{}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&hitungPersegi); err != nil {
				log.Fatal(err)
			}
		} else {
			SisiP := r.FormValue("sisi")
			getSisi, _ := strconv.Atoi(SisiP)

			hitungPersegi.Sisi = float64(getSisi)

		}
		if hitungPersegi.Sisi < 0 {
			http.Error(rw, "Nilai tidak boleh bernilai tinggi", http.StatusBadRequest)
			return
		}

		getValue := make(chan float64)
		go LuasPersegi(hitungPersegi.Sisi, getValue)
		getLuas := <-getValue
		hitungPersegi.Hitung = getLuas

		dataPersegi, _ := json.Marshal(hitungPersegi)
		nilaiPersegi = append(nilaiPersegi, hitungPersegi)
		rw.Write(dataPersegi)
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

func postNilaiKelPersegi(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var hitungPersegi = Persegi{}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&hitungPersegi); err != nil {
				log.Fatal(err)
			}
		} else {
			SisiP := r.FormValue("sisi")
			getSisi, _ := strconv.Atoi(SisiP)

			hitungPersegi.Sisi = float64(getSisi)

		}
		if hitungPersegi.Sisi < 0 {
			http.Error(rw, "Nilai tidak boleh bernilai tinggi", http.StatusBadRequest)
			return
		}

		getValue := make(chan float64)

		go KelPersegi(hitungPersegi.Sisi, getValue)
		getKel := <-getValue
		hitungPersegi.Hitung = getKel

		dataPersegi, _ := json.Marshal(hitungPersegi)
		nilaiPersegi = append(nilaiPersegi, hitungPersegi)
		rw.Write(dataPersegi)
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

func postNilaiLuasPersegiPanjang(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var hitungPersegiPanjang = PersegiPanjang{}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&hitungPersegiPanjang); err != nil {
				log.Fatal(err)
			}
		} else {
			panjangPP := r.FormValue("panjang")
			lebarPP := r.FormValue("lebar")
			getPanjang, _ := strconv.Atoi(panjangPP)
			getLebar, _ := strconv.Atoi(lebarPP)

			hitungPersegiPanjang.Panjang = float64(getPanjang)
			hitungPersegiPanjang.Lebar = float64(getLebar)

		}
		if hitungPersegiPanjang.Panjang < 0 || hitungPersegiPanjang.Lebar < 0 {
			http.Error(rw, "Nilai tidak boleh bernilai tinggi", http.StatusBadRequest)
			return
		}

		getValue := make(chan float64)
		go LuasPersegiPanjang(hitungPersegiPanjang.Panjang, hitungPersegiPanjang.Lebar, getValue)
		getLuas := <-getValue
		hitungPersegiPanjang.Hitung = getLuas

		dataPersegiPanjang, _ := json.Marshal(hitungPersegiPanjang)
		nilaiPersegiPanjang = append(nilaiPersegiPanjang, hitungPersegiPanjang)
		rw.Write(dataPersegiPanjang)
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

func postNilaiKelPersegiPanjang(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var hitungPersegiPanjang = PersegiPanjang{}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&hitungPersegiPanjang); err != nil {
				log.Fatal(err)
			}
		} else {
			panjangPP := r.FormValue("panjang")
			lebarPP := r.FormValue("lebar")
			getPanjang, _ := strconv.Atoi(panjangPP)
			getLebar, _ := strconv.Atoi(lebarPP)

			hitungPersegiPanjang.Panjang = float64(getPanjang)
			hitungPersegiPanjang.Lebar = float64(getLebar)

		}
		if hitungPersegiPanjang.Panjang < 0 || hitungPersegiPanjang.Lebar < 0 {
			http.Error(rw, "Nilai tidak boleh bernilai tinggi", http.StatusBadRequest)
			return
		}

		getValue := make(chan float64)

		go KelPersegiPanjang(hitungPersegiPanjang.Panjang, hitungPersegiPanjang.Lebar, getValue)
		getKel := <-getValue
		hitungPersegiPanjang.Hitung = getKel

		dataPersegiPanjang, _ := json.Marshal(hitungPersegiPanjang)
		nilaiPersegiPanjang = append(nilaiPersegiPanjang, hitungPersegiPanjang)
		rw.Write(dataPersegiPanjang)
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

func postNilaiLuasLingkaran(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var hitungLingkaran = Lingkaran{}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&hitungLingkaran); err != nil {
				log.Fatal(err)
			}
		} else {
			jari_Jari := r.FormValue("jarijari")
			getJarijari, _ := strconv.Atoi(jari_Jari)

			hitungLingkaran.JariJari = float64(getJarijari)
		}
		if hitungLingkaran.JariJari < 0 {
			http.Error(rw, "Nilai tidak boleh bernilai tinggi", http.StatusBadRequest)
			return
		}

		getValue := make(chan float64)
		go LuasLingkaran(hitungLingkaran.JariJari, getValue)
		getLuas := <-getValue
		hitungLingkaran.Hitung = getLuas

		dataLingkaran, _ := json.Marshal(hitungLingkaran)
		nilaiLingkaran = append(nilaiLingkaran, hitungLingkaran)
		rw.Write(dataLingkaran)
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

func postNilaiKelLingkaran(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var hitungLingkaran = Lingkaran{}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&hitungLingkaran); err != nil {
				log.Fatal(err)
			}
		} else {
			jari_Jari := r.FormValue("jarijari")
			getJarijari, _ := strconv.Atoi(jari_Jari)

			hitungLingkaran.JariJari = float64(getJarijari)
		}
		if hitungLingkaran.JariJari < 0 {
			http.Error(rw, "Nilai tidak boleh bernilai tinggi", http.StatusBadRequest)
			return
		}

		getValue := make(chan float64)

		go KelLingkaran(hitungLingkaran.JariJari, getValue)
		getKel := <-getValue
		hitungLingkaran.Hitung = getKel

		dataLingkaran, _ := json.Marshal(hitungLingkaran)
		nilaiLingkaran = append(nilaiLingkaran, hitungLingkaran)
		rw.Write(dataLingkaran)
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

func postNilaiLuasJajarGenjang(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var hitungJajarGenjang = JajarGenjang{}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&hitungJajarGenjang); err != nil {
				log.Fatal(err)
			}
		} else {
			sisiJG := r.FormValue("sisi")
			alasJG := r.FormValue("alas")
			tinggiJG := r.FormValue("tinggi")
			getSisiJG, _ := strconv.Atoi(sisiJG)
			getAlasJG, _ := strconv.Atoi(alasJG)
			getTinggiJG, _ := strconv.Atoi(tinggiJG)

			hitungJajarGenjang.Sisi = float64(getSisiJG)
			hitungJajarGenjang.Alas = float64(getAlasJG)
			hitungJajarGenjang.Tinggi = float64(getTinggiJG)
		}
		if hitungJajarGenjang.Sisi < 0 || hitungJajarGenjang.Alas < 0 || hitungJajarGenjang.Tinggi < 0 {
			http.Error(rw, "Nilai tidak boleh bernilai tinggi", http.StatusBadRequest)
			return
		}

		getValue := make(chan float64)
		go LuasJajarGenjang(hitungJajarGenjang.Alas, hitungJajarGenjang.Tinggi, getValue)
		getLuas := <-getValue
		hitungJajarGenjang.Hitung = getLuas

		dataJajarGenjang, _ := json.Marshal(hitungJajarGenjang)
		nilaiJajarGenjang = append(nilaiJajarGenjang, hitungJajarGenjang)
		rw.Write(dataJajarGenjang)
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

func postNilaiKelJajarGenjang(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var hitungJajarGenjang = JajarGenjang{}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&hitungJajarGenjang); err != nil {
				log.Fatal(err)
			}
		} else {
			sisiJG := r.FormValue("sisi")
			alasJG := r.FormValue("alas")
			tinggiJG := r.FormValue("tinggi")
			getSisiJG, _ := strconv.Atoi(sisiJG)
			getAlasJG, _ := strconv.Atoi(alasJG)
			getTinggiJG, _ := strconv.Atoi(tinggiJG)

			hitungJajarGenjang.Sisi = float64(getSisiJG)
			hitungJajarGenjang.Alas = float64(getAlasJG)
			hitungJajarGenjang.Tinggi = float64(getTinggiJG)
		}
		if hitungJajarGenjang.Sisi < 0 || hitungJajarGenjang.Alas < 0 || hitungJajarGenjang.Tinggi < 0 {
			http.Error(rw, "Nilai tidak boleh bernilai tinggi", http.StatusBadRequest)
			return
		}

		getValue := make(chan float64)

		go KelJajarGenjang(hitungJajarGenjang.Alas, hitungJajarGenjang.Sisi, getValue)
		getKel := <-getValue
		hitungJajarGenjang.Hitung = getKel

		dataJajarGenjang, _ := json.Marshal(hitungJajarGenjang)
		nilaiJajarGenjang = append(nilaiJajarGenjang, hitungJajarGenjang)
		rw.Write(dataJajarGenjang)
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

//===========================================================================================
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			rw.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if username == "admin" && password == "admin" {
			next.ServeHTTP(rw, r)
			return
		}
		rw.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func main() {

	http.HandleFunc("/nilai-luas-segitiga", getNilaiSegitigaSamaSisi)
	http.Handle("/nilai-luas-segitiga/create", Auth(http.HandlerFunc(postNilaiLuasSegitigaSamaSisi)))

	http.HandleFunc("/nilai-kel-segitiga", getNilaiSegitigaSamaSisi)
	http.Handle("/nilai-kel-segitiga/create", Auth(http.HandlerFunc(postNilaiKelSegitigaSamaSisi)))

	http.HandleFunc("/nilai-luas-persegi", getNilaiPersegi)
	http.Handle("/nilai-luas-persegi/create", Auth(http.HandlerFunc(postNilaiLuasPersegi)))

	http.HandleFunc("/nilai-kel-persegi", getNilaiPersegi)
	http.Handle("/nilai-kel-persegi/create", Auth(http.HandlerFunc(postNilaiKelPersegi)))

	http.HandleFunc("/nilai-luas-persegi-panjang", getNilaiPersegiPanjang)
	http.Handle("/nilai-luas-persegi-panjang/create", Auth(http.HandlerFunc(postNilaiLuasPersegiPanjang)))

	http.HandleFunc("/nilai-kel-persegi-panjang", getNilaiPersegiPanjang)
	http.Handle("/nilai-kel-persegi-panjang/create", Auth(http.HandlerFunc(postNilaiKelPersegiPanjang)))

	http.HandleFunc("/nilai-luas-lingkaran", getNilaiLingkaran)
	http.Handle("/nilai-luas-lingkaran/create", Auth(http.HandlerFunc(postNilaiLuasLingkaran)))

	http.HandleFunc("/nilai-kel-lingkaran", getNilaiLingkaran)
	http.Handle("/nilai-kel-lingkaran/create", Auth(http.HandlerFunc(postNilaiKelLingkaran)))

	http.HandleFunc("/nilai-luas-jajargenjang", getNilaiJajarGenjang)
	http.Handle("/nilai-luas-jajargenjang/create", Auth(http.HandlerFunc(postNilaiLuasJajarGenjang)))

	http.HandleFunc("/nilai-kel-jajargenjang", getNilaiJajarGenjang)
	http.Handle("/nilai-kel-jajargenjang/create", Auth(http.HandlerFunc(postNilaiKelJajarGenjang)))

	fmt.Println("server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
