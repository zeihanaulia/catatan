# Docker

Docker adalah software yang melakukan virtualisasi atau kita menyebutnya dengan istilah kontainer. Yang dibuat oleh perusahaan Docker Inc.

Cerita awal mula docker itu dimulai dari Solomon Hykes, mengerjakan project internal di dotCloud. Dan pada tahun 2013 di opensource kan [source code](https://github.com/docker/docker-ce)

Sekarang docker menjadi normal baru dalam kegiatan pemrograman.

## Masalah apa yang coba dipecahkan oleh Docker

Dulu gw pernah mengerjakan beberapa project website. Project ini dikerjakan dengan beberapa programmer lain. Setiap programmer memiliki setingan laptop yang berbeda beda. Sehingga yang terjadi adalah aplikasi yang gw buat bisa berjalan di laptop gw, tapi pas dicoba di laptop mereka ada yang bisa dan ada yang tidak. Ternyata versi bahasa pemrograman dilaptopnya berbeda.

Akhirnya, harus disetting ulang, disesuaikan dengan settingan laptop yang bisa tadi. Setelah beberapa menit atau mungkin beberapa jam melakukan setting laptop, baru lah bisa berjalan.

Tapi jadi ada masalah baru, yaitu. Project-Project dia jadi tidak ada yang jalan. Artinya kalau dia mau menjalankan lagi, dia harus ganti settingan lagi, atau dia harus update kodenya disesuaikan sama setingan barunya. rumit sekali.

Lalu muncul lah beberapa tools untuk menyelesaikan masalah itu, seperti VirtualBox, VMWare, Vagrant dan Docker.

Seperti yang kita tau, belajar dari pengalaman, gonta ganti settingan itu ide yang gak bagus, kenapa? karena terlalu rumit dan memakan waktu cukup lama, belum lagi jika ada kesalahan setting. Dan juga itu tidak skalabel.

Menggunakan virtual mesin itu setelah dicoba memakan banyak resources dan juga gak bisa di reuse dengan mudah. Maksudnya kita bisa tau settinganya, tapi susah buat langsung dibagi. tetep aja musti seting satu satu VMnya. Misal nih, kalian semua udah dapet settingannya. tapi pada saat development ada satu dari programmer yang membutuhkan settingan berbeda. Ini harus diinfokan juga dan yang lain harus ikut setting itu juga.

Lalu munculah Vagrant, Vagrant membuat semua setting setting menjadi lebih mudah, 
karena semua settingan ditulis didalam satu file dan akan dijalankan menjadi VM. 
Kita cukup share file itu aja. Tapi Vagrant gak menyelesaikan masalah besarnya resources yang dipakai. Karena masih menggunakan virtual machine. Tapi itu dulu sih, karena sekarang vagrant juga udah support docker, dan dia bisa manage docker container.

Nah apa itu docker, kenama bisa menyelesaikan masalah resource yang ada di VM. 
Bukannya docker VM juga?

Kalo ngomongin docker bakal banyak cakupannya, kaya networking, security, deployment.
Tapi kita akan cari tau dulu tentang containerization dan isolate deployment.

Docker itu benar benar ringan dibanding VM. Aplikasi bisa dipisahkan ke beberapa container. Artinya container bisa dideploy secara independen dan modular. 
Kita bisa deploy per container sendiri sendiri.

Docker itu extendable, Kita bisa membuat base image yang berisi settingan global diserver dan kita bisa extend di image lain.

Contoh:
[serverless-ci](https://cloud.docker.com/u/zeihanaulia/repository/docker/zeihanaulia/serverless-ci)

## Container VS Virtualzation

Virtualisasi membutuhkan resource yang berlebih dan membayar semuanya diawal.
Sedangkan container lebih ringan. Kita membutuhkan sesuai dengan apa yang dibutuhkan dan bisa scaling dengan lebih mudah.

Kalo virtual setup, biasanya dimulai dengan membuat host machine, bisa di server atau di laptop. VM punya virtual machine monitor yang bisa digunakan untuk membuat virtual machine yang memiliki virtual hardware dan software. Dan tentu saja VM memiliki operating sistem. 

Tapi ingat, VM itu cuma mensimulasikan environment hardware. Memecah dan mengalokasikan hardware yang dibutukan. yang harus diputuskan diawal.

Sedangkan Container, hanya membutuhkan os host, dan docker engine yang akan mengelola container. yang terkoneksi dengan kernel os host. 

Bukan cuma kernelnya linux, di windows juga bisa digunakan. Dan perlu diingiat container ini cuma bersifat sementara.

Nantinya diproduction, bisa ada banyak container yang sama akan berjalan berdampingan atau diganti tanpa ada gangguan pada service.

## Apa saja yang harus dimasukan ke container

Dalam membuat aplikasi pada umumnya kita install apa aja sih? Server, Database, Logging System, Monitoring System. Ya kan?

Bisa aja kita isi satu container dengan itu semua. Tapi itu jadi membingungkan. Semua diletakan dalam satu kontainer. Baiknya semua itu dipisah pisah menjadi beberapa container. sehingga setiap container memiliki domainnnya sendiri sendiri.

Aplikasi sendiri, Database sendiri, Log dan monitoring sistem sendiri. Database datanya juga jangan disimpan didalam container, tapi kita bisa menggunakan docker volume untuk menyimpan file database kita dihost os.

Ya, idealnya satu container cuma ngehandle 1 proses saja dalam sebuah sistem yang besar.

Misal kita punya satu aplikasi, dan aplikasi itu menyimpan data. Menyimpan data jangan di dalam container, melainkan gunakan docker voluma.
Ketika terjadi masalah pada container database dan kita butuk memperbaiki dan restart / rebuild container, data tidak akan hilang. sedangkan jika kita simpan di dalam container,setiap direstart dia akan menconfigurasi ulang dan mengapus semua data data.

Intinya, docker container itu haruslah kecil dan satu. Jangan ada beberapa hal didalamnya dan dapat selalu digantikan.

## Show me the code

Berapa lama biasanya kita setup web server?

```bash

docker run -d --name nginx1 -p 8080:80 nginx

```

Terus bisa gak kita bikin lagi?

```bash

docker run -d -p --name nginx2 8081:80 nginx

```

Sekarang kita mau liat container apa aja yang sedang jalan?

```bash

docker ps

```

Kalo ada container yang sedang gak jalan bisa diliat?

```bash

docker stop nginx1

docker ps -a

```

Kalo mau hapus container bisa?

```bash

docker rm nginx1

```

Kalo mau hapus paksa container yang lagi nyala bisa?

```bash

docker rm -f nginx1

```

Tadi kan container, kalo liat images gimana?

```bash

docker images

```

Images bisa dihapus?

```bash

docker rmi nginx

docker images

```

## Apa itu docker engine

Oke kita cari tau apa yang bisa bikin docker ini.
Ada beberapa komponen didalamnya ada server dan ada client, sama kaya bikin restful API.
Server berlajan dari komponen yang bernama dockerd. Tugasnya adalah membuat dan memanage docker object kaya container, images, swarm, networks, volume.

Server menunggu permintaan dari docker socket.

Docker client itu adalah cli, yang barusan kita coba untuk menjalain perintah perintah. Perintah itu yang akan dikirim ke Server melalui API.

API ini menperoses banyak tugas. Misal ketika kita ingin membuat container tetapi images gak ada dia akan mencarikan dulu di docker hub lalu mendownload lalu membuat images.

API memungkinkan kita untuk:

- create , run and control docker container
- build, manage, and commit, pull, and push Docker images
- read and monitor logs from containers
- build and manage networks
- create and control data volumes with persistent storage
- provision Docker swarms and scale services
- dll

Bisa kita lihat, docker bukan hanya tentang containers, tapi ada banyak hal yang bisa dilakukan dan sangan berguna untuk devops dan web developer.

Oke kita review ulang ya:

Containers adalah rumah dari semua hal hal yang dibutuhkan aplikasi dan berjalan dalam isolasi.

Container berbagi kernel host os yang membuat start cepat dan memakai RAM sedikit.

Docker hub, tempat dimana dikita menyimpan registri images. Docker file yang sudah kita buat bisa disimpan disana dan dibuild imagesnya, sehingga ketika ingin digunakan tinggal kita download saja.

Disana kita juga bisa mencari kumpulan image image yang kita butuhkan, misal postgres, python dan ubuntu os, dll

Docker compose, ini untuk memanage local environtment. ketika kita punya banyak container dalam satu aplikasi, docker compose ini bisa menjalankannya bersamnaan.

Misal,  tadi kita sudah tau ya, kalo contanier itu idealnya dibuat kecil kecil dan terpisah, kalo terpisah kaya gitu, ribet kan jalaninnya kalo harus satu persatu. nyalain contain erdatabase dulu, nyalain container log, nyalain container apps. dengan adanya docker compose, kita bisa jalankan semua dengan satu perintah saja. 
Docker compose disusun dengan file YAML. Dan docker compose juga sangat berguna untuk mendeploy multi stage, semisal ada development, staging dan production.

Docker images adalah layer yang membuat container.

Container logs, dimana kita bisa melihat logs yang ada di setiap container yang berjalan.

Data volume adalah untuk berbagi dan menyimpan data

Docker swarm dibutuhkan ketika kita mau melakukan scaling: Replication, load balancing, rolling image, and configuration updates, multi-host networking, and cluster management.

The Docker Machine virtual host yang mensetup si docker engine. Docker engine bisa berjalan di linux mac dan windows.

## Container VS Images

soon

## How to pull docker images

[Pull](https://docs.docker.com/engine/reference/commandline/pull/)

Docker pull adalah command yang dapat mengambil images dari [https://hub.docker.com](https://hub.docker.com) 

```bash

docker pull nginx:1.13.0

```

cek images

```bash

docker images

```

## Apa itu [Dockerfile](https://docs.docker.com/engine/reference/builder/)

Docker file itu cuma sebuat kumpulan instruksi. yang nantinya akan dijalankan oleh docker mechine.

## Dockerfile Structure

```Dockerfile

FROM ubuntu:16.10

LABEL maintainer="John Doe <john@doe.com>"

RUN apt-get update && apt-get upgrate -y

RUN apt-get install -y vim

```

struktur docker file selalu diawali dengan `FROM`. `FROM` itu adalah keyword dari docker
nantinya bisa aja kita bikin.

- `FROM images`
- `FROM images:tag`
- `FROM images@digest`

Contoh diatas menjelaskan kalau kita ingin ambil images ubuntu dengan tag/versi 16.10
Kita juga bisa menggunakan tag :latest untuk mendapatkan versi paling terbaru.

Lalu ada `RUN` ini juga keyword dari docker, yang merepresntasikan commandline. jadi kita bisa menjalankan perintah yang ada dicommand line ubuntu, kasus diatas adalah meminta install vim dengan mengupdate apt get terlebihd dahulu.

`FROM` `RUN` `RUN` itu adalah layer yang terpisah. Jadi image dibangun berdasarkan layer layer

## [Build Docker](https://docs.docker.com/engine/reference/commandline/build/)

Untuk menjalankan docker build cukup dengan jalankan

```bash
docker build .
```

Bisa gak kalo dockernya pake nama lain. misal Dockerfile.base? bisa dong

```bash
docker build -f ./Dockerfile.base
```

## Image History

melihat history yang terjadi pada build

```bash
docker images
docker history images image-id
```

## Inspect Docker

melihat informasi dari image

```bash
docker images
docker history images image-id
```

## Refactor Dockerfile

```Dockerfile
FROM ubuntu:16.10

LABEL maintainer="John Doe <john@doe.com>"

RUN apt-get update \ 
    && apt-get upgrate -y \
    && apt-get install -y \ 
    vim \
    nginx
```

Build dockerfile yang baru dan bandingkan.

Kalo terpisah kaya sebelumnya masalahnya dimana? Jadi bisa aja terjadi masalah di cache. Docker cache akan membaca cache layer `RUN apt-get update && apt-get upgrate -y` sehingga tidak diesekusi lagi. dan yang `RUN apt-get install -y vim nginx` akan tetap diesekusi. masalahnya, apt tidak terupdate lagi, ada dapat kemungkinan nginx versi lama.

## Copy dan Add

```Dockerfile

COPY /code/mysite /var/www/mysite
or
COPY ["/code/my site", "/var/www/mysite"]
```

Yang sebelah kiri adalah source / file code di local kita
sedangkan yang kanan adalah destinasi atau kita ingin menempatkan didalam container di folder apa.

Pastikan tanda `/` itu benar.

```Dockerfile
COPY /code/mysite /var/www/mysite  #/var/www/mysite
COPY /code/mysite var/www/mysite   #<WORKDIR>/var/www/mysite
COPY /code/my* /var/www  #ambil semua folder yang berawalan my
COPY /code/mysite/page?.html /var/www  #ambil page.html dan pages.html
```

Add sama seperti copy tapi ada tambahan. ADD bisa menganbil dari url

```Dockerfile
ADD http://www.web.com/css/main.css /var/www/mysite/temp.css
```
Add juga bisa melakukan extract file on the fly

```Dockerfile
ADD resource/file.tar.gz /usr/local
```

kalo pakai copy, musti ada beberapa instruction lagi misal

```Dockerfile
COPY resource/file.tar.gz /tmp
RUN tar -zxvf /tmp/resource/file.tar.gz -C /usr/local
RUN rm /tmp/resource/file.tar.gz
```

Tidak disaran kan begini

```Dockerfile
COPY ../../mysite /var/www/mysite
ADD ../../mysite /var/www/mysite
```

semuanya harus di current directory


## Environment Variable

Kita bisa menyisipkan environment variable seperti ini

```Dockerfile
ENV DOC_ROOT /var/www/mysite
ADD code/mysite ${DOC_ROOT}
```

Kita coba running

```Dockerfile
FROM ubuntu:16.10

LABEL maintainer="John Doe <john@doe.com>"

ENV DOC_ROOT /var/www/mysite

RUN apt-get update \ 
    && apt-get upgrate -y \
    && apt-get install -y \ 
    vim \
    nginx \
    php7.0 \
    && rm -rf /var/lib/apt/lists/*
ADD code/sites/mysite ${DOC_ROOT}
```

build error , cek error

```bash
mkdir -p code/sites/mysite
```

di windows

```bash

ls = dir

md code/sites/mysite
mkdir code/sites/mysite
```

build lagi.


Bikin pake docker file yang ambil file dari server

```Dockerfile
FROM ubuntu:16.10

LABEL maintainer="John Doe <john@doe.com>"

ENV DOC_ROOT /var/www/mysite
ENV JQUERY_VERSION 3.2.1

RUN apt-get update \ 
    && apt-get upgrate -y \
    && apt-get install -y \ 
    vim \
    nginx \
    php7.0 \
    && rm -rf /var/lib/apt/lists/*
COPY code/sites/mysite ${DOC_ROOT}
ADD https://code.jquery.com/jquery-${JQUERY_VERSION}.min.js ${DOC_ROOT}/js
```

## Gimana cara passing env

```Dockerfile
FROM ubuntu:16.10

LABEL maintainer="John Doe <john@doe.com>"

ENV DOC_ROOT /var/www/mysite
ARG JQUERY_VERSION=3.2.1
ENV JV ${JQUERY_VERSION}

RUN apt-get update \ 
    && apt-get upgrate -y \
    && apt-get install -y \ 
    vim \
    nginx \
    php7.0 \
    && rm -rf /var/lib/apt/lists/*
COPY code/sites/mysite ${DOC_ROOT}
ADD https://code.jquery.com/jquery-${JV}.min.js ${DOC_ROOT}/js
```

Lalu running

```bash
docker build --build-arg JQUERY_VERSION=3.0.0 .
```

## Gimana menghapus images

```bash
docker rmi image-id
docker image rm image-id
docker image $(docker images -q) #hapus semua image
```

image akan tidak bisa dihapus jika ada container yang mengguakannya
jadi harus diremove dulu containernya


## Gimana cara tag image

Tag image untuk memberikan nama pada docker yang kita build tadi. tadi kita sudah nge build image. sekara kita ingin mentag image terse but dengan cara 

```bash
docker tag image-id webserver:1.0
```

atau kita bisa buat pada saat build dengan cara

```bash
docker build -t <imagename>:<TAG> .
docker build -t webserver:1.0 .

# dengan repositoru
docker build -t <repositoryname>/<imagename>:<TAG> .
docker build -t zeihanaulia/webserver:1.0 .
```

## Register docker hub

```bash
docker login

# kalau punya registry sendiri

docker login -u username -p password http://registry-sendiri.com
```

## push docker image ke registru docker hub

untuk melakukan push ke registry, membutuhkan repository. Jadi kita bisa melakukan tagging untuk image yang belum ada repositorynya atau build ulang dengan tag beserta repository

```bash
docker build -t <imagename>:<TAG> .
docker build -t webserver:1.0 .

# dengan repositoru
docker build -t <repositoryname>/<imagename>:<TAG> .
docker build -t zeihanaulia/webserver:1.0 .
```

lalu

```bash
docker push zeihanaulia/webserver:1.0
```

## gimana cara pull images

biasanya yang push ke registry cukup devops saja. yang lainnya tinggal pull aja

```bash
docker pull zeihanaulia/webserver:1.0
```

## gimana cara update remote image

```Dockerfile
FROM ubuntu:16.10

LABEL maintainer="John Doe <john@doe.com>"
LABEL image_type="web server dengan nginx dan php 7"

ENV DOC_ROOT /var/www/mysite
ARG JQUERY_VERSION=3.2.1
ENV JV ${JQUERY_VERSION}

RUN apt-get update \ 
    && apt-get upgrate -y \
    && apt-get install -y \ 
    vim \
    nginx \
    php7.0 \
    && rm -rf /var/lib/apt/lists/*
COPY code/sites/mysite ${DOC_ROOT}
ADD https://code.jquery.com/jquery-${JV}.min.js ${DOC_ROOT}/js
```

lau build

```bash
docker build -t zeihanaulia/webserver:1.1 .

docker push zeihanaulia/webserver:1.1
```

## gimana cara running docker images

```bash
docker run zeihanaulia/webserver:1.1
```

umumnya ketika ngerun kita menjalankan dengan command tambahan

- i = interactive / kita bisa masuk ke bash docker
- d = detach / jalan di background
- t = untuk jalan di terminal, psudo TTY

jadi kalo kita 

```bash
docker run -it zeihanaulia/webserver:1.1
```

kita ada didalam container

```bash
docker run -d zeihanaulia/webserver:1.1
```
jalan dibackground

## build container dari registry

```bash
docker run -it zeihanaulia/webserver:1.1 bin/bash
```

## gimana cara menamakan container

kalo misal kita udah run container gampangnya kita bikin gini

```bash
docker rename nama_container nama_container_baru
```

kalau sekalian buat

```bash
docker run --name webserver -it zeihanaulia/webserver:1.1 bash
```

# Gimana cara stop container

```bash
docker stop webserver / container-id
```

## kalo mau nyalain lagi

```bash
docker start webserver / container-id
```

## kalo mau restart lagi

```bash
docker restart webserver / container-id
docker restart -t 20 webserver / container-id #restart dalam 20 detik
```

## Cara run command di container

```bash
docker run -it zeihanaulia/webserver:1.1 php -v
docker run -it zeihanaulia/webserver:1.1 php -a
```

## Gimana cara masuk ke container

setelah running kita juga bisa masuk kedalam container

```bash
docker run webserver -td zeihanaulia/webserver:1.1 bash
docker exec webserver bash
docker exec -it webserver bash

#ganti user
docker exec -it -u www-data webserver bash

#pakai env
docker exec -it -u www-data -e DOC_ROOT="/var/www/mysite" webserver bash
echo $DOC_ROOT
```

## menyeting user pada images

```Dockerfile
FROM ubuntu:16.10

LABEL maintainer="John Doe <john@doe.com>"
LABEL image_type="web server dengan nginx dan php 7"

ENV DOC_ROOT /var/www/mysite
ARG JQUERY_VERSION=3.2.1
ENV JV ${JQUERY_VERSION}

RUN apt-get update \ 
    && apt-get upgrate -y \
    && apt-get install -y \ 
    vim \
    nginx \
    php7.0 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR ${DOC_ROOT}

USER www-data:www-data

COPY code/sites/mysite ${DOC_ROOT}
ADD https://code.jquery.com/jquery-${JV}.min.js ${DOC_ROOT}/js
```

## bikin web server

```Dockerfile
FROM ubuntu:16.10

LABEL maintainer="John Doe <john@doe.com>"
LABEL image_type="web server dengan nginx dan php 7"

ENV DOC_ROOT /var/www/mysite
ARG JQUERY_VERSION=3.2.1
ENV JV ${JQUERY_VERSION}

RUN apt-get update \ 
    && apt-get upgrate -y \
    && apt-get install -y \ 
    vim \
    nginx \
    php7.0 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR ${DOC_ROOT}

USER www-data:www-data

COPY code/sites/mysite ${DOC_ROOT}
ADD https://code.jquery.com/jquery-${JV}.min.js ${DOC_ROOT}/js
EXPOSE 80 443
```

jalan kan 

```bash
docker run --name con-tes webserver

## publish
docker run --name con-tes -p 80:80 webserver

```

Running apache web server

```Dockerfile
FROM ubuntu:16.10

LABEL maintainer="John Doe <john@doe.com>"
LABEL image_type="web server dengan apache dan php 7"  #update labels

ENV DOC_ROOT /var/www/mysite
ARG JQUERY_VERSION=3.2.1
ENV JV ${JQUERY_VERSION}

RUN apt-get update \ 
    && apt-get upgrate -y \
    && apt-get install -y \ 
    php7.0 \ 
    && rm -rf /var/lib/apt/lists/* # remove nginx dan vim

WORKDIR ${DOC_ROOT} 

COPY code/sites/mysite ${DOC_ROOT}
ADD code/apache/sites-available/000-default.conf /etc/apache2/sites-available/000-default.conf
ADD https://code.jquery.com/jquery-${JV}.min.js ${DOC_ROOT}/js
EXPOSE 80

CMD apachectl -D FOREGROUND #run command
```

isi file  000-default.conf

```conf

<VirtualHost *:80>

    ServerAdmin webmaster@localhost
	DocumentRoot ${DOC_ROOT}

	ErrorLog ${APACHE_LOG_DIR}/error.log
	CustomLog ${APACHE_LOG_DIR}/access.log combined

</VirtualHost>
```

isi didalam code/sites/mysite

```php
<?php phpinfo();
```

```bash
docker run --name con-tes -d -p 80:80 webserver
```

setelah itu tag dan push


## docker volume

https://docs.docker.com/storage/volumes/