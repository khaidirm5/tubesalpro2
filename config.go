package main

// STRUKTUR DATA UTAMA
type Pengeluaran struct {
	ID       int
	Nama     string
	Kategori string // "transportasi", "akomodasi", "makanan", "hiburan"
	Jumlah   int
}

const MaksData = 1000

// Variabel Global yang akan diakses oleh file crud.go dan sorting.go
var daftarPengeluaran [MaksData]Pengeluaran
var totalData int = 0
var budgetAwal int = 0