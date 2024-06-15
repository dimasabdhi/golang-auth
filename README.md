# golang-auth

Proyek ini bertujuan untuk menyediakan fitur login dan sign up dengan keamanan yang baik tanpa menggunakan database. Autentikasi dilakukan menggunakan JSON Web Token (JWT) untuk mengelola sesi pengguna secara aman.

Fitur utama proyek ini meliputi:

Signup: Pengguna dapat mendaftar akun baru dengan menyediakan nama pengguna (username) dan kata sandi (password). Kata sandi yang disimpan dalam bentuk terenkripsi untuk keamanan tambahan.

Login: Pengguna yang sudah mendaftar dapat masuk ke akun mereka dengan memasukkan kredensial yang benar. Sistem akan mengotentikasi pengguna dan memberikan token JWT sebagai tanda bukti sesi yang valid.

JWT Authentication: Setelah login, pengguna akan mendapatkan token JWT yang disimpan dalam cookie untuk otentikasi di endpoint API lainnya. Token ini memastikan bahwa pengguna yang mengakses API memiliki akses yang sesuai dan sesi yang valid.

Keamanan: Proyek ini menggunakan teknik keamanan yang baik, termasuk penggunaan JWT untuk otentikasi dan enkripsi kata sandi pengguna sebelum disimpan.

Tidak Memerlukan Database: Proyek ini dirancang tanpa menggunakan database. Sebagai gantinya, data pengguna disimpan secara sementara dalam memori untuk keperluan autentikasi.
