# Budget Traveling Management System

Aplikasi manajemen budget perjalanan berbasis terminal yang dikembangkan menggunakan bahasa Go (Golang). Program ini membantu pengguna mengatur anggaran perjalanan, mencatat pengeluaran, melakukan pencarian dan pengurutan data, serta menghasilkan laporan keuangan sederhana untuk memantau kondisi budget selama perjalanan.

## Features

### Budget Management
- Menentukan dan mengubah budget awal perjalanan.
- Menghitung sisa budget secara otomatis.

### Expense Management (CRUD)
- Menambahkan data pengeluaran.
- Menampilkan seluruh data pengeluaran.
- Mengubah data pengeluaran berdasarkan ID.
- Menghapus data pengeluaran berdasarkan ID.

### Sorting
- Selection Sort berdasarkan jumlah pengeluaran.
- Insertion Sort berdasarkan kategori pengeluaran.

### Searching
- Sequential Search berdasarkan kategori.
- Binary Search berdasarkan kategori.

### Financial Report
- Menampilkan total pengeluaran.
- Menampilkan sisa budget.
- Menampilkan rincian pengeluaran per kategori.
- Memberikan status kondisi budget.

---

## Data Structure

Program menggunakan struktur data statis berupa array dan record (struct).

```go
type Pengeluaran struct {
	ID       int
	Nama     string
	Kategori string
	Jumlah   int
}
```

Kategori yang digunakan:

- Transportasi
- Akomodasi
- Makanan
- Hiburan

---

## Algorithms Implemented

### Selection Sort
Digunakan untuk mengurutkan data berdasarkan jumlah pengeluaran.

Complexity:
- Best Case: O(n²)
- Average Case: O(n²)
- Worst Case: O(n²)

### Insertion Sort
Digunakan untuk mengurutkan data berdasarkan kategori.

Complexity:
- Best Case: O(n)
- Average Case: O(n²)
- Worst Case: O(n²)

### Sequential Search
Digunakan untuk mencari kategori tanpa memerlukan data terurut.

Complexity:
- O(n)

### Binary Search
Digunakan untuk mencari kategori setelah data diurutkan berdasarkan kategori.

Complexity:
- O(log n)

---

## Project Structure

```text
budget-traveling/
│
├── main.go
├── config.go
├── crud.go
├── sorting.go
├── searching.go
├── laporan.go
├── go.mod
└── README.md
```

### File Description

| File | Description |
|--------|--------|
| main.go | Menu utama dan kontrol program |
| config.go | Struktur data dan variabel global |
| crud.go | Fitur Create, Read, Update, Delete |
| sorting.go | Implementasi Selection Sort dan Insertion Sort |
| searching.go | Implementasi Sequential Search dan Binary Search |
| laporan.go | Pembuatan laporan keuangan dan analisis budget |

---

## How to Run

### Clone Repository

```bash
git clone https://github.com/username/budget-traveling.git
cd budget-traveling
```

### Run Program

```bash
go run .
```

atau

```bash
go run *.go
```

### Build Executable

```bash
go build .
```

Jalankan executable:

```bash
./budget-traveling
```

---

## Example Menu

```text
=== APLIKASI BUDGET TRAVELING ===

1. Atur / Ubah Budget Awal
2. Tambah Pengeluaran
3. Tampilkan Semua Pengeluaran
4. Ubah Pengeluaran
5. Hapus Pengeluaran
6. Urutkan Berdasarkan Jumlah (Selection Sort)
7. Urutkan Berdasarkan Kategori (Insertion Sort)
8. Cari Kategori (Sequential Search)
9. Cari Kategori (Binary Search)
10. Lihat Laporan Keuangan & Saran
0. Keluar Aplikasi
```

---

## Example Report

```text
========================================
          LAPORAN BUDGET TRAVEL
========================================
Budget Awal       : Rp500000
Total Pengeluaran : Rp445000
Sisa Budget       : Rp55000

RINCIAN PER KATEGORI
----------------------------------------
Transportasi : Rp25000
Akomodasi    : Rp100000
Makanan      : Rp70000
Hiburan      : Rp250000
----------------------------------------
Status : Budget masih aman.
========================================
```

---

## Technologies

- Go (Golang)
- Terminal-Based Application
- Array
- Struct
- Selection Sort
- Insertion Sort
- Sequential Search
- Binary Search

---

## Learning Outcomes

Proyek ini dibuat untuk mengimplementasikan konsep:

- Algoritma dan Pemrograman
- Struktur Data Sederhana
- Array dan Struct
- Modular Programming
- Searching Algorithms
- Sorting Algorithms
- CRUD Operations
- Budget Tracking System

---

## Authors

- Abdurrahman
- Khaidir Maulana

Telkom University
