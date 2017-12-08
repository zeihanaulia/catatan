# Array data structure

## Arrays

Array adalah structure penyimpanan data yang paling umum digunakan, hampir disemua bahasa pemrogramman ada.

### Insert (tidak bisa diduplikasi)

Dalam array  insert sangat cepat, hanya membutuhkan satu langkah data baru sudah terisi.
Knp? karena insert item baru akan selalu mengisi cell kosong pertama dalam array

### Search (tidak bisa diduplikasi)

Mencari items dalam array membutuhkan banyak step, proses yang dilakukan lama karena, progam akan terus mencari di setiap cell hingga menemukan item yang ingin dicari, Sehingga akan memakan waktu lebih lama.

contoh:
dik:
1, 2, 3, 4, 5, 6, 7, 8

dit:
cari item 3.

Jawab:
array:
[0] = 1
[1] = 2
[2] = 3
[3] = 4
[4] = 5
[5] = 6
[6] = 7
[7] = 8

angka didalam [] adalah key dari array dan angka disebelah kanan adalah value
ketika proses pencarian program akan melooping array dan mencari nilai sesuai target

- apakah nilai index 0 sama dengan 3; tidak;
- apakah nilai index 1 sama dengan 3; tidak;
- apakah nilai index 2 sama dengan 3; ya;

untuk menemukan nilai tiga array perlu melakukan pengecekan ke index 0 dan satu.

### Delete (tidak bisa diduplikasi)

Delete juga sama seperti search, bahkan lebih lama dari search, karena, sebelum mengeksekusi penghapusan item. progam akan mencari di setiap cell, jika ketemu akan dihapus lalu satu persatu item cell setelahnya akan dinaikan menggatinkan item tang telah dihapus, begitu seterusnya

contoh:
dik:
1, 2, 3, 4, 5, 6, 7, 8

dit:
hapus item 3.

Jawab:
array:
[0] = 1
[1] = 2
[2] = 3
[3] = 4
[4] = 5
[5] = 6
[6] = 7
[7] = 8

angka didalam [] adalah key dari array dan angka disebelah kanan adalah value
ketika proses pencarian program akan melooping array dan mencari nilai sesuai target

- apakah nilai index 0 sama dengan 3; tidak;
- apakah nilai index 1 sama dengan 3; tidak;
- apakah nilai index 2 sama dengan 3; ya;
- hapus item 3;
- pindahkan nilai index 3 menjadi index 2;
- pindahkan nilai index 4 menjadi index 3;
- ...
- pindahkan nilai index 7 menjadi index 6;

untuk menemukan nilai tiga array perlu melakukan pengecekan ke index 0 dan 1 hingga menemukan nilai 3 ada di index 2. setelah ketemu lalu nilai index 2 dihapus, lalu ketika index 2 sudah kosong maka nilai index 3 dipindahkan ke index 2 dan seterusnya.

diatas adalah pengujian terhadap array yang valuenya tidak bisa diduplikasi. lalu bagaimana dengan yang bisa diduplikasi. Pastinya makin kompleks algoritmanya

## Refrensi

Data Structures and Algorithms in Java, 2nd Edition
https://www.safaribooksonline.com/library/view/data-structures-and/9780134849775/ch02.xhtml