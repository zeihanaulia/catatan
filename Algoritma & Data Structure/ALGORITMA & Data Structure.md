# Algoritma & Data Structure

## Kenapa harus belajar Data Structure

Pengetahuan data structure sangat penting untuk perancangan dan implementasi perangkat lunak. Tujuannya untuk menghasilkan perangkat lunak yang berkinerja dengan baik dan efisien, bukan hanya menghasilkan yang benar saja (yang penting jadi). Struktur data yang dirancang dengan baik dapat meningkatkan kecepatan suatu program dan mengurangi jumlah memori yang dibutuhkan untuk menyimpan data pada program yang di proses. Selain itu, belajar Data Structure berfungsi sebagai jembatan antara topik pemrogramman dasar dan topik yang lebih maju seperti sistem operasi, jaringan dan organisasi komputer.

## Apa itu data?

> Data adalah informasi

Setiap informasi (**data**) yang datang pasti akan ada sumbernya (**source**). Jadi, **Informasi** itu seperti nama, alamat, email dan lain-lain sedangkan **sumbernya** bisa jadi dari device yang dipasangkan ke sistem komputer seperti keyboard, mouse , modem , disk drive atau perangkat penyimpanan lainnya. kadang program yang kita tulis juga bisa menjadi source dari data itu sendiri, selama program tersebut memproses dan mengenerate data baru.

> Faktanya, Studi menunjukkan bahwa program menghabiskan 80% waktu eksekusi mereka untuk mencari melalui memori untuk menemukan data yang mereka proses.

## Data Structure adalah

The National Institute of Standards and Technology, mendefinisikan data structure sebagai berikut:
> A data structure is an organization of information, usually in memory, for better algorithm efficiency.

Loonycorn - Udemy Course From 0 to 1: Data Structures & Algorithms in Java:
> Data Structure are way to organize information in such way that we can do useful things with it.

Premis:

- Computing is all about performing operations on information
- The used of data structres is to help perform operations efficiently
- Data adalah informasi, dan penggunaan data structure untuk membantu melakukan operasi secara efisien.

Jadi data terstruktur adalah cara untuk mengatur informasi dengan sedemikian rupa sehingga, kita bisa melakukan hal-hal yang bermanfaat dengannya.

Standart untuk kebaikan -  William McAllister, Data Structures and Algorithms Using Java

> A good data structure is one that organizes the data in a way that facilitates the operations performed on the data, while keeping its total storage requirement at, or close to, the size of the data set.

Struktur data yang baik adalah data yang mengatur data dengan cara memudahkan operasi yang dilakukan pada data, sekaligus menjaga agar kebutuhan penyimpanan total atau mendekati ukuran dari data set.

## Ada 2 tipe Data Structure

### built-in data structure

Struktur data built-in adalah skema untuk menyimpan data yang merupakan bagian dari bahasa pemrograman.

### programmer defined data structure

Struktur data yang diprogram oleh pemrogram adalah skema untuk menyimpan data yang dipahami dan diterapkan oleh pemrogram program tertentu. Contohnya ada dibeberapa bahasa pemrograman yang sudah mengimplementasikan data structure didalam standart librarynya

- <http://php.net/manual/en/spl.datastructures.php> (PHP)
- <http://www.geeksforgeeks.org/the-c-standard-template-library-stl/> (C++)
- <https://golang.org/pkg/> (go)

## Dampak Data Structure terhadap performa

Ada 2 kriteria evaluasi untuk program yang ditulis:

### 1. The “penmanship” of the program is good

disini kita dapat mengevaluasi hal hal seperti, tidak adanya syntax error, penamaan variabel yang baik, indentation yang baik, pengaturan code yang baik (termasuk jika menggunakan subprogram) dan penggunaan bahasa pemrogramman yang baik (contoh, penggunaan pernyataan switch vs pernyataan if-else)

### 2. For all sets of valid inputs, the program produces the correct outputs

Untuk semua rangkaian masukan yang valid, program menghasilkan keluaran yang benar.

Selain 2 kriteria itu, masih ada 2 kriteria tambahan lagi

1. Efficient use of storage—both main memory and external storage.
1. Speed of execution.

Dalam dunia nyata yang memiliki data set yang besar 2 kriteria tambahan perlu juga diperhatikan. Karena akan berdampak pada program yang kita buat.

## Apa yang dapat dilakukan dengan Data Structure

- Make common operations fast
- Make difficult operations possible
- Occupy less space and still represent the complexity of information and it's interrelationalship in an intuitive way

## Data Structure VS Algorithm

- Data structure lend itself to efficient algorithms
- Data struture dibutuhkan untuk membuat algoritma dan efisien
- Data structure form the core in many algorithms, such as sorting, graph traversal, searching, etc
- Data structure adalah inti dari banyak algoritma
- Data structure and algorithm go hand in hand
- Data structure dan algoritma berjalan beriringan

## Tugas seperti apa yang membutuhkan Data Structure

- A **set** might be useful for extremely fast membership and containment queries
- Compiler use **Hash Tables** as look up tables for operations such as runtime method binding
- **Stacks** can be used for undo functionality in application as well as the back functionality in browser
- Index data structures such as a **suffix tree** or and **inverted index** as used in search engine indexing
- **Graphs** are used to represent relationship such as on social networking

> Data structures influence Algorithms as much as Algorithms influence the design of Data sctrutures

## Data Structure VS Data Types

### Abstract Data

are mathematical model of data types, where the data type is defined by how it is **used** i.e from the point of view of the user

abstrak data adalah model mathematical dari data type. dimana data type didefinisikan bagaimana itu berkerja dan biasanya ini dibuat dari sisi penggunanya.

These define the operations to be performed on data, and what the expected behavior of these operations are.
Menentukan operasi apa yang akan dilakukan kepada data dan perilaku apa yang diharapkan dari operasi ini.

### Data structures

Data structure are **concrete representations** of data from the point of view of an implementor.
Jika data type kita bisa tau apa saja yang akan dilakukan oleh suatu task, maka data structure akan membuat implementasinya.

This specifies the actual implementation of the structure in code to meet the expected behavior

An Abststact data type can be defined as:
> Class of objects whose logical behavior is defined by set of values and set of operations

Kelas objek yang perilaku logisnya ditentukan oleh seperangkat nilai dan rangkaian operasi.

It does not specify how the type will actualy exhibit that behavior,
that is where data structures come in

## Sebenarnya data structure itu baik untuk apa sih?

Seperti yang sudah kita pelajadi diatas, bahwa data structure itu bisa dibilang adalah susunan data didalam memori komputer. Data structure termasuk array, linked list, stacks, binary trees dan hash table. Sedangkan Algoritma, dia memanipulasi data dalam struktur ini dengan berbagai cara, seperti mencari item data tertentu dan menyortir data.

Lalu masalah macam apa yang bisa diselesaikan dalam dunia nyata?

1. Real world data storage
1. Programmer Tools
1. Modeling

Maksudnya?

### Real world data storage

Didalam dunia komputer contoh kasusnya seperti ini. misalnya, record personel menjelaskan data terkini dari setiap orang, record inventory yang dapat menunjukan lokasi atau stok barang dan atau record financial yang dapat memberikan rincian pengeluaran atau tagihan yang harus dibayar.

Contoh non komputer, misalnya tumpukan kertas yang berisi data-data. jika kertas tersebut berisi nama, alamat, nomo telepon maka kertas itu akan menjadi buku alamat. jika kertas tersebut berisi nama, lokasi dan harga rumah maka kertas itu menjadi inventory perumahan.

Semua itu adalah data yang bisa kita olah menjadi suatu informasi yang bernilai, dan sekarang pertanyaannya adalah

- Bagaimana kamu menyimpan data tersebut didalam komputer
- Apakah metode kamu bisa berkerja dengan ratusan data? ribuan atau bahkan jutaan data?
- Apakah metode anda memungkinkan untuk menambahkan data dan menghapus data lama?
- Apakah metode anda memungkinkan untuk melakukan pencarian data dengan cepat?
- Jika anda ingin mengurutkannya, bagaimana caranya?

Didalam dunia nyata pertannyaan-pertanyaan itu akan terus ditanyakan agar sistem yang kita buat menjadi efisien dan baik.

### Programmer Tools

Tidak selalu data structure digunakan untuk penyimpanan, biasanya data diakses secara langsun dengan programm yang dibuat. Ada data data yang tidak bisa diakses oleh user tetapi harus melalui program. Progammer menggunakan data structure untuk mengatasinya. Stacks, antrian dan priority antrian  paling sering digunakan.

## Overview Data struture

Berikut ini adalah kekuatan dan kelemahan dari masing-masing data structure. dengan melihat kelemahan dan kekurangan kita bisa dengan mudah melihat dan membedakan data structure.

### Array

- Keuntungan : Quick insertion, very fast access if index known
- Kelemahan  : Slow search, slow deletion, fixed size.

### Ordered Array

- Keuntungan : Quicker search than unsorted array.
- Kelemahan  : Slow insertion and deletion, fixed size.

### Stack

- Keuntungan : Provides last-in, first-out access.
- Kelemahan  : Slow access to other items.

### Queue

- Keuntungan : Provides first-in, first-out access
- Kelemahan  : Slow access to other items.

### Linked list

- Keuntungan : Quick insertion, quick deletion.
- Kelemahan  : Slow search.

### Binary tree

- Keuntungan : Quick search, insertion, deletion (if tree remains balanced).
- Kelemahan  : Deletion algorithm is complex.

### Red-black tree

- Keuntungan : Quick search, insertion, deletion. Tree always balanced.
- Kelemahan  : Complex.

### 2-3-4 tree

- Keuntungan : Quick search, insertion, deletion. Tree always balanced. Similar trees good for disk storage.
- Kelemahan  : Complex.

### Hash table

- Keuntungan : Very fast access if key known. Fast insertion.
- Kelemahan  : Slow deletion, access slow if key not known, inefficient memory usage.

### Heap

- Keuntungan : Fast insertion, deletion, access to largest item.
- Kelemahan  : Slow access to other items.

### Graph

- Keuntungan : Models real-world situations.
- Kelemahan  : Some algorithms are slow and complex.

## Overview algorithms 

Banyak dari algoritma yang secara sepesifik menuju data structure. Dari banyak data structure anda perlu tau caranya

- Insert a new data item
- Search for specific item
- Delete a specific item

Question:

1. In many data structures you can **insert** a single record, **search** it, and **delete** it.
1. Rearranging the contents of a data structure into a certain order is called **sort**.
1. In a database, a field is 

    a. a specific data item.

    b. **a specific object.**

    c. part of a record.

    d. part of an algorithm.
1. The field used when searching for a particular record is the **index**

1. In object-oriented programming, an object

    a. **is a class**.

    b. may contain data and methods.

    c. is a program.

    d. may contain classes.

1. A class
    a. is a blueprint for many objects.

    b. **represents a specific real-world object**.

    c. will hold specific values in its fields.

    d. specifies the type of a method.
1. In Java, a class specification:

    a. creates objects.

    b. **requires the keyword new.**

    c. creates references.

    d. none of the above.
1. When an object wants to do something, it uses a **method or function**.
1. In Java, accessing an object’s methods requires the **. (dot)** operator.
1. In Java, boolean and byte are **primitive type**.


## Referensi

- Why are Data Structures And Algorithms important? - <https://www.safaribooksonline.com/library/view/from-0-to/9781788626767/>
- Overview and Java Review - 1.1 Overview and Java Review <https://www.safaribooksonline.com/library/view/data-structures-and/9780763757564/chap01.xhtml>
- Data Structures and Algorithms in Java, 2nd Edition - <https://www.safaribooksonline.com/library/view/data-structures-and/9780134849775/ch01.xhtml#ch01>