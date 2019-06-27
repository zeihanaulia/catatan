# Mengingat kembali javascript

Javascript pertama kali muncul pada tahun 1995. Sekarang ditahun 2019 sudah 24 tahun usianya dan sudah banyak sekali mengalami perubahan. Dulu biasanya javascript digunakan hanya untuk pemanis website seperti membuat tombol atau membuat validasi form. Sekarang semenjak ada Node.Js javascript sudah menjadi bahasa pemrograman yang "sebenarnya" karena bisa untuk membuat aplikasi.

Javascript dikelola oleh komite European Computer Manufacturers Association (ECMA), Tugasnya mereka adalah mengatur dan memprioritaskan proposal lalu memutuskan untuk dimasukan kedalam spesifikasi [ECMAScript](https://www.ecma-international.org/memento/tc39-rf-tg.htm). [Semua orang boleh memasukan proposal](https://tc39.es/process-document/) untuk perubahan ECMAScript ke komite ECMA. Bisa dicek [github](https://github.com/tc39)-nya

Untuk mengejar ketinggalan saya coba mebahas sedikit sejarah dari ECMAScript. Saat ini ECMAScript sudah ada ECMAScript 1 - 9. ES1 dikeluarkan pada tahun 1997 dan yang terakhir ES9 pada tahun 2018. Dari list tersebut hanya ES4 yang tidak jadi dikeluarkan karena terjadi banyak [masalah](https://auth0.com/blog/the-real-story-behind-es4/). Dan dilanjutkan ke ES5 ditahun 2009 yang membawa banyak fitur seperti properti object, array dan mendukung [json](https://en.wikipedia.org/wiki/JSON).

Dan ditahun 2015 keluarlah ES5 dan semakin menarik dunia per-JS-an ini, kenapa? karena semakin banyak orang-orang memberi proposal ke komite dan ketika banyak orang tertarik pada fitur itu akan dilakukan pengerjaan pembuatan fitur. Tetapi setelah semua fitur itu selasai, tergantung ke para penyedia browser seperti chrome, Firefox, safari dan lain-lain, Apakah mereka ingin mengiplementasikan fiturnya atau tidak. Lalu bagaimana cara agar kita tau kapan fitur yang kita inginkan akan dapat digunakan di browser yang ingin kita targetkan. Itu bisa kita pantau situs ini [http://kangax.github.io/compat-table/esnext/](http://kangax.github.io/compat-table/esnext/), bisa dilihat sekarang sudah masuk ESNext ditahun 2019 ini yang tapi belum secara resmi dikeluarkan. 

Secara pribadi saya sudah banyak ketinggalan, mungkin saya hanya tau javascript yang disebelum versi ES6 atau ES2015. Jadi berikut saya coba cari perubahan apa yang ada di ES6 ini agar bisa mengejar ketinggalan.

## Deklarasi variabel

Yang saya tau untuk mendeklarasi di javascript bisa menggunakan `var`. Tetapi setelah es 2016 kita bisa menggunakan kata kunci lain seperti `const` dan `let`. Apa sih bedanya?

### Konstanta

Konstanta, seperti pada bahasa pemrogramannya lainnya. Konstanta digunakan untuk menyatakan bahwa variabel tersebut tidak dapat ditiban atau diubah dengan value lain. Contoh

```js
var coffee = true;
coffee = false
console.log(coffee) // false
```

Variable diatas akan mencetak `false` karena kita menugasi variable `coffee` dibaris kedua dengan `false`, sehingga penugasan variabel `coffee` dibaris pertama menjadi berubah.

jika kita ganti `var` dengan `const` kita akan mendapatkan error seperti ini.

![alt text](/javascript/const-error.png "Error Konstanta")

### let

Let, ini cukup baru buat saya, yaitu mendeklarasikan variable yang hanya ber-efek didalam cakupan terntentu, Seperti didalam cakupan kurung kurawal {}. Langsung contohnya:

```js
var coffee = "Arabica";
if (coffee) {
    var coffee = "Robusta";
    console.log("Didalam ", coffee) // Diluar Robusta
}
console.log("Diluar ", coffee) // Diluar Robusta
```

Dari contoh diatas, maka hasil dari cetak yang ada diluar `if` akan mencetak `Diluar Robusta`. Padahal kita inginnya hanya mengubah variable `coffee` yang didalam saja. Untuk menyelesaikan masalah diatas maka kita gunakan variabel `let`

```js
var coffee = "Arabica";
if (coffee) {
    let coffee = "Robusta";
    console.log("Didalam ", coffee) // Diluar Robusta
}
console.log("Diluar ", coffee) // Diluar Arabica
```

Dengan seperti ini cetak yang diluar dari cakupan `if` akan mencetak `Diluar Arabica`

### Mengabungkan `string`

Ini yang menurut saya menarik sih, karena kalo dulu itu ribet banget untuk menggabungkan string yang cukup banyak, harus dengan menggunakan `+` sekarang cukup mudah dan enak dibaca dengan menggunakan *template strings*. Contohnya

```js
console.log(firstName + " " + lastName);
```

Sekarang cukup dengan menggunakan `backtick (``)`, gini;

```js
console.log(`${fistName} ${lastName}`);
```

Menarik bukan? contoh diatas masih dalam kata-kata yang sederhana, bagaimana dengan kata yang cukup banyak seperti membuat surat atau yang berisi html? fitur ini akan sangat membantu karena kita bisa dengan mudah membacanya.

## Membuat Fungsi

Fungsi (`function`) adalah kata kunci yang digunakan untuk membungkus kode agar dapat dipanggil dan dapat digunakan berulang-ulang. Semua bahasa pemrograman memiliki fitur ini, Tapi ada yang menarik di ES 6 yaitu adalah `Arrow Function`.

## Arrow Function

Fungsi Arrow/Panah (=>) adalah cara membuat fungsi tanpa kata kunci `function` melainkan `=>`, contohnya:

biasanya kita membuat fungsi seperti ini:

```js

const brew = function(coffee) {
    return `brewing coffee ${coffee}`
}

console.log(brew('Arabica'));
console.log(brew('Robusta'));

```

dengan fungsi arrow kita dapat membuatnya menjadi 1 baris alih-alih 3 baris seperti diatas. kita jadi menghilangkan kata kunci `function` dan `return` karena alaminya fungsi arrow akan langsung mengembalikan.

```js
const brew = coffee => `brewing coffee ${coffee}`
console.log(brew('Arabica'))
```

Kalau ada 2 argumen juga bisa dengan menambahkan kurung bukan dan tutup / parentheses `()`, seperti ini

```js
const brew = (coffee, method) => `brewing coffee ${coffee} with ${method}`
console.log(brew('Arabica', 'V60'))
```

Bagaimana jika menggunakan konsisi `if/else`, seperti ini

```js
const brew = (coffee, method) => {
    if (!coffee) {
        console.log('coffee is require');
        return
    }

    if (!coffee) {
        console.log('brewing method is require');
        return
    }

    return `brewing coffee ${coffee} with ${method}`
}
```

Bagaiana kalau ingin mengembalikan object, seperti ini:

```js
const coffee = (name, varietal, altitude, proccess) => ({
    name: name,
    varietal: varietal,
    altitude: altitude,
    proccess: proccess
})

console.log(coffee('Finca Maman', 'Geisha', '1400-2000 mts', 'Natural'))
```

dengan bantuan parentheses, kita dapat mengembalikan object. Sekarang sudah dapat update terbaru tentang membuat function di ES6.

## Spread operator

Seperti halnya ada dibeberapa bahasa pemrograman lainnya, di javasript terutama di ES7 atau ES2016 sudah ada fungsi yang dinamakan `spread operator` fungsinya adalah bisa mendapat variabel array yang tersisa. contohnya:

```js
const coffee = ['Arabica', 'Robusta', 'Excelsa', 'Liberica']
const [arabica, robusta, excelsa, liberica] = coffee
console.log(arabica) // Arabica
```

Kita gunakan operator spread

```js
const coffee = ['Arabica', 'Robusta', 'Excelsa', 'Liberica']
const [arabica, ...others] = coffee
console.log(others) //(3) [ "Robusta" , "Excelsa" , "Liberica"]
```

cukup menambahkan `...others` kita bisa mengambil semua array selain arabica. Yang menarik lagi, teryata operator spread melakukan `copy` alih-alih melakukan mutasi ke array `...others`. kita coba lakukan reverse, tanpa operator spread:

```js
const coffee = ['Arabica', 'Robusta', 'Excelsa', 'Liberica']
const [first] = coffee.reverse()
console.log(first) // Liberica
console.log(coffee) // ['Liberica', 'Excelsa', 'Robusta', 'Arabica']
```

Hasil yang didapat, kita bisa mendapatkan item pertama dari variable `coffee` yang dibalik dan jika kita cetak variabel `coffee`, isinya juga ikut terbalik atau termutasi. Lalu bagaimana solusinya? kita bisa menambahkan opertator spread agar variabel coffee di copy saja alih-alih dimutasi.

```js
const coffee = ['Arabica', 'Robusta', 'Excelsa', 'Liberica']
const [first] = [...coffee].reverse()
console.log(first)// Liberica
console.log(coffee) // [ "Arabica" ,"Robusta" , "Excelsa" , "Liberica"]
```

Selain itu masih ada manfaat lain, operator spread bisa digunakan untuk meyusun pengembalian dari REST API dalam bentuk json atau bisa juga digunalan didalam object.

## Asyncronus Javascript

Asyncronus adalah proses yang tidak langsung. Intinya adalah proses menunggu pada sebuah tugas. Misalnya ketika kita membutuhkan pengambilan data dari datatabase atau ketika kita mengambil data dari API lain. Pasti akan ada proses menunggu dalam pengambilan datanya. Oke, langsung dicontohin aja dengan memgambil data dari API [https://uinames.com](https://uinames.com).

Dulu waktu membuat pengamblan API kita butuh banyak menulis AJAX dengan baris kode yang lumayan banyak, sekarang kita hanya membutuhkan [fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API).

```js
console.log(fetch("https://uinames.com/?ext"));

// Promise {
// }
```

Setelah kita coba cetak ternyata `fetch` mengembalikan `Promise`, Apa itu Promise? Promise adalah cara agar kita dapat memahami asyncronus di javascript. `Promise` seperti namanya adalah memberi harapan atau janji, dia yang akan mengambilkan datanya dan berjanji jika ada terjadi sesuatu apapu dia akan mengembalikannya entah itu sukses atau gagal. `fetch` didalamnya disertakan juga method `then()`, method ini yang digunakan untuk menangkap pengembalian dari `Promise`.

```js
fetch("https://uinames.com/api/?ext")
.then(res => res.json())
.then(console.log)
.catch(console.error);
```

Lalu ada cara lain lagi, yaitu menggunakan Async/Await. Mungkin ini adalah cara yang lebih gamblang untuk digunakan, contohnya.

```js
const getRandomUser = async () => {
  try {
    let res = await fetch("https://uinames.com/api/?ext")
    let result = await res.json()
    console.log(result);
  } catch (error) {
    console.error(error);
  }
};
getRandomUser();
```

Kode ini prosesnya sama seperti kita menggunakan `then` hanya saja kita menggunakan kata kunci `async` diawal pembukaan fungsi dan `await` disetiap pemrosesan yang dilakukan di dalam cakupan fungsi yang di `async` kan. `Promise` membuat permintaan dengan asyncronus menjadi lebih mudah.

## Class

Di Javascript ada `Class` ? ini pertanyaan waktu pertama kali mendengar `Class` dijavascript biasanya dulu hanya bermain function-function saja. Entah harus senang atau sedih mendengarnya, kenapa jadi ikut ikutan bahasa lain yang ada Classnya. Tapi sebentar, biar kita lihat isi dari classnya.

Kelas di javascript menggunakan [Inheritance and the prototype chain](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Inheritance_and_the_prototype_chain).

```js
class Brewing {
    constructor(coffee, method) {
        this.coffee = coffee
        this.method = method
    }

    brew() {
        console.log(`Brewing ${this.coffee} using ${this.method}`)
    }
}

const coffee = new Brewing("Arabica", "Kalita Wave");

coffee.brew(); // Brewing Arabica using Kalita Wave
```

Saama seperti `Class` pada umumnya dimana dapat dilakukan perwarian. Di Javascript juga sama dengan menggunakan kata kunci `extends`

```js
class Parameter extends Brewing {
  constructor(coffee, method, parameter) {
    super(coffee, coffee);
    this.parameter = parameter;
  }

  brew() {
    super.brew();
    console.log(`Brewing with ${this.parameter.join(" and ")}`);
  }
}

const brewing = new Parameter("Arabica", "V60", [
  "Coffee 20gr",
  "Grind size medium",
  "Water 170",
  "Temperatur 90°"
]);

brewing.brew();

// 'Brewing Arabica using V60'
// 'Brewing with Coffee 20gr and Grind size medium and Water 170 and Temperatur 90°'
```

## Modules

Javascript modules nah ini bagian menarik juga, bagaimana javacript dapat melakukan export dan import function yang ada didalam file js, sehingga dapat dengan mudah dilakukan penggunaan kembali. Coba kita liat fiitur import dan export di Javascript itu seperti apa?

### Export

fitur export ini dipergunakan untuk mengexpor / menerbitkat fungsi atau kelas yang ada didalam suatu file. contohnya kita membuat file dengan nama `brewing.js`

```js
export brewing(coffee) => console.log(`Brewing ${this.coffee}`)
export const heating = "heating water"
```

arti kode diatas adalah kita menerbitkan fungsi `brewing` yang siap dipanggil dari manapun.

Ada satu kata kunci lagi yang dapat digunakan dalam melakukan `export` yaitu adalah `default`. `default` dapat digunakan ketika Anda ingin mengekspor hanya satu jenis tipe. Dan `export` dan `export default` dapat menggunakan tipe apa pun. contoh penggunaan dari `export default`, kita bisa letakan fungsi brewing parameter itu ada di file brew-parameter.js

```js
export default Parameter("Arabica", "V60", [
  "Coffee 20gr",
  "Grind size medium",
  "Water 170",
  "Temperatur 90°"
]);
```


## Import

Berlawanan dengan expor, kata kunci `import` digunakan untuk melakukan pemanggilan fungsi yang telah di `export`. Jika yang diexpor lebih dari satu fungsi, kita dapat melakukan destruturisasi obyek.

```js

import { brewing, heating } from './brewing.js'
import parameter as coffee from './brew-parameter.js'

brewing("Arabica')
console.log(heating)
coffee.brewing()
```

Kita juga dapat membuat alias dari variable yang kita impor. Atau kita juga bisa mengeluarkan semua fungsi yang diexpor dengan menggunakan simbol bintang (*).

```js
import * from './brewing.js'
```

Sekedar mengigatkat, `export` dan `import` belum sepenuhnya diimplentasi digunakan di browser. Maka dari itu kita membutuhkan pustaka tambahan untuk melakukan compiling atau melakukan konversi dari fitur yang belum di implemen dibrowser menjadi dapat kita gunakan. Pustaka itu dinamakan [babel](https://babeljs.io/)

## Babeljs.io

Seperti yang sudah disinggung sebelumnya, [babeljs](https://babeljs.io/) memproses sintak yang belum ada dibrowser dan lalu dikonversi ke sintak javascript yang sudah di implemen atau sudah ada di browser. Kita coba dengan membuka [https://babeljs.io/repl](https://babeljs.io/repl) dan pastekan code ini.

```js
const brew = coffee => `brewing coffee ${coffee}`
```

kita akan mendapatkan hasil konversi seperti ini:

```js
"use strict";

var brew = function brew(coffee) {
  return "brewing coffee ".concat(coffee);
};
```

Nah, Proses diatas umumnya akan dilakukan secara otomatis dengan menggunakan [webpack](https://webpack.js.org/) atau [parcel](https://parceljs.org/), Mungkin dilain tulisan saya akan membahas tentang webpackjs atau parceljs, karena belum sepenuhnya paham.

Referensi:

- [github.com/lukehoban/es6features](https://github.com/lukehoban/es6features)
- [developer.mozilla.org/en-US/docs/Web/JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript)
- [github.com/tc39](https://github.com/tc39)