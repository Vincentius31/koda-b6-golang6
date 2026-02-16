# 🍜 Aplikasi Antrian Pemesanan Makanan

Repositori ini berisi implementasi sistem antrian pemesanan makanan yang dikonversi dari JavaScript 3 Repositori ke **Golang** menggunakan Goroutines, Channels, dan WaitGroup.

---

## 📂 Struktur Kode
Aplikasi ini menangani data berikut:
| Nama Variabel | Tipe Data | Deskripsi |
| :--- | :--- | :--- |
| `Name` | `string` | Nama pemesan makanan. |
| `Wait` | `int` | Waktu tunggu dalam detik. |

---

## 🛠️ Penjelasan Kode

### 1. Channels (`chan`)
Channel digunakan sebagai wadah untuk mengirim data antara Goroutine.
* **Deklarasi**: `queueChan := make(chan Person)` membuat saluran khusus untuk mengirim objek `Person`.
* **Pengiriman**: `queueChan <- p` memasukkan data pesanan ke dalam antrian.
* **Penerimaan**: `for p := range queueChan` di dalam Goroutine akan terus menunggu dan mengambil data dari channel selama channel tersebut belum ditutup.

### 2. WaitGroup (`sync.WaitGroup`)
WaitGroup berfungsi sebagai penghitung (counter) untuk memastikan aplikasi tidak berhenti sebelum semua proses asinkron selesai.
* **`wg.Add(1)`**: Dipanggil setiap kali ada pesanan baru yang akan diproses (menambah antrian).
* **`wg.Done()`**: Dipanggil menggunakan `defer` di dalam fungsi proses untuk memberitahu bahwa satu pesanan telah selesai.
* **`wg.Wait()`**: Menahan *main program* agar tetap berjalan sampai seluruh counter kembali ke angka nol (semua pesanan selesai).

---

## 💻 Cara Menjalankan
1.  Buka terminal di VS Code.
2.  Pastikan kode Go sudah ada di file `main.go`.
3.  Jalankan perintah:
    ```
    go run main.go
    ```

---

## 📝 Logika Program
Program akan memproses daftar pesanan secara berurutan:
1.  **Validasi**: Jika nama tersedia, program mencetak status "Menunggu antrian...". Jika kosong, mencetak peringatan.
2.  **Simulasi Waktu**: Menggunakan `time.Sleep` berdasarkan atribut `Wait` dari tiap person.
3.  **Selesai**: Setelah semua pesanan diproses dan channel dikosongkan, pesan "Selesai" akan muncul sebagai penutup.

