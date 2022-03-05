# Module

Pembuatan module sederhana mulai dari cara membuat modul, memakai nya di code lain, handling error dan beberapa lainnya dg sumber utama dari [sini](https://go.dev/doc/tutorial/create-module#prerequisites).

Membuat module, sama dengan sebelumnya yaitu pertama menjalankan command ** go mod init 02_Module_01/greetings ** dimana "02_Module_01" bisa di ganti dg repository source code(apabila di publish di inet) dan "greetings" sbg nama dari module nya.

Pembuatan func berstruktur seperti ini ** func nama_func (nama_parameter tipe_parameter) ** tipe_yang_dikembalikan.
Pada bahasa pemograman Go, untuk func yang huruf awalnya Kapital, maka Func tersebut bisa dipanggil di Code lain atau di Package lain, atau istilah lain di bahasa Go disebut dengan Exported Name.


Membuat simple module yang fungsinya mengucapkan selamat datang pada code yang memanggil fungsi nya.
Struktural dari folder akan seperti ini:

```
02_Module_01
|	greetings
|	| greetings.go
| hello
|	|	hello.go

```

Setelah membuat masing masing file Go di direktorinya masing-masing, pada file "hello.go" di import module greetings dari "02_Module_01/greetings", dikarenakan kita menaruh di lokal, maka perlu ada pengeditan yaitu dengan mengubah file "go.mod" lalu run ``` go mod tidy ``` di folder hello dengan melakukan command:
``` bash
go mod edit -replace 02_Module_01/greetings=../greetings

```