# Kubernetes

[Kubernetes](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/) adalah sistem [orkestrasi](https://en.wikipedia.org/wiki/Orchestration_(computing)) kontainer. Kubernetes bisa disingkat K8S [K (ubernete) S]. Untuk mengerti kenapa K8S ada kita perlu memahami kenapa docker sang kontainer sistem diciptakan. Fungsi docker yang paling bermanfaat adalah dapat memisah-misah aplikasi yang tadinya [monolitik](https://en.wikipedia.org/wiki/Monolithic_application) menjadi [microservice](https://en.wikipedia.org/wiki/Microservices). Setelah kita paham fungsi kontainer, kita akan mendapatkan banyak sekali kontainer-kontainer, sehingga kita perlu memiliki alat yang dapat mengatur dan mengorkestrasikan kontainer itu dengan baik.

Orkestrasi contohnya gini, misalnya kita punya puluhan microservice artinya setiap service sama dengan puluhan kontainer, Apa yang terjadi ketika salah satu kontainer disana mati. Bagaimana kita menyalakannya kembali dan menidentifikasi kontainer yang mati. Bagaimana komunikasi antar service. Atau bayangkan [seperti ini banyaknya](https://www.theregister.co.uk/2014/05/23/google_containerization_two_billion/)

Untuk menyelesaikannya, Maka terciptalah K8S ini. Ada juga sih alat selain K8S.

Docker sendiri sudah terpasang docker swarm untuk orkestrasi, Setupnya pun lebih mudah dibandingkan dengan K8S. Tapi K8S lebih kaya fitur dan komunitasnya lebih banyak. Selain itu ada juga [Amazon ECS](https://aws.amazon.com/ecs/), [Apache Mesos](https://mesos.apache.org/), [Azure Service Fabric](https://docs.microsoft.com/en-us/azure/service-fabric/), [Crossplane](https://crossplane.io/), [Nomad](https://www.nomadproject.io/). Bisa langsung cek [disini](https://landscape.cncf.io/category=scheduling-orchestration&format=card-mode&grouping=category&selected=nomad)

Dengan K8S kita dapat mengatur kontainer sebanyak itu, hanya dengan mengelola text file berformat [.yaml](https://yaml.org). Dengan file yaml tersebut kita dapat mengatur apa pun tentang server yang kita perlukan, mulai dari mengembalikan container yang mati, menambah server, dan lain lain.

K8S dibuat digoogle dan sekarang dimanage oleh [Cloud Native Computing Foundation](https://www.cncf.io/).

## Minikube

Minikube adalah alat yang digunakan untuk menjalankan kubernetes di lokal komputer kita. Cara installnya cukup mudah bisa dilihat [disini](https://kubernetes.io/docs/setup/learning-environment/minikube/#installation).

## Docker

Docker adalah teknologi kontainer dimana aplikasi yang kita buat akan dimasukan kedalam kontainer.

Kontainer adalah deploying / penyebaran sistem yang paling popular saat ini.

Kenapa kita harus menggunakan docker atau kontainer ini? Bisa lasung cek [disini](https://docs.docker.com/get-started/)

## POD

Pod itu bisa disebut juga kaya pembungkus atau wrapper dari container. Jadi aplikasi kubernetes akan mengesekusi pod dimana didalam pod terdapat container yang berisi aplikasi kita.

Kita coba bahas konfigurasi sederhana.

```yaml

apiVersion: v1 # versi format yang digunakan
kind: Pod      # untuk membuat Pod
metadata:
    name: pod-example # menamakan Pod
    labels: # Labels adalah key value yang nantinya digunakan oleh services
        app: pod-example
spec: #spesifikasi yang digunakan per-Pod
    containers:
    - name: ubuntu          # nama container
      image: ubuntu:trusty  # image of docker ubuntu
      command: ["echo"]     # command
      args: ["Hello World"] # argument

```

Untuk mencoba kode diatas pastikan minikube sudah terinstal dan berjalan. Coba cek dengan sintak:

```bash
minikube status
```

Untuk menjalankan file yaml gunakan sintak:

```bash
kubectl apply -f <nama-file-yaml>
```

Untuk melihat deskripsi dari pod yang sudah kita deploy

```bash
kubectl describe pod <nama-pod>
```

Untuk melihat isi dari pod dan memberikan command

```bash
kubectl exec <nama-pod> ls
```

Untuk melihat isi dari pod dan memberikan command secara interaktif

```bash
kubectl -it exec <nama-pod> sh
```

Untuk menghapusnya

```bash
kubectl delete pod <nama-pod>
```

Hasil dari apply hanya dapat dipanggil dari dalam cluster kubernetes aja. Meskipun kita sudah tau IP nya, tetap belum bisa dibuka. Jadi jika misalnya kita membuat web app, untuk memanggilnya dibrowser kita gunakan ip yang bisa kita cek menggunakan minikube, sintaknya gini:

```bash
minikube ip
```

kita coba untuk mengetest dari dalam cluster dengan menggunakan interaktif exec.

```bash
kubectl -it exec <nama-pod> sh
wget http://localhost:80
```

Dengan cara itu, seharusnya kita bisa membuktikan bawah ternyata kita bisa mendownload file yang ada dicluster kita, dan webapp kita berhasil di deploy. Tetapi kenapa kita tidak bisa mengakses menggunakan api cluster dari browser kita?

Jadi ada mekanisme lain selain dari Pod, Yaitu Service


### Referensi Link

- [https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/](https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/)
- [https://kubernetes.io/docs/reference/](https://kubernetes.io/docs/reference/)

## Service

[Service](https://kubernetes.io/docs/concepts/services-networking/service/) kita gunakan untuk mengexpose Pod yang sudah terdeploy. Filosofinya Pod sifatnya hanya sementara akan dapat dimatikan dan dibuat lagi. 

Prosesnya gini 

```text
Service             ->      Pod
selector:                   labels:
    app: webapp                 app:webapp
```

Service memiliki selector dan Pod memiliki label. Service akan menugasi selector untuk mencatat pod dengan label.

Untuk mengexpose Pod maka kita membutuhkan service. Caranya dengan membuat file yaml lagi seperti ini misalnya

```yaml

apiVersion: v1 # versi format yang digunakan
kind: Service      # untuk membuat Service
metadata:
    name: service-example # menamakan service 
spec: #spesifikasi yang digunakan per-service
    ports:
    # Menerima trafik dan mengirimkan ke port 80
    - name: http          # nama port
      port: 80  # port yang diexpose
      targetPort: 80     # port yang diambil dari Pod

    selector:
      # Load balancing trafik yang melewati Pod
      app: pod-example # nama label dari POD

    # type load balancer hanya bisa digunakan di cloud provider
    # untuk dilocal kita bisa menggunakan ClusterID yaitu hanya bisa digunakan dengan id dari cluseter aja
    # atau kita bisa gunakan NodePort dimana kita bisa akses melalui browser dimana rangenya mulai dari 30000-32767
    # nodeport cuma temporary karena nantinya jika sudah dicloud menggunakan LoadBalancer
    type: LoadBalancer

```

## Bagaimana cara mengupdate app agar tidak ada down

Untuk melakukan update kita bisa saja mengganti secara langsung image yang ada dipod lalu `apply` lagi. Tapi itu akan memakan waktu karena ada proses build. Cara kedua dengan menggunakan label pada Pod dan service. Kita coba membuat Pod baru dengan membatasi dengan `---` dash / minus sebanyak tiga kali.

```yaml
# Pod

apiVersion: v1 # versi format yang digunakan
kind: Pod      # untuk membuat Pod
metadata:
    name: pod-example # menamakan Pod
    labels: # Labels adalah key value yang nantinya digunakan oleh services, dan value harus string
        app: pod-example
        release: "0"
spec: #spesifikasi yang digunakan per-Pod
    containers:
    - name: ubuntu          # nama container
      image: ubuntu:trusty  # image of docker ubuntu
      command: ["echo"]     # command
      args: ["Hello World"] # argument

--- # dash membatasi antara pod satu dan dua

apiVersion: v1 # versi format yang digunakan
kind: Pod      # untuk membuat Pod
metadata:
    name: pod-example # menamakan Pod
    labels: # Labels adalah key value yang nantinya digunakan oleh services
        app: pod-example
        release: "1"
spec: #spesifikasi yang digunakan per-Pod
    containers:
    - name: ubuntu          # nama container
      image: ubuntu:trusty  # image of docker ubuntu
      command: ["echo"]     # command
      args: ["Hello World"] # argument
```

```yaml

# Service

apiVersion: v1 # versi format yang digunakan
kind: Service      # untuk membuat Service
metadata:
    name: service-example # menamakan service 
    labels:
        release: "0"
spec: #spesifikasi yang digunakan per-service
    ports:
    # Menerima trafik dan mengirimkan ke port 80
    - name: http          # nama port
      port: 80  # port yang diexpose
      targetPort: 80     # port yang diambil dari Pod

    selector:
      # Load balancing trafik yang melewati Pod
      app: pod-example # nama label dari POD

    # type load balancer hanya bisa digunakan di cloud provider
    # untuk dilocal kita bisa menggunakan ClusterID yaitu hanya bisa digunakan dengan id dari cluseter aja
    # atau kita bisa gunakan NodePort dimana kita bisa akses melalui browser dimana rangenya mulai dari 30000-32767
    # nodeport cuma temporary karena nantinya jika sudah dicloud menggunakan LoadBalancer
    type: LoadBalancer

```

Tinggal kita lagukan re-apply service jika Pods versi baru sudah ready.

Untuk mengecek Pods kita bisa menggunakan sintak:

```bash
kubectl get pods                            # standart
kubectl get po                              # disingkat jadi po
kubectl get po --show-labels                # mencari pods dengan infomasi label
kubectl get po --show-labels -l release=0   # mencari pods dengan filter label
```

## ReplicaSet

Jadi kalau sebelumnya kita mendeploy langsung menggunakan pods, sekarang kita coba deploy ke service menggunakan ReplicaSet. Kenapa harus menggunakan ReplicaSet? Bayangkan ketika malam hari diwaktu kita tertidur tiba tiba pod kita mati. Maka kita harus bangun dan segera menyalakan kembali. Ya kalau itu ada yang memperhatikan dan memberi tau. Jika tidak gimana?

ReplicaSet akan menyelesaikan itu, caramnya kita akan mendeploy tidak langsung menggunakan pods, tetapi menggunakan replica set yang bisa kita seting berapa replica yang kita inginkan?

Struktur file yaml yang dibutuhkan seperti ini:

```yaml
apiVersion: apps/v1 # kenapa kok beda? ada apps nya? bisa dicek didokuentasi, jika core tidak perlu pake apps/ kalau tulisannya ada group apps, maka dibuat apps/v1
kind: ReplicaSet
metadata:
  name: webapp
spec:
  selector:
    matchLabels:
      app: webapp
  replicas: 2
  template: #template for the pods
    metadata:
      labels:
        app: webapp
    spec: #spesifikasi yang digunakan per-Pod
      containers:
      - name: ubuntu          # nama container
        image: ubuntu:trusty  # image of docker ubuntu
        command: ["echo"]     # command
        args: ["Hello World"] # argument
```

Dengan adanya ReplicaSet, kita dapat dengan tenang meninggalkan aplikasi, jika sesuatu terjadi yang menyebabkan pod kita mati ReplicaSet dengan cepat akan mengantikannya ke pods yang menyala dan merecreating ulang pods yang baru.

## [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#creating-a-deployment)

Ketika belajar ReplicaSet Di APInya ada warning, dibeberapa case baiknya menggunakan Deployment, apa itu deployment?. Sebenarnya Deployment, juga melakukan ReplicaSet, Hanya saja Deployment menambahkan kemampuan untuk bisa melakukan update versi dari pods.
Jadi cukup menganti Kind ReplicaSet menjadi Deployment.

```yaml
apiVersion: apps/v1 # kenapa kok beda? ada apps nya? bisa dicek didokuentasi, jika core tidak perlu pake apps/ kalau tulisannya ada group apps, maka dibuat apps/v1
kind: Deployment
metadata:
  name: webapp
spec:
  selector:
    matchLabels:
      app: webapp
  replicas: 2
  template: #template for the pods
    metadata:
      labels:
        app: webapp
    spec: #spesifikasi yang digunakan per-Pod
      containers:
      - name: ubuntu          # nama container
        image: ubuntu:trusty  # image of docker ubuntu
        command: ["echo"]     # command
        args: ["Hello World"] # argument
```

Jadi ketika ada perubahan versi dari image container dapat dengan mudah mengganti menjadi versi misalnya 1 atau 2 tanpa ada waktu berhenti bagi service kita.

redeploy

```bash
kubectl apply -f pods.yaml
```

untuk melihat status

```bash
kubectl rollout status deploy webapp
```

untuk melihat history rollout

```bash
kubectl rollout status deploy webapp
```

untuk membatalkan service yang telah dideploy ke sebelumnya

```bash
kubectl rollout undo deploy webapp --to-revision=2 # kalau tanpa --to-revision defaultnya kembali ke 1 versi sebelumnya
```

Prosess rollout ini cukup lama, tapi tidak membuat aplikasi kita menjadi berhenti karena ada rollout ini. Jadi kita dapat dengan mudah melakukan rollout ke versi berapapun tanpa ada aplikasi yang berhenti. Yang menarik lagi adalah, misalnya developer salah memberikan image yang akan dideploy. Kegagalan dalam membuat image tidak akan mengakibatkan aplikasi yang sedang berjalan di versi sebelumnya menjadi rusak juga. Secara otomatis K8S tidak akan melanjutkan proses untuk menghentikan versi sebelumnya.

## Namespace

SOON

## Menyimpan daya dihardisk fisik bukan kontainer

Istilah teknologinya itu volumes. Ketika membuat aplikasi tentunya terasa kurang kalau tidak melakukan penyimpanan data.
Tidak mungkin kita menyimpan didalam container, karena kontainer itu selalu berubah dan pastinya setiap ada pembuatan kontainer baru data akan hilang. Maka dari itu kita butuh melakukan `volumes` yang mana, mengcopy data dari container ke hardisk fisik komputer kita.

Sebagai contoh membuat deployment mongodb, kita sebut `mongo-stack.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mongodb
spec:
  selector:
    matchLabels:
      app: mongodb
  replicas: 1
  template: #template for the pods
    metadata:
      labels:
        app: mongodb
    spec:
      containers:
      - name: mongodb
        image: mongo:3.6.5-jessie             # image langung ambil aja dari dockerhub
        volumeMounts:
          - name: mongo-persistent-storage    # penamaan volume [container]
            mountPath: /data/db               # path apa yang mau di mount
      volumes:
        - name: mongo-persistent-storage      # penaaan volume [fisik]
          persistentVolumeClaim:              # persistent volume yang ingin di claim
            claimName: mongo-pvc              # nama claim (nanti akan aka yaml terpisaj)

--- # Tetap dibutuhkan service agat dapat diakses oleh kontainer lain
kind: Service
apiVersion: v1
metadata:
  name: fleetman-mongodb
spec:
  selector:
    app: mongodb
  ports:
    - name: mongoport
      port: 27017
  type: ClusterIP
```

Sebenarnya kiita bisa saja melakukan volume tanpa menggunakan Persistent Volume Claim, Hanya saja jika sudah banyak sekali kontainer yang kita maintain akan sangat merepotkan. Merepotkan dimana? Untuk contoh diatas baru 1 deployment satu database. Bayangkan jika ada puluhan dan kita ingin memindahkan ke cloud provider lain. Kita harus mengubah puluhan deployment hanya untuk mengubah setingan volumes persisten database kita. Jadi kita bikin saja satu file yaml lagi untuk pusat setingan. misal kita namakan `storage.yaml`

```yaml
# Apa yang kita inginkan?
# Maksudnya ini adalah permintaan ke kubernetes untuk memenuhi settingan dibawah
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mongo-pvc                   # nama ini yang digunakan untuk claim
spec:
  storageClassName: mylocalstorage
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 20Gi                 # permintaan untuk menyediakan storage sebanyak 20 gigabyte

---
# Bagaimana cara kita mengimplementasikannya
# Maksudnya ini adalah jawaban atas permintaan diatas
apiVersion: v1
kind: PersistentVolume
metadata:
  name: local-storage
spec:
  storageClassName: mylocalstorage
  capacity:
    storage: 20Gi                           # meminta ke kubernetes untuk menyetingkan 20 giga byte di cloud provider
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/mnt/some/directory/structure/"  # letak dari data fisik
    type: DirectoryOrCreate
```

