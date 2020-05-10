# MongoDB Docker

Bagaimana cara bermain mongodb dan docker

## Playground

- Install mongodb

    ```bash
        <!-- latest -->
        docker pull mongo

        <!-- version 4.0.4 -->
        docker pull mongo:4.0.4
    ```

- Menjalankan mongodb di docker

    ```bash
        <!-- jalan kan docker (-d) secara background (-p) diport 27017-27019 dengan (--name) nama mongodb dari image mongo -->
        docker run -d -p 27017-27019:27017-27019 --name mongodb mongo
    ```

- Pastikan docker sudah berjalan

    ```bash
        docker ps
    ```

- Akses kedalam mongodb jika sudah berhasil berjalan

    ```bash
        docker exec -it mongodb bash
    ```

- Coba mejalankan perintah mongo db

  - Ketik mongo untuk masuk kedalam cli mongo db

    ```bash
        mongo
    ```

  - Kita coba melihat database yang ada

    ```bash
        show dbs;
    ```

  - Kita coba membuat database

    ```bash
        use yourdbname
    ```

  - Coba untuk memasukan collection dan datanya

    ```bash
        db.user.save({ firstname: "Zei", lastname: "Han" })
    ```  

  - Coba untuk mengambil data yang tadi dimasukan

    ```bash
        db.user.find({})
    ```

  - Coba untuk mengambil data dengan nama spesifik

    ```bash
        db.user.find({ firstname: "Zei" })
    ```

- Keluar dari console mongodb

    ```bash
        exit
    ```