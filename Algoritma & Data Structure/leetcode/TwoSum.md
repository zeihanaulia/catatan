# Two Sum Problem

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

## Penjelasaan

Inti dari masalahnya begini, Temukan index dari array integer, yang kalau dijumlahkan hasilnya sama seperti target.

Jika targetnya 9 dengan value dari array dimulai dari index 0 (9 - (nums[1] = 2) = 7)

## Solusi

### 1. Simpan seluruh hasil pengurangan, untuk di cari pasangannya diiterasi selanjutnya

1. Definisikan hash map kosong untuk menampung index dan value dari hasil pengurangan
2. Buat perulangan untuk semua array integer
3. Ambil nilai integer dari array untuk dikurangkan dengan target
4. Masukan hasil pengurangan kedalam hash map.
5. Di iterasi selanjutnya ulangi step no 3 dan hasilnya dicari kedalam hash map dan masukan seperti no 4
6. Ulangi iterasi hingga hasil pengurangan bertemu

```bash

inputan = [2, 7, 11, 15]
target  = 9

definisikan map kosong []

perulangan sejumlah isi dari inputan = 3
ambil inputan yang indexnya 0 = 2
kurangi dengan 9 - 2 = 7
cari nilai 7 didalam map, jika tidak ada maka masukan nilai index 2 kedalam map. menjadi [2:0] key 2 value 0

ulangi iterasi selanjutnya dengan mengambil inputan yang indexnya 1 = 7
kurangi dengan 9 - 7 = 2
cari nilai 2 kedalam map , seharusnya ketemu karena diiterasi sebelumnya index 0 isinya adalah 2, jika sudah ketemu maka masukan nilai kedalam array dari integer untuk diisini nilai index dari key 2 yaitu 0 dan nilai index di iterasi sekarang yaitu 1, maka kita sudah mendapatkan hasilnya yaitu [0,1]

Di iterasi selanutnya, ambil inputan yang indexnya 1 = 7, lalu
```
