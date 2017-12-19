# Pengenalan laravel

## Versioning Scheme

Paradigma relis versi terbaru laravel mengenal istilah major dan minor. Rilis versi major akan selalu keluar setiap 6 bulan sekali (February dan Agustus) sedangkan rilis versi minor keluar setiap 2 minggu sekali. Minor rilis tidak akan merusak kode, untuk major rilis bisa merusak.

Solusinya selalu gunakan version constaint 5.5.* (sesuai versi lts). Tapi laravel selalu berusaha untuk membuat pembaharuan lebih cepat.

## Kenapa sih laravel gak menggunakan SemVer (Semantic Versioning)

Banyak jg komponent laravel yang menerapkan Semver, Tapi laravel sendiri tidak, alasannya karena Semver itu "reduksionis" untuk menentukan apakah dua kode kompetibel. Kalaupun laravel menggunakan Semver , kita juga musti install package yang telah di upgrade dan menjalankan test suit otomatis untuk mengetahui kalau ada yang tidak sesuai dengan kode kita.

## Kebijakan dukungan

untuk versi LTS 5.5, bug fix akan terus dilayani selama 2 tahun dan untuk security akan dilayani selama 3 tahun. untuk general rilis akan dilayani selama 6 bulan dan security fix selama satu tahun.

## Laravel 5.5 (LTS)

Dari versi 5.4 ke 5.5 laravel menambahkan beberapa paket seperti:

- Auto-Detection
- API resource / transformation
- Auto-registration of console commands
- Queued job chaining
- Queue job rate limiting
- Time based job attempt
- Renderable mailables
- Renderable and reportable exception
- More consisten exception handling
- Database testing improvement
- Simpler custom validation rules
- React front-end preset
- Route::view and Route::redirect method
- "locks" for memcached and redis cache drivers
- on-demand notification
- headless chrome support in Dusk
- Convenient blade shortcut
- Improve trusted proxy support
- more

rilisnya laravel 5.5 bertepatan dengan rilisnya laravel-horizon,  **a beautiful new queue dashboard and configuration system for your Redis based Laravel queues**. 

## Package Discovery / Auto-detection

Cukup menggunakan composer package yang kita inginkan langsung masuk kedalam aplikasi laravel kita

## API resource / transformation

Ketika membuat API kita membutukan transformasi layer dari model ke JSON. Fitur ini mempermudah kita, tanpa harus membuat lagi transformasi secara manual

## Auto-registration of console commands

Laravel sekarang sudah bisa scaning command file menggunakan new load

## React front-end preset / New Frontend Presets

Laravel dan react sebagai frontend , sekarang sudah bisa

## Queued Job Chaining

Laravel juga bisa dan memungkinkan untuk membuat job secara berurutan menggunakan method withChain

## Queued Job Rate Limiting

Laravel juga bisa membatasi job misal, kita ingin membuat job yang berjalan 10 kali setiap 60 detik

## Time Based Job Attempts

Laravel juga bisa membatasi berapa banyak job yang dieksekusi kemudian digagalkan

## Validation Rule Objects

Laravel memberikan kemudahan dalam pembuatan rule validasi menggunakan artisan

cukup menjalankan perintah:

```bash
php artisan make:rule ValidName
```

Maka laravel akan membuat file bernama ValidName

## Trusted Proxy Integration

Laravel memasukan Trusted Proxies package by Chris Fidao ke dalam standart laravel

## On-Demand Notifications

Laravel membuat kita bisa menggunakan notifikasi sesuai dengan kebutuhan

## Renderable Mailables

Laravel juga memungkinkan kita melihat email kita sebelum dikirimkan

## Renderable & Reportable Exception

Laravel juga memudahkan untuk melakukan kustomisasi exception

## Request Validation

Laravel juga bisa melakukan validasi request masuk dengan cepat

## Consistent Exception Handling

Laravel juga menawarkan konsistensi error handling dengan menggunakan App\Exceptions\Handler

## Cache Locks

Laravel sekarang memiliki method locak Cache

## Blade Improvements

Laravel juga melakukan perbaikan pada bladenya, sekarang blade memiliki Blade::if dimana sekarang bisa melakukan custom directve menggunakan closures

## New Routing Methods

Laravel juga memberikan method baru pada routing yaitu **redirect** dan **view**

## "Sticky" Database Connections

Laravel memberikan pilihan, kalau kita menginginkan koneksi yang sama ketika melakukan write and read

## Kesimpulan

Diatas adalah release note dari laravel
laravel menawarkan banyak sekali feature yang bisa digunakan dapat membuat produktifitas meningkat, disamping ketika release versi baru ada banyak kode yang akan berubah

komunitas laravel cukup besar sehingga dapat mengurangi banyak relearning framework, artinya banyak sekali orang yang menggunakan project menggunakan laravel

Dokumentasi dari laravel diwebsitenya juga sangat lengkap. cukup untuk memulai membuat project dari laravel tanpa harus kebingungan mencari trainer.

