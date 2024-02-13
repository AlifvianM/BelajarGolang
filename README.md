# BelajarGolang
Repositori untuk belajar golang dari 0 sampai dapat cuan

## Installasi golang
### Requirements
Usahakan menginstall Ubuntu menggunakan windows subsystem for linux supaya memudahkan penggunaan, penginstalan dan penyelarasan lingkungan kerja.

### Instalasi
* Update package ```sudo apt update``` & ```sudo apt upgrade```
* Download go ```wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz```
* Jalankan program berikut ``` rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz ```. Pada repositori ini kami menggunakan go versi 1.22.0
* Export lokasi go ke PATH environment ```export PATH=$PATH:/usr/local/go/bin```
* Check apakah go sudah terinstall ```go version```

# Memulai Project Baru
* Buat suatu folder dengan nama bebas
* Masuk ke folder ```cd <nama-folder>```
* Initialize folder ```go mod init <nama-folder>```
* Buat satu file dengan ekstensi ```<nama-file>.go``` lalu jalan kan file tersebut ```go run <nama-file>.go```
