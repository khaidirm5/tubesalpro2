package main

import "fmt"

// Selection Sort Berdasarkan JUMLAH (Ascending)
func SortByJumlahSelection() {
	for i := 0; i < totalData-1; i++ {
		idxMin := i
		for j := i + 1; j < totalData; j++ {
			if daftarPengeluaran[j].Jumlah < daftarPengeluaran[idxMin].Jumlah {
				idxMin = j
			}
		}
		temp := daftarPengeluaran[i]
		daftarPengeluaran[i] = daftarPengeluaran[idxMin]
		daftarPengeluaran[idxMin] = temp
	}
	fmt.Println("✓ Daftar berhasil diurutkan berdasarkan JUMLAH terkecil.")
	TampilSemua() // Memanggil fungsi dari crud.go
}

// Insertion Sort Berdasarkan KATEGORI (Ascending)
func SortByKategoriInsertion() {
	for i := 1; i < totalData; i++ {
		key := daftarPengeluaran[i]
		j := i - 1
		for j >= 0 && daftarPengeluaran[j].Kategori > key.Kategori {
			daftarPengeluaran[j+1] = daftarPengeluaran[j]
			j = j - 1
		}
		daftarPengeluaran[j+1] = key
	}
	fmt.Println("✓ Daftar berhasil diurutkan berdasarkan KATEGORI (A-Z).")
	TampilSemua()
}