# SOLID Principle

- The **S**ingle Responsibility (SRP) Principle
- The **O**pen/Close Principle
- The **L**iskov Subtitution Principle
- The **I**nterface Segregation Principle
- The **D**ependency Inversion Principle

Software yang baik dimulai dari "Clean Code". Dengan kata lain batu bata akan menjadi tembok, dan tembok akan menjadi bangunan. kalo susunan batu batanya aja udah kacau udah pasti bangunannya bakal kacau juga. Dengan ini makannya diperlukan prinsip - prinsip agar kode yang deibuat menjadi baik. Prinsip ini dikenal dengan nama SOLID.

Tujuan dari prinsip-prinsip diatas adalah sebagai berikut:

- Toleransi terhadap perubahan
- Mudah untuk dimengerti
- Merukapan dasar komponen yang bisa digunakan di banyak software

## SRP: The Single Responsibility Principle

Konsekuensi dari hukum Conway: Struktur terbaik untuk software sangat dipengaruhi oleh struktur sosial organisasi yang menggunakan sehingga setiap modul software memiliki satu dan hanya satu alasan untuk berubah.

> Do One Thing and Do It Well

> A class should have one, and only one, reason to change. - Uncle Bob (Robert C. Martin)

### Identifikasi pelanggaran SRP

- Class yang besar dan banyak statement bercabang
- Nama class biasanya bernama Manager, Processor dll
- Susah memberikan nama class secara sadar
- Unit test yang panjang

### Create SRP class

- Hanya ada satu alasan untuk mengubah class
- Class kecil
- High Cohesion and lose copeling

### Contoh:

- Controller seharusnya tidak ada query kedatabase, hanya sebagai pengkontrol antara request dan data
- Class Payment, tidak seharusnya berisi fungsi mengirim notifikasi. Juga seharusnya tidak berisi metode pembayaran. class Payment haruslah hanya berfungsi sebagai penghitung pembayaran saja.
- Class User tidak seharusnya ada validasi email, sebaiknya hal hal yang berkaian dengan email dibuatkan class baru yaitu Email.
- Class Controller tidak seharusnya melakukan perintah print statement, Class controller seharusnya hanya mengontrol, menerima request , mengambil data berdasarkan request lalu memilih view untuk ditampilkan dimana.

### Solusi 

Dapat menerapkan The Facade pattern

## OCP: The Open-Closed Principle

Betrand Meyer, membuat prinsip ini terkenal pada tahun 1980an, intinya adalah software harus mudah untuk berubah perilakunya, tetapi hanya boleh dengan menambahkan kode, bukan mengubah kode yang sudah ada.

Ini aturannya agak aneh, mengizinkan tapi dibatasi hanya boleh dengan menambah. kode boleh beruab perilakukanya tapi hanya boleh dengan cara menambahkan bukan mengubah kode. 

Uncle Bob berkata 

> Open for extension, Closed for modification. - Uncle Bob (Robert C. Martin)

### Open for extension

Artinya perilaku module bisa berubah tapi dengan cara di extend.

### Close for modification

Exending perilaku dari module bukan berarti mengubah kode sumber dari modul itu. Kode sumber tidak boleh diubah ubah.

## LSP: The Liskov Substitution Principle

Definisi subtipe barbar liskov yang terkenal dari tahun 1988, intinya prinsip ini mengatakan bahwa membangun software dari bagian yang dapat dipertukarkan. bagian tersebut harus mematuhi kontrak yang memungkinkan bagian tersebut diganti satu sama lain,

## ISP: The Interface Segregation Principle

Prisip ini menyarankan agar perancang software untuk menghindari ketergantungan pada hal hal yang tidak digunakan.

## DIP: The Dependency Inversion Principle

Kode dilevel tertinggi tidak boleh memiliki ketergantungan dengan kode yang berlevel rendah, Sebaliknya rincian harus bergantung ke master.