# Distributed Tracing

>Kenapa harus ada Distributed Tracing?


Karena sudah masuk era Microservice dan Faas atau distributed system.

Beberapa tahun belakangan ini pengembangan perangkat lunak sudah berpindah ke era modern yang mana Cloud Computing dan Containerization Teknologi memungkinkan generasi baru desain sistem terdistribusi. Munkin pernah dengar istilah Microservice atau Functional as a Service? itu adalah salah satu pengaplikasian dari pengembangan perangkat lunak era modern.

> Apa itu Microservice?

Tidak ada definisi resmi sih, tapi kalau mengambil argumen dari [Martin Fowler](https://martinfowler.com/articles/microservices.html) ada 8 karakteristik dari MicroService. yaitu:

1. Componentization via Services
   
    Definisi komponen, menurut artikel:
    
    > Our definition is that a component is a unit of software that is independently replaceable and upgradeable.

    Artinya, Aplikasi yang kompleks dipecah menjadi beberapa komponen melalui servis yang secara independen mudah tergantikan dan diperbaharui

    Misalnya: Diawal-awal pembuatan aplikasi semua disatukan, istilahnya dikenal dengan monolith. Contoh Odoo. Seiring bertambahnya jumlah pengguna dan semakin kompleksnya bisnis perlahan aplikasi di odoo akan dipecah pecah kedalam service yang kecil kecil.

2. Organized around Business Capabilities
3. Products not Projects
4. Smart endpoints and dumb pipes
5. Decentralized Governance
6. Decentralized Data Management
7. Infrastructure Automation
8. Design for failure
9. Evolutionary Design

