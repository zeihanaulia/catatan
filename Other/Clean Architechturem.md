INTRODUCTION
Banyak dari orang orang yang baru berkecimpung di dunia programming, kode yang mereka produksi
tidak bagus, tetapi itu berkerja, itu berkerja karena: membuat aplikasi yang berkerja "sekali"
memang mudah.

Membuat software yang benar dan baik itu sulit, membutuhkan pengetahuan dan kemampuan yang dimana
para programmer baru tidak memilikinya. itu membutuhkan pemikiran dan wawasan, dan sebagian programmer
tidak pernah ada waktu untuk membangun kemampuannya. Membangun kemampuan membutuhkan disiplin
dan dedikasi, dan juga passion untuk membangun sebuah karya sebagai profesional

jika software di bangun dengan benar, itu hanya membutuhkan sedikit orang untuk membuat dan memelihara
perubahan menjadi mudah dan cepat.

WHAT IS DESIGN AND ARCHITECTURE?
Architecture = Design
Tujuan dari software Architecture adalah untuk meminifikasi tenaga kerja yang dibutuhkan untuk
membuat dan memelihara sistem.
Membuat code dangan benar, clean code,TDD dapat membantu untuk mempercepat produksi.
Membuat kode asal asalan , yang dibangung berdasarkan yang penting jadi akan mengambat produksi
selanjutnya.
satu satunya cara untuk melaju dengan cepat yaitu, buat sistem dengan baik.
membuat kode yang asalal asalan (kacau), pasti akan lebih lambat dari membuatnya dengan bersih.
disiplin untuk menjaga kode.

Chapter 2. A Tale of Two Values
ada dua nilai bagi software enginneer pertama adalah Behavior yang kedua adalah Architectuer
Behavior adalah perilaku umum yang bisa dikatakan software enginneer harus membuat system yang
bisa menghasikan uang untuk stakeholder

Architecture adalah ketika software enginneer harus membuat sistem dengan baik dan benar, sistem mudah
diberubah untuk kebutuhan stakeholder selanjutnya. sistem yang sulit untuk dimodifikasi akan menghambat
bisnis stakeholder.

nah kalo berdasarkan dari matrix mendesak dan penting maka:
Behavior = mendesak tapi tidak sangat penting
Architecture = sangat penting tapi tidak mendesak

dan sebagai software enginneer yang baik Architecture harus diperjuangkan, dan ingat sebagai software
enginneer, saya adalah pemilik dari sistem tersebut, yang memiliki andil dari software yang perlu
saya awasi. dan ini adalah bagian dari peran saya, dan bagian dari tugas saya.
Dan ini adalah sebagian besar alasan mengapa saya dipekerjakan.

Ingat, jika Architecture datang diakhir akhir, dan sistem akan menjadi lebih mahal untuk didevelop,
itu karena tim software enginneer tidak cukup berjuang keras untuk apa yang mereka kerjakan

PART II. Starting with the Bricks: Programming Paradigms
menceritakan ttg sejarah singkat bahasa pemrograman, dimana ada tokoh bernama alan turing yang
pertamakali memahami bahwa programming hanyalah data. dan dimana pada awanya programming dibuat hanya
dengan bahawa binary.

Chapter 3. Paradigm Overview
menceritakan ttg paradigma dari programming, selama ini ada 3 paradigma yang dapat diadopsi.

Structured Programming.
Paradigma ini yang paling terbaru, ditemukan oleh Edsger Wybe Dijkstra pada tahun 1968.
Dijkstra menunjukan bahwa statement jumps sangat menbahayakan, dan dia mengganti statement
tersebut menjadi syntax yang lebih familiar if/then/else and do/while/until.

alasan Dijkstra mengatakan bahwa goto bisa membahayakan karena they complicate two important
and related tasks: Proving that a program fragment is correct and describing what a program
has done so far.  We'll consider each of these in turn.
http://www.drdobbs.com/cpp/what-dijkstra-said-was-harmful-about-got/228700940

ringkasan dari paradigma Structured programming adalah
structured programming menerapkan disiplin on direct transfer of control.

Object oriented programming
dua tahun sebelumnya ditahun 1966 Ole Johan Dahl, dab Kristen Nygaard, memperkenalkan keduanya mencatat
bahwa fungsi pada bahasa ALGOL bisa menjadi bertumpuk, yang membuat fungsi local variable yang dideklarasikan
oleh fungsi yang sudah ada jauh setelah fungsi sudah dikembalikan.
fungsi menjadi constructor untuk class yang akan berisi variable, dan function lainnya
akan menjadi method. dan ini menjadi adanya polymophism yang mungginakan function pointer secara disiplin.

ringkasan paradima :
object oriented programming menerapkan disiplin on indirect transfer of control

Functional Programming
Paradima yang pertama diadopsi, paradigma ini ada sebelum adanya komputer,
functional programming merupakan kerjaan Alonzo Church  yang pada tahun1936 yang sedang
mengerjakan λ-calculus bersama dengan alan turing.

ringkasan paradigma:
functional programming menerapkan disiplin upon assignment.

kesimpulannya , apahubugannya dengan architectur, dan architecture akan menerapkan yang mana?
architecture akan menerapkan semuanya, kita menggunakan polymophism sebagai mekanisme untuk
melintasi batas batas arsitekturan. kita juga menggunakan functional programming untuk menerapkan
disiplin pada lokasi dan akses data.dan kita menggunakan structure Programming sebagai dasar
algoritma dari module yang kita buat.

dan itulah 3 hal besar yang penting dalam architecture, functiona, batasan dari komponen,data management

Chapter 4. Structured Programming



