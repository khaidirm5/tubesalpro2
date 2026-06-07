package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
    fmt.Print(prompt)
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    return strings.TrimSpace(text)
}

// Set Budget Awal
func SetBudget() {
    fmt.Print("Masukkan Total Budget Perjalanan (Rp): ")
    fmt.Scan(&budgetAwal)
    fmt.Println("✓ Budget awal berhasil diatur!")
    fmt.Println("Budget saat ini: Rp", budgetAwal)
}

// Tambah Data Pengeluaran (CREATE)
func TambahPengeluaran() {
	if totalData >= MaksData {
		fmt.Println("Kapasitas penyimpanan penuh!")
		return
	}

	var p Pengeluaran
	p.ID = totalData + 1

	p.Nama = readLine("Masukkan Nama Pengeluaran: ")
	p.Kategori = readLine("Masukkan Kategori (transportasi/akomodasi/makanan/hiburan): ")
	fmt.Print("Masukkan Jumlah Pengeluaran (Rp): ")
	fmt.Scan(&p.Jumlah)
    
	daftarPengeluaran[totalData] = p
	totalData++
	fmt.Println("✓ Data pengeluaran berhasil ditambahkan!")
}

// Tampilkan Semua Pengeluaran (READ)
func TampilSemua() {
	if totalData == 0 {
		fmt.Println(" Data pengeluaran masih kosong.")
		return
	}

	fmt.Println("\n=========================================================")
	fmt.Printf("%-4s | %-20s | %-15s | %-12s\n", "ID", "Nama", "Kategori", "Jumlah (Rp)")
	fmt.Println("=========================================================")
	for i := 0; i < totalData; i++ {
		p := daftarPengeluaran[i]
		fmt.Printf("%-4d | %-20s | %-15s | %-12d\n", p.ID, p.Nama, p.Kategori, p.Jumlah)
	}
	fmt.Println("=========================================================")
}

// Ubah Data Pengeluaran (UPDATE)
func UbahPengeluaran() {
    if totalData == 0 {
        fmt.Println(" Data kosong, tidak ada yang bisa diubah.")
        return
    }

    var idTarget int
    fmt.Print("Masukkan ID Pengeluaran yang ingin diubah: ")
    fmt.Scan(&idTarget)

    indexFound := -1
    for i := 0; i < totalData; i++ {
        if daftarPengeluaran[i].ID == idTarget {
            indexFound = i
            break
        }
    }

    if indexFound == -1 {
        fmt.Println(" ID tidak ditemukan!")
        return
    }

    fmt.Printf("Data Lama: %s (%s) - Rp%d\n", daftarPengeluaran[indexFound].Nama, daftarPengeluaran[indexFound].Kategori, daftarPengeluaran[indexFound].Jumlah)
    daftarPengeluaran[indexFound].Nama = readLine("Masukkan Nama Baru: ")
    daftarPengeluaran[indexFound].Kategori = readLine("Masukkan Kategori Baru: ")
    fmt.Print("Masukkan Jumlah Baru (Rp): ")
    fmt.Scan(&daftarPengeluaran[indexFound].Jumlah)
    
    fmt.Println("✓ Data berhasil diperbarui!")
}

// Hapus Data Pengeluaran (DELETE)
func HapusPengeluaran() {
	if totalData == 0 {
		fmt.Println(" Data kosong, tidak ada yang bisa dihapus.")
		return
	}

	var idTarget int
	fmt.Print("Masukkan ID Pengeluaran yang ingin dihapus: ")
	fmt.Scan(&idTarget)
    
	indexFound := -1
	for i := 0; i < totalData; i++ {
		if daftarPengeluaran[i].ID == idTarget {
			indexFound = i
			break
		}
	}

	if indexFound == -1 {
		fmt.Println(" ID tidak ditemukan!")
		return
	}

	// SHIFTING ARRAY
	for i := indexFound; i < totalData-1; i++ {
		daftarPengeluaran[i] = daftarPengeluaran[i+1]
	}
	totalData--
	fmt.Println("✓ Data berhasil dihapus dan posisi array dirapikan!")
}
