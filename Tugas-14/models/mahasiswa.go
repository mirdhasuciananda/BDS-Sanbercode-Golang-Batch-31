package models

type NilaiMahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"matakuliah"`
	IndeksNilai string `json:"indeksnilai"`
	Nilai       uint   `json:"nilai"`
}
