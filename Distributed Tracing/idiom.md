# Kosakata yang baru tau

## Metadata/Context Propagation

Di dalam distributed tracing, kita dapat meindentifikasi penggunaan dari hardware yang kita akan gunakan, seperti CPU cores atau disk space. Sehingga kita dapat memprediksi biaya dari operational aplikasi kita.

Sebagai contoh, Misalnya kita memiliki dua lini bisnis dan keduanya memerlukan 1000 cores CPU. Satu lini bisnis kita asumsikan akan bertumbuh sebanyak 10% dan satu lagi akan bertumbuh 100% ditahun depan. Sampai sini kita tidak bisa membedakan kedua lini bisnis karena keduanya digabung. Kalau tenyata Lini bisnis 1 menggunakan 90% hardware artinya dia menggunakan 900 to 990 CPU cores dan lini bisnis 2 sebanyak 10% artinya dia akan menambah dari 100 ke 200 CPU cores. Maka penambahan seharusnya hanya 190 cores. Tapi kalau ternyata dibagi 50/50 maka perhitungannya berubah menjadi 500 *1.1 + 500* 2.0=1550 cores.

Lalu bagaimana caranya mengidentifikasi hal seperti dengan cara mencatat `Resource usage attribution` atau implementasi sistem yang support multi-tenancy. Dan akan sangat bagus kalau kita bisa menghitung berapa CPU time yang digunakan per-customer. Logika sederhanyanya, kita bisa mencatat semua request dengan menambahkan customer id pada api kita, tapi itu kurang effisien, didalam distributed tracing dengan opentracing ada cara lebih baik yaitu dengan menerapkan Distributed Metadata/Context Propagation. Ada istilah `Baggage` yang dapat menyimpan informasi didalamnya dan akan dapat dipanggil di service service berikutnya.
