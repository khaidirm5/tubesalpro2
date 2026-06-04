package main

import "fmt"

func main() {
	pilihan := -1
	for pilihan != 0 {
		fmt.Println("\n=== APLIKASI BUDGET TRAVELING ===")
		
		// --- MENU BAGIANMU (CRUD & SORTING) ---
		fmt.Println("1. Atur / Ubah Budget Awal")
		fmt.Println("2. Tambah Pengeluaran")
		fmt.Println("3. Tampilkan Semua Pengeluaran")
		fmt.Println("4. Ubah Pengeluaran")
		fmt.Println("5. Hapus Pengeluaran")
		fmt.Println("6. Urutkan Berdasarkan Jumlah (Selection Sort)")
		fmt.Println("7. Urutkan Berdasarkan Kategori (Insertion Sort)")
		
		// --- MENU BAGIAN kahidir (DIBUAT NANTI) ---
		fmt.Println("8. [Belum Aktif] Cari Kategori (Sequential Search)")
		fmt.Println("9. [Belum Aktif] Cari Kategori (Binary Search)")
		fmt.Println("10. [Belum Aktif] Lihat Laporan Keuangan & Saran")
		
		fmt.Println("0. Keluar Aplikasi")
		fmt.Print("Pilih menu (0-10): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		// Menghubungkan ke fungsi di file kamu (crud.go & sorting.go)
		case 1:
			SetBudget() 
		case 2:
			TambahPengeluaran() 
		case 3:
			TampilSemua() 
		case 4:
			UbahPengeluaran() 
		case 5:
			HapusPengeluaran() 
		case 6:
			SortByJumlahSelection() 
		case 7:
			SortByKategoriInsertion() 
			
		case 8:
            SequentialSearchKategori()

        case 9:
            BinarySearchKategori()

        case 10:
            LaporanBudget()
			
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
