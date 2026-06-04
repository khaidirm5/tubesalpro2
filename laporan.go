package main

import "fmt"

func LaporanBudget() {
    if totalData == 0 {
        fmt.Println("Belum ada data pengeluaran.")
        return
    }

    var totalPengeluaran int
    var transportasi int
    var akomodasi int
    var makanan int
    var hiburan int

    for i := 0; i < totalData; i++ {

        totalPengeluaran += daftarPengeluaran[i].Jumlah

        switch daftarPengeluaran[i].Kategori {

        case "Transportasi", "transportasi":
            transportasi += daftarPengeluaran[i].Jumlah

        case "Akomodasi", "akomodasi":
            akomodasi += daftarPengeluaran[i].Jumlah

        case "Makanan", "makanan":
            makanan += daftarPengeluaran[i].Jumlah

        case "Hiburan", "hiburan":
            hiburan += daftarPengeluaran[i].Jumlah
        }
    }

    sisaBudget := budgetAwal - totalPengeluaran

    fmt.Println()
    fmt.Println("========================================")
    fmt.Println("          LAPORAN BUDGET TRAVEL")
    fmt.Println("========================================")

    fmt.Printf("Budget Awal       : Rp %d\n", budgetAwal)
    fmt.Printf("Total Pengeluaran : Rp %d\n", totalPengeluaran)
    fmt.Printf("Sisa Budget       : Rp %d\n", sisaBudget)

    fmt.Println("\nRINCIAN PER KATEGORI")
    fmt.Println("----------------------------------------")
    fmt.Printf("Transportasi : Rp %d\n", transportasi)
    fmt.Printf("Akomodasi    : Rp %d\n", akomodasi)
    fmt.Printf("Makanan      : Rp %d\n", makanan)
    fmt.Printf("Hiburan      : Rp %d\n", hiburan)

    fmt.Println("----------------------------------------")

    if sisaBudget < 0 {
        fmt.Println("Status : Budget terlampaui.")
    } else {
        fmt.Println("Status : Budget masih aman.")
    }

    fmt.Println("========================================")
}
