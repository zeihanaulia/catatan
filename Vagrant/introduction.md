# Introduction & Install

## Apaitu virtualisasi?

Virtualisai disini konteksnya hanya sebatas lingkungan pengembangan software saja. Kehidupan sebelum adanya vitualisasi seperti ini:

- Install semua infrastructur didalam komputer
- Ada beberapa komponent yang harus dibeli terlebih dahulu misal, jaringan, kabel, atau host
- akan sulit jika kita membuat beberapa project dengan environtment berbeda
- dan semuanya hanya menambah biaya dan waktu saja

Jadi apa itu virtualisasi?

- jaman dulu virtualisasi hanya ada didalam server tapi sekarang sudah ada didesktop
- virtualisasi mengacu kepada kemampuan mesin untuk memiliki multiple virtual mesin yang berjalan besamaan
- jadi misalnya kita punya komputer dengan os Windows 10 lalu menggunakan Virtual box atau VmWare untuk menjalankan os lain seprti ubuntu, red hat dan lain lain
- Virtual mesin juga memilikin memory network card yang terpisah dari komputer aslinya

Tetapi itu saja belum cukup

- Memang virtualisasi bisa memudahkan developer
- tapi seiring cepatnya perubahan pada pengembangan perangkat lunak dan website, kalo cuma virtualisasi aja gak cukuk, kenapa?

    1. Aplikasi menjadi lebih kompleks dengan perkembangan teknologi yang baru
    1. Server Configurasi juga makin kompleks
    1. Jamannya Agile development, yang mana membutuhkan pembuatan yang cepat dan deploy aplikasi di beberapa environtment yang berbeda
    1. cepat untuk mengadopsi trend devops terbaru, dimana sejarahnya batasan antara developers, system administrator dan operator menjadi blur
    1. dan virtualisasi akhirnya semuanya membutuhkan automatisasi dengan docker dan vagrant

## Apa itu vagrant?

Vagrant adalah tools yang memanfaatkan virualisasi, untuk membuat environtment development secara lengkap. semuanya di otomatisasi sehingga menghemat waktu bagi developer untuk melakukan setup infrastructure.

Ketika menggunakan vagrant kita akan menghemat waktu dan usaha kita, apaja yang kita hemat, ini:

- Membuat virtual mesin disemua operating sistem
- Configurasi VM resurces (Memory, CPU, network, dll)
- Tersedia juga share folder diantara host machine dan virtual machine jadi kita bisa mengubah file secara langsung.
- memberikan hostname ke mesin
- boots up the machine
- Deploy software ke mechine menggunakan script atau configuration management system seperti chef atau puppet
- bagaimanapun tugas inti dari vagrant ada dibagian deployment: semuanya bisa masuk kedalam oprating sistem apapun

Tapi, bukannya itu kaya docker ya? terus apa bedanya sama docker?

- Docker adalah container base solution untuk mengirimkan bundel aplikasi dengan infrastructure. satu keunggulannya docker lebih sedikit menggunakan resouces dibandingkan vagrant
- bagaimanapu solusi dengan menggunakan kontainer mengutilisasi host system kernel agar itu berkerja, jadi ketika deploying docker kemesin lain membutuhkan build ulang lagi agar bisa berjalan.
- intinya mau menggunakan docker atau vagrant gak masalah yang penting tujuannya sama

## Membuat machine

### Cara lama membuat virtual mesin

1. Buka virtual box 
1. Klik button new beri nama misal, webserver 
1. Pilih tipe linux
1. Pilih version ubuntu
1. Lalu klik continue
1. Seting memory sesuai keinginan lalu klik continue
1. Pilih Create a virtual hardisk now
1. Pilih VDI (Virtualbox Disk Image)
1. Pilih Dinamic Allocated
1. Setting Hardisk lalu klik create

Masih banyak lagi caranya, diatas hanya cara membuat virtual machine belum dengan infrastructur development

### Cara Vagrant

Kita hanya perlu mengertik sintax seperti ini

```bash
vagrant init bento/ubuntu-16.04; vagrant up --provider virtualbox
```

syntax diatas penjelasannya seperti ini

- vagrant init : memberi perintah kepada vagrant untuk menginisialisasi atau menginstall bento/ubuntu-16.04. darimana kita tau bento/ubuntu tersebut, bisa dilihat di https://app.vagrantup.com/bento/boxes/ubuntu-16.04
- vagrant up : memberi perintah kepada vagrant untuk menjalankan virtual mechine
- --provider: dengan tambahan syntax ini vagrant bisa memilih providernya, saat ini kita menggunakan virtualbox tapi bisa saja diganti dengan vmware atau hyperv
