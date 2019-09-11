# Distributed Tracing


## Apasih alasan kenapa ada distributed tracing?

Ya jawaban paling rasional karena sudah jamannya, Dimana ada kebutuhan tracing atau pelacakan untuk semua service dari aplikasi kita. Diawali dari maraknya istilah seperti microservice dan sekarang Cloud Native. Ya kalau belum microservice gak perlu distributed tracing juga sih.
Intinya adalah kita tidak menginginkan ketika aplikasi kita mati atau tidak berjalan kita tidak bisa mengindentifikasinya karena tidak ada jejaknya.

### Apa itu Microservice?

Tidak ada definisi resmi sih, tapi kalau mengambil pendapat dari [Martin Fowler](https://martinfowler.com/articles/microservices.html) ada 8 karakteristik dari MicroService. yaitu:

1. Componentization via Services
2. Organized around Business Capabilities
3. Products not Projects
4. Smart endpoints and dumb pipes
5. Decentralized Governance
6. Decentralized Data Management
7. Infrastructure Automation
8. Design for failure
9. Evolutionary Design

Lalu apa hubunganya, menelaah karakteristik yang pertama, dimana setiap service akan terpisah dari service lainnya sehingga ini akan menambah banyak masalah. Seperti latency, identifikasi dimana service yang mati dimana, dan lain lain. Ya gak masalah kalau servicenya cuma sedikit dalam hitungan jari, bisa langsung ketemu masalahnya dimana, bagaimana kalau seperti gambar ini.

![alt text](/images/visual-uber-microservice.jpg "Gambar dari buku mastering distributed tracing")

Gambar diatas adalah gambar dari Jaeger UI DAG(Directed Acyclic Graph). Yang merepresntasikan microservice yang ada di uber, besar dari lingkaran / node itu merepresentasikan service yang paling banyak digunakan atau semua bergantung pada service itu.

Kalau dilihat dari kompleksnya gambar diatas, kira kira apa aja tantangannya?

1. Orkestrasi, pastinya kita butuh otomatisasi dari deploy container, auto scale, restarting service yang bermasalah, dan lain lain. Tapi ini udah dihandle sama kubernetes. Gak mungkin handle itu secara manual. Pasti yang jagainnya abis onboarding langsung resign.
2. Komunikasi, Gimana cara service mengetahui service lain di dalam jaringan, gimana cara load balancing?, gimana cara implementasi rate limiting?. Tapi ini akan dihandle oleh services meshes atau network proxy, ini ada di jalur pembelajaran cloud native kita? siapa yang mau explore ini?
3. Seperti yang kita sama sama mengerti kalau membuat sistem dari monolith ke microservice itu bakalan mengurangi ke handalan system / reliability. Misalnya kita punya satu request yang menjalankan 30 proses, kalau dimonolitik satu saja error maka akan error semuanya. Beda dengan di microservice. Satu error mungkin sistem tidak akan error semuanya tetapi kita akan sulit untuk menemukan penyebab error tersebut. Mungkin untuk antisipasinya kita bisa implemntasi retry untuk memaintain ketersediaan service. ketika ada error kita coba ulangi requestnya.
4. Masalah selanjutnya adalah latency, sudah pasti latency akan bertambah ketika memproses semua service.
5. Bagaimana cara kita memonitor jika hanya mengandalkan tradisional monitoring.

Dan lalu ketika ada request yang lambat atau bermasalah bagaimana cara kita meneliti penyebab lambatnya? dan Bagaimana kita menjawab pertanyaan-pertanayaan dibawah ini?

- Service apa apa saja sih yang dilewati sebelumnya? Apa ada transaksi? Misal reserved balance, Redeem Coupon, dll.
- Apa yang dilakukan service sebelumnya? Kalo reserved balance dia akan mengurangi dari balance sementara. User akan melihat balancenya ter-debet.
- Kalo requestnya lambat, siapa yang jadi bottleneck, kenapa bisa itu terjadi?
- Kalo requestnya error kenapa bisa itu bisa terjadi, karena apa?
- Bagaimana kita tau bedanya antara sistem yang normal dan gak normal?
  - Apakah ada penambahan service baru? Dan membuat beberapa service yang dibutuhin gak kepanggil?
  - Apakah karena ada service yang terlalu lama memproses?
- Atau mungkin kasarnya, tim apa yang harus dipanggil untuk menghandle ini?

Semua itu tidak dapat terhandle hanya dengan mengandalkan traditional monitoring

Mungkin kalian bertanya, Yang dimaksud dengan traditional monitoring itu apa?

Intinya adalah monitoring hanya dapat menghandle per instance atau per service saja. tidak menghiraukan konteks dari requestnya. Semisal ketika PPOB system berjalan, mungkin akan ada beberapa service yang dilewati, misalnya, inquiry data ppob, menambah poin untuk user, mendebet balance di ovo sementara. membeli ke ppob provider lain. jika gagal mengembalikan balance yang sudah didebet. Jika berhasil mencatat kembali jika transaksi berhasil. Traditional monitoring hanya mencatat per service saja, tidak ada konteks apakah proses menambah poin untuk user karena transaksi PPOB atau transaksi dari Purchase Order.

Maka dari itu kita perlu memiliki sistem yang dapat diteliti / Observability.

## Apa itu observability?

Kalau kita ambil dari term observability, maka sistem bisa dikatakan observable ketika?

> Kondisi dan perilaku dari suatu sistem dapat ditentukan hanya dengan melihat input dan output saja.

Bryan Cantrill, CTO Joyent berpendapat definisi diatas itu tidak dapat dipraktekan di software, kenapa? karena hanya dengan membaca input dan output itu terlalu kompleks sehingga kita tidak bisa mengetahui keadaan sistem secara lengkap dan karena itu cara mengukur observability jawabannya selalu nol. (tidak bisa diukur, tidak ada jawaban).

Jadi definisi yang seharusnya yaitu, observability adalah "kemampuan untuk memungkinkan manusia untuk bertanya dan menjawab pertanyaan". Lebih banyak pertanyaan yang bisa kita jawab, artinya sistem kita lebih observable. itu sudah.

Lantas apakah yang dimaksud dengan observability sama seperti maksud dari monitoring?

Mengacu pada artikel ini [monitoring vsobservability](https://www.fastly.com/blog/monitoring-vs-observability).

Monitoring adalah aktifitas yang mengamati suatu sistem setiap waktu. Dan otomatis tidak membutuhkan manusia. Monitoring dapat memberi peringatan ke kita, jika terjadi sesuatu yang kita tidak inginkan. Semisal service yang dijalankan kehabisan memori. 

Disana juga disebut tentang 3 pillar dari observability, yaitu Logs, metrics dan traces. Akan tetapi tidak lama disanggah oleh Ben Sigelman dalam artikelnya [three pillars zero answers towards new scorecard observability](https://lightstep.com/blog/three-pillars-zero-answers-towards-new-scorecard-observability/).

Keresahannya ada di ketiga pilar tersebut, pertama tentang metrics, Metrics dapat sangat membantu ketika tags yang digunakan tidak terlalu banyak.

TODO: apa itu tags, show demo

Lalu Logs, keresahan di logs dari pendapat Ben adalah Mahal. Hitungan seperti ini

- Transaction rate
- x all microservice
- x cost of net + storage
- x weeks of rentetion

Lalu Distributed Tracing membutuhkan banyak sampling data. Yamh mengakibatkan biaya dari tracing akan menjadi mahal. Pendekatannya bisa diligat di artikel ini [Sampling in Observability - Jaana B. Dogan](https://medium.com/observability/sampling-in-observability-db0142cdda5b)

Videonya bisa lihat disini [https://www.youtube.com/watch?v=TLv1YhGNVGU](https://www.youtube.com/watch?v=TLv1YhGNVGU)

Ben Sigelman, menambahkan 2 penilaian, untuk menjawab keresahan diatas. yaitu Detection (deteksi) dan Refinement (perbaikan).

Deteksi:
    - Specificity: Unlimited Cardinality untuk seluruh stack.
    - Fidelity: Status yang benar dan frekuensi stats yang tinggi
    - Freshness: dibawah 5 detik

Refinement:
    - Identifying variance: Unlimited Cardinality, hi-fi histograms, data rentetions
    - "Suppress the messengers"

    Faktanya user kita gak peduli dan gak mau tau kita implementasi mikroservice ato engga. Yang penting aplikasi berjalan dengan baik tidak ada error. Dengan menambahkan mikroservis kita akan terus menghadapi banyak kegagalan. Intinya dari refinement ini adalah kita harus mengurangi lama pencarian kesalahan tersebut.

Tapi ketiga pillar tersebut hanyal tools atau instrumen untuk mengumpulkan data.

"If you want to talk about (metrics, logs, and traces) as pillars of observability–great.
The human is the foundation of observability!"

-- Brian Cantrill

## Observability adalah tantangan dalam microservice

Apa alasan atau apa yang diharapkan oleh perusahaan atau organisasi untuk mengimplementasikan Microservice?

Pasti harapannya adalah skalibilitas yang lebih baik dan lebih produktif.

Tapi jangan salah Microservice itu punya tantangan tersendiri dan kompleks, Dan kenapa google, facebook, netflix atau twitter mengimplementasi microservice dengan sukses ya karena dibaliknya banyak sekali orang pintar yang dapat menemukan  jalan yang paling efisien untuk memecahkan kompleksitas itu.

Menurut pendapat Vijay Gill, Senior VP of Engineering at Databricks dari artikel ini [the only good reason to adopt microservices](https://lightstep.com/blog/the-only-good-reason-to-adopt-microservices/)

"the right way to scale your engineering organization because you‘re going to ship your org chart, no matter what."

Menurut data tentang tantangan microservice di 2018

• Expect challenges and changes on the microservices journey:

   - 99% report challenges with using microservices
   - 56% say each additional microservice increases operational challenges
   - 73% find troubleshooting is harder in a microservices environment
   - 98% of those that face issues identifying the root cause of issues in microservices environments report it has a
   direct business impact
   - 87% of those in production report microservices generate more application data

Referensi Tambahan:

- [Metrics, tracing, and logging](https://peter.bourgon.org/blog/2017/02/21/metrics-tracing-and-logging.html)
- [Distributed Systems Observability](https://learning.oreilly.com/library/view/distributed-systems-observability/9781492033431/ch04.html)
