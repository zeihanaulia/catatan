# Retry with Backoff Mechanism

Pada sistem terdistribusi kegagalan komunikasi antar service adalah hal yang tidak bisa dihindarkan. Tidak ada yang bisa menjamin bahwa servis yang kita panggil akan 100% tersedia. Pasti akan ada kalannya dimana terjadi kegagalan sementara, entah karena database terlalu lama meresponse atau ketika sedang ada pembaharuan servis. Lalu bagaimana cara menghadapi masalah ini. Kalau mengingat kembali pendapat [Martin Fowler tentang karakteristik microservice](https://martinfowler.com/articles/microservices.html) ada karakteristik Design for Failure atau kita mendesain sistem untuk menghadapi kegagalan. Untuk tau mekanismenya apa saja bisa lihat di blog risingstack tentang[designing-microservices-architecture-for-failure](https://blog.risingstack.com/designing-microservices-architecture-for-failure/).

Ditulisan kali ini, pembahasannya spesifik ke Retry logic diblog itu juga dijelaskan retries dengan jumlah yang sangat besar juga bisa memperparah kondisi dimana akan terus membebani server. Agar tidak terlalu parah ditambahkan lah algoritma Backoff. Jadi,

>Mekanisme Retry adalah mekanisme yang **mengulang proses** jika terjadi kegagalan pada request sebelumnya.
>
>Mekanisme Backoff adalah mekanisme atau algoritma untuk **menentukan waktu tunggu** dan **membatalkan** saat mencapai hitungan yang ditetapkan.

Ada beberapa algoritma backoff, Apa saja? dikutip dari [retry-strategies-for-transient-failures](https://dev.to/ayushsharma/retry-strategies-for-transient-failures-4ci6)

- No Backoff          = Tidak ada delay pada setiap usaha retry-nya
- Constant Backoff    = Delaynya sudah tetap, misal diset setiap 5 detik akan mengulang
- Linear Backoff      = Delay yang ditambahkan dari setiap iterasi usaha retry
- Fibonacci Backoff   = Delay yang penambahannya berdasarkan fibonaci  0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
- Quadratic Backoff   = Delay yang penambahannya berdasarkan Quadratic
- Exponential Backoff = Delay yang penambahannya berdasarkan exponensial
- Polynomial Backoff  = Delay yang penambahannya berdasarkan polinomial 0, 1, 3, 6, ...

Kebutuhan dari mekanisme retry seperti ini.

- Identifikasi kalau kesalahan itu adalah kesalahan sementara.
  - Biasanya di http code 500, karena jika di 400 itu bisa saja kesalahan karena input dan memang clientnya yang harus diperbaiki
- Definisikan maksimum retry. Misal 3 kali coba ulang setelah itu stop.
- Setiap kali retry naikan hitungan usaha retry
- Jika sukses maka ya sudah lanjutkan ke client.
- Jika masih gagal, lanjutkan hingga hitungan retry tercapai.
- Jika sampai hitungan gagal, itu sama saja service memang tersedia. 
  - Biasanya kita butuh satu mekanisme Circuit Breaker, untuk menutup semua request

Tapi, Mengulang terus menerus bukanlah taktik yang efektif. Kenapa? coba kalo itu dilakukan dengan secara bersamaan ratusan atau ribuan request **ke service yang sedang overload / sibuk**. Nah ini akan membuat masalah baru thundering-herd problem, Client akan terus merequest tanpa batas ke service yang memang sibuk.
Solusi untuk mencegah masalah itu adalah dengan [Exponential Backoff](https://en.wikipedia.org/wiki/Exponential_backoff) dimana kita menambahkan periode waktu tunggu untuk melakukan retry.

## Exponential Back-off

Example bisa dicoba disini, menggunakan Library [github.com/cenkalti/backoff](github.com/cenkalti/backoff)

>Mekanisme Exponential Backoff adalah pengembangan dari mekanisme Backoff hanya saja penentuan waktunya dikalikan secara eksponensial, misalnya melakukan 10 kali retry maka akan selesai pada 43 detik.

```bash
2019/08/29 03:01:22 listening at :8091
2019/08/29 03:01:26 Retry took 71.576µs
2019/08/29 03:01:26 Retry took 564.754444ms
2019/08/29 03:01:28 Retry took 1.649490488s
2019/08/29 03:01:29 Retry took 2.968337811s
2019/08/29 03:01:30 Retry took 4.555636113s
2019/08/29 03:01:33 Retry took 6.899606963s
2019/08/29 03:01:37 Retry took 11.41400335s
2019/08/29 03:01:41 Retry took 14.640195828s
2019/08/29 03:01:46 Retry took 20.255575221s
2019/08/29 03:01:54 Retry took 27.911903345s
2019/08/29 03:02:09 Retry took 43.312680191s
2019/08/29 03:02:09 End of Retry Get http://localhost:8090: dial tcp [::1]:8090: connect: connection refused
```

## Jitter Back-off

Example bisa dicoba disini, menggunakan Library [https://github.com/Rican7/retry](https://github.com/Rican7/retry)

>Mekanisme Jitter Back-off adalah 

```bash

// Exponential Without Jitter

2019/08/29 03:34:13 listening at :8091
2019/08/29 03:34:18 Retry took 1.335µs
2019/08/29 03:34:19 Retry took 520.141067ms
2019/08/29 03:34:20 Retry took 1.522019494s
2019/08/29 03:34:21 Retry took 3.024834141s
2019/08/29 03:34:24 Retry took 5.529506299s
2019/08/29 03:34:27 Retry took 9.036588219s
2019/08/29 03:34:33 Retry took 14.547195571s
2019/08/29 03:34:41 Retry took 23.054750308s
2019/08/29 03:34:54 Retry took 35.562393486s
2019/08/29 03:35:13 Retry took 54.569926995s
2019/08/29 03:35:41 Retry took 1m23.077054828s
2019/08/29 03:35:41 End of Retry Get http://localhost:8090: dial tcp [::1]:8090: connect: connection refused
```

```bash

// Exponential With Jitter

2019/08/29 03:39:57 listening at :8091
2019/08/29 03:40:02 Retry took 18.856µs
2019/08/29 03:40:02 Retry took 555.989332ms
2019/08/29 03:40:04 Retry took 2.021498698s
2019/08/29 03:40:06 Retry took 3.92785303s
2019/08/29 03:40:07 Retry took 5.522857266s
2019/08/29 03:40:11 Retry took 8.893533715s
2019/08/29 03:40:16 Retry took 14.462596471s
2019/08/29 03:40:21 Retry took 19.325151204s
2019/08/29 03:40:33 Retry took 31.455846444s
2019/08/29 03:41:01 Retry took 59.676399012s
2019/08/29 03:41:21 Retry took 1m19.44061763s
2019/08/29 03:41:21 End of Retry Get http://localhost:8090: dial tcp [::1]:8090: connect: connection refused
```

### Menggunakan library backoff

Sebelum mengimplementasi dilambda coba pelajari dulu apa itu library backoff, coba dengan membuat golang file untuk menginvole ke api world

Referensi:

- [designing-microservices-architecture-for-failure](https://blog.risingstack.com/designing-microservices-architecture-for-failure/)
- [understanding-retry-pattern-with-exponential-back](https://dzone.com/articles/understanding-retry-pattern-with-exponential-back)
- [retries-timeouts-backoff](https://namc.in/2019-04-15-retries-timeouts-backoff)
- [retry-strategies-for-transient-failures](https://dev.to/ayushsharma/retry-strategies-for-transient-failures-4ci6)
