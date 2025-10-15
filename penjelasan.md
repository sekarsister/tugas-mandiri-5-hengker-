TUGAS 1 MENENTUKAN BILANGAN POLINDROM

1. Pendahuluan Program
Program ini merupakan implementasi algoritma untuk mendeteksi apakah suatu bilangan bulat merupakan bilangan polindrom atau bukan. Bilangan polindrom didefinisikan sebagai bilangan yang nilainya sama ketika dibaca dari depan maupun belakang.

2. Struktur Input dan Output
Program menerima satu input berupa bilangan bulat dari pengguna melalui fungsi baca input. Hasil deteksi akan ditampilkan dalam bentuk teks "polindrom" atau "bukan polindrom" sebagai output tunggal.

3. Alur Logika Utama
Program menggunakan pendekatan kondisional kompleks yang terbagi menjadi dua skenario utama berdasarkan tanda bilangan:

3.1. Penanganan Bilangan Negatif

    Bilangan negatif terlebih dahulu dikonversi menjadi nilai positif

    Nilai positif tersebut kemudian dibalik urutan digitnya menggunakan algoritma modulus

    Hasil pembalikan dibandingkan dengan nilai positif awal

    Jika sama, bilangan negatif tersebut diklasifikasikan sebagai polindrom

3.2. Penanganan Bilangan Non-Negatif

    Bilangan dengan satu digit (0-9) otomatis diklasifikasikan sebagai polindrom

    Untuk bilangan multi-digit, dilakukan perhitungan jumlah digit melalui iterasi pembagian

    Dilakukan analisis simetri digit dengan membandingkan digit dari kedua ujung menuju tengah

    Program hanya memeriksa hingga setengah jumlah digit untuk optimasi

4. Mekanisme Pembalikan Bilangan
Algoritma pembalikan bilangan menggunakan operasi matematika dasar:

    Digit terakhir diekstrak menggunakan operator modulus 10

    Digit tersebut ditambahkan ke hasil pembalikan dengan menggeser nilai sebelumnya

    Bilangan asli dikurangi digit terakhir melalui pembagian integer 10

    Proses diulang hingga semua digit terproses

5. Analisis Simetri Digit
Untuk bilangan multi-digit, program menerapkan:

    Perhitungan posisi digit dari kiri dan kanan

    Ekstraksi digit spesifik menggunakan fungsi pangkat dan modulus

    Perbandingan berpasangan digit kiri dan kanan

    Penghentian dini jika ditemukan ketidakcocokan

6. Optimasi dan Efisiensi
Program mengimplementasikan beberapa teknik optimasi:

    Early termination ketika ditemukan ketidaksesuaian digit

    Pembatasan iterasi hanya hingga setengah jumlah digit

    Penanganan kasus khusus tanpa melalui proses komputasi penuh

    Penggunaan operasi matematika murni tanpa konversi tipe data

7. Kompleksitas Algoritma

    Kompleksitas waktu: O(n) untuk bilangan dengan n digit

    Kompleksitas ruang: O(1) dengan penggunaan memori konstan

    Efisiensi tercapai melalui pembatasan iterasi hingga n/2

8. Keunggulan Pendekatan
Pendekatan yang digunakan memiliki keunggulan:

    Deteksi akurat untuk semua bilangan bulat

    Efisiensi komputasi yang tinggi

    Penggunaan resource memori minimal

    Kemampuan menangani bilangan besar tanpa overflow


PROGRAM FIBONACI

    Input: Membaca integer n dari user

    Validasi: Jika n < 0, program berhenti

    Output Jumlah Elemen: Cetak n+1 (total bilangan yang akan dicetak)

    Base Cases:

        Jika n >= 0, cetak 0 (Fibonacci ke-0)

        Jika n >= 1, cetak 1 (Fibonacci ke-1)

    Perhitungan Fibonacci:

        Inisialisasi a = 0, b = 1

        Untuk i dari 2 hingga n:

            next = a + b

            Cetak next

            Update: a, b = b, next
