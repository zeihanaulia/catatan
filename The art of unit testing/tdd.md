# The three core of TDD (Test Driven Development) (Chapter 1 section 6)

1. Knowing how to write good test
1. Writing them first
1. Designing them well

Jika hanya karena menulis test diawal bukan berarti kode akan menjadi mudah dimaintain, dibaca, dan dipercaya. Lalu jika sudah menulis test dengan mudah dimaintain, dibaca, dan dipercaya bukan berarti kita bisa mendapat keuntungan yang sama seperti kita menulis testnya pertama kali.  Lalu jika sudah menulis test dengan mudah dimaintain, dibaca, dan dipercaya bukan berarti kode yang kita tulis akan memiliki desain yang bagus.

Rekomendasi buku bacaan:

- Kent beck'sTest-Driven Development: by Example (Addison-Wesley Profesional, 2002)
- Growing Object-Oriented Software, Guided by Test by Steve Freeman and Nat Pryce (Addison-Wesley Profesional, 2009)
- Clean Code by Robert C. Martin (Prentince Hall, 2008)

Nasihan cara belajar:

Sebaiknya ketiga hal tersebut dipelajari secara terpisah dan berurutan. Agar bisa fokus kesetiap inti dari TDD. Dengan mengambil pendekatan tambahan untuk mempelajari bidang ini, Anda akan menghilangkan rasa takut bahwa Anda salah dalam area yang berbeda dari yang Anda fokuskan saat ini.

## Catatan latihan

1. 1=1 satu class test mengetest satu class production, satu untit test project per test project, dan setidaknya satu unit method per pertest
1. penamaan class harus jelas UnitOfWork_Scenario_ExceptedBehavior
1. Gunakan factory method untuk mengulang kode. seperti membuat dan menginisialisasi object
1. jangan gunakan setup dan tear down. ini bisa membuat test jadi susah dimengerti.