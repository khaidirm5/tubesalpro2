package main

import "fmt"

func SequentialSearchKategori() {
    if totalData == 0 {
        fmt.Println("Data pengeluaran masih kosong.")
        return
    }

    var kategori string
    var ditemukan bool

    fmt.Print("Masukkan kategori yang dicari: ")
    fmt.Scan(&kategori)

    fmt.Println("\nHASIL PENCARIAN (Sequential Search)")
    fmt.Println("---------------------------------------------------------")
    fmt.Printf("%-4s %-20s %-15s %-12s\n",
        "ID", "Nama", "Kategori", "Jumlah")
    fmt.Println("---------------------------------------------------------")

    for i := 0; i < totalData; i++ {
        if daftarPengeluaran[i].Kategori == kategori {
            fmt.Printf("%-4d %-20s %-15s Rp %-10d\n",
                daftarPengeluaran[i].ID,
                daftarPengeluaran[i].Nama,
                daftarPengeluaran[i].Kategori,
                daftarPengeluaran[i].Jumlah)

            ditemukan = true
        }
    }

    if !ditemukan {
        fmt.Println("Data tidak ditemukan.")
    }

    fmt.Println("---------------------------------------------------------")
}

func BinarySearchKategori() {
    if totalData == 0 {
        fmt.Println("Data pengeluaran masih kosong.")
        return
    }

    SortByKategoriInsertion()

    var kategori string

    fmt.Print("Masukkan kategori yang dicari: ")
    fmt.Scan(&kategori)

    left := 0
    right := totalData - 1
    found := -1

    for left <= right {
        mid := (left + right) / 2

        if daftarPengeluaran[mid].Kategori == kategori {
            found = mid
            break
        } else if daftarPengeluaran[mid].Kategori < kategori {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    if found == -1 {
        fmt.Println("Data tidak ditemukan.")
        return
    }

    fmt.Println("\nHASIL PENCARIAN (Binary Search)")
    fmt.Println("---------------------------------------------------------")
    fmt.Printf("%-4s %-20s %-15s %-12s\n",
        "ID", "Nama", "Kategori", "Jumlah")
    fmt.Println("---------------------------------------------------------")

    i := found

    for i >= 0 && daftarPengeluaran[i].Kategori == kategori {
        i--
    }

    i++

    for i < totalData && daftarPengeluaran[i].Kategori == kategori {
        fmt.Printf("%-4d %-20s %-15s Rp %-10d\n",
            daftarPengeluaran[i].ID,
            daftarPengeluaran[i].Nama,
            daftarPengeluaran[i].Kategori,
            daftarPengeluaran[i].Jumlah)
        i++
    }

    fmt.Println("---------------------------------------------------------")
}
