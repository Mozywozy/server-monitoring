# Server Monitoring

Proyek ini adalah aplikasi monitoring server sederhana yang memantau status beberapa server dan mengirimkan email pemberitahuan ketika salah satu server tidak dapat dijangkau.

## Fitur

- Memantau status server menggunakan ping.
- Mengirim email pemberitahuan ketika server tidak dapat dijangkau.
- Mengkonfigurasi interval monitoring.
- Menggunakan SMTP untuk pengiriman email.

## Persyaratan

Sebelum menjalankan aplikasi, pastikan Anda memiliki:

- Go (versi terbaru)
- Akses ke server SMTP (misalnya, Gmail)
- Alamat email yang valid dan sandi aplikasi untuk autentikasi.

# Cara buat password khusus aplikasi

## 1. Membuat Password Khusus Aplikasi
- Masuk ke Akun Google Anda:

## 2. Kunjungi Akun Google dan masuk dengan kredensial Anda.
- Akses Keamanan:
- Di menu sebelah kiri, klik Keamanan.
- pastikan bahwa Verifikasi Dua Langkah sudah diaktifkan.

## 3. Buat Password Aplikasi:
- Klik pada opsi Buat password aplikasi.
- Pilih jenis aplikasi (Mail) dan perangkat (Other atau nama yang Anda pilih, seperti "GoMonitor").
- Klik Buat.
- Salin password yang dihasilkan (akan terlihat seperti: abcdefghij123456).

.env
```
EMAIL_ADDRESS=youremail@gmail.com
EMAIL_PASSWORD=abcdefghij123456  # Gantilah ini dengan password khusus aplikasi Anda
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```
