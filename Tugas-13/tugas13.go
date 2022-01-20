package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilaiMahasiswa, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilaiMahasiswa)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func postNilaiMahasiswa(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var newNilaiMahasiswa NilaiMahasiswa

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&newNilaiMahasiswa); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			nama := r.PostFormValue("Nama")
			mataKuliah := r.PostFormValue("MataKuliah")
			getNilai := r.PostFormValue("Nilai")
			nilai, _ := strconv.Atoi(getNilai)
			id := len(nilaiNilaiMahasiswa) + 1

			var indeksNilai string
			if nilai <= 100 && nilai >= 0 {
				if nilai >= 80 {
					indeksNilai = "A"
				} else if nilai >= 70 {
					indeksNilai = "B"
				} else if nilai >= 60 {
					indeksNilai = "C"
				} else if nilai >= 50 {
					indeksNilai = "D"
				} else {
					indeksNilai = "E"
				}
			} else {
				// fmt.Println("Nilai haruslah dalam rentang 0-100")
				rw.Write([]byte("Nilai haruslah dalam rentang 0-100"))
				return
			}

			newNilaiMahasiswa = NilaiMahasiswa{
				Nama:        nama,
				MataKuliah:  mataKuliah,
				IndeksNilai: indeksNilai,
				Nilai:       uint(nilai),
				ID:          uint(id),
			}
		}

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilaiMahasiswa)
		dataNilaiMahasiswa, _ := json.Marshal(newNilaiMahasiswa) // to byte
		rw.Write(dataNilaiMahasiswa)                             // cetak di browser
		return
	}
	http.Error(rw, "NOT FOUND", http.StatusNotFound)
	return
}

// Fungi Log yang berguna sebagai middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			rw.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if username == "admin" && password == "login" {
			next.ServeHTTP(rw, r)
			return
		}
		rw.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func main() {
	fmt.Println(nilaiNilaiMahasiswa)
	http.HandleFunc("/nilai-mahasiswa", getNilaiMahasiswa)
	http.Handle("/post-nilai-mahasiswa", Auth(http.HandlerFunc(postNilaiMahasiswa)))

	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
