# Protocol Buffers 3

Didalam dunia cloud native, [protocol buffers 3](https://developers.google.com/protocol-buffers) atau biasa disingkat protobuf ini berperan sebagai [serialisasi](https://id.wikipedia.org/wiki/Serialisasi) data struktur yang digunakan oleh framework [grpc](https://grpc.io/). Intinya protobuf itu sama seperti XML atau JSON hanya saja ukurannya lebih kecil, lebih cepat dan lebih sederhana.

Keuntungan menggunakan protobuf:

- Data diketik sepenuhnya
- Data dikompresi secara otomatis (mengurangi penggunaan CPU)
- Menggunakan schema untuk menghasilkan kode dan untuk membaca data
- Dokumentasi dapat dimasukan kedalam skema
- Data dapat berkembang kapanpun dan aman, perubahannya tidak merusak.
- 3-10x lebih kecil, 20-100x lebih cepat dari XML
- Kode dapat dihasilkan secara otomatis!

Kekurangan menggunakan protobuf:

- Belum mensuport semua bahasa pemrograman, bisa dilihat [disini](https://developers.google.com/protocol-buffers/docs/reference/other)
- Data yang sudah diserialisasi tidak dapat dibuka untuk dibaca dengan text editor

Bagaimana protobuf digunakan?

Untuk memahami cara berkerja, kita harus paham terlebih dulu kenapa protobuf diciptakan?. Ide dasarnya adalah untuk membagi data lintas bahasa. Dimulai dengan membuat file .proto yang bisa dibaca oleh manusia, Lalu secara otomatis file .proto akan menghasilkan kode secara otomatis ke bahasa pemrogaman yang ditargetkan. Dengan bahasa pemrograman yang kita targetkan kita dapat membuat object untuk mengserialisasikan data dari file .proto dan dengan kode tersebut kita dapat meinterpreted agar dapat dibaca.

Beberapa database sudah ada yang support protobuf seperti [mysql](https://github.com/google/mysql-protobuf). Juga beberapa framework [RPC](https://en.wikipedia.org/wiki/Remote_procedure_call) seperti [grpc](https://grpc.io/) dan project besar seperti [etcd](https://etcd.io/) juga menggunakan protobuf.

## Mencoba protobuf

### Dasar

syntax protobuf seperti ini

```proto
syntax = "proto3";

message MyMessage {
    int32 id = 1;
    string first_name = 2;
    bool is_valid = 3;
}
```

- syntax             = versi dari protobuf yang digunakan.
- message MyMessage  = mendefinisikan pesan, `message` adalah keyword dari protobof, sedangkan `MyMessage` adalah nama pesan.
- `id, first_name, is_valid` = ini disebut dengan nama field 
- setiap field memiliki type seperti `int32, string, bool`
- penomoran `=3, = 2, =1` ini adalah tags berupa number

### Scalar Type

1. Number

Number terdiri dari

## Referensi

- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [Complete Introduction to Protocol Buffers 3 - Stéphane Maarek](https://learning.oreilly.com/videos/complete-introduction-to/9781789349344)
