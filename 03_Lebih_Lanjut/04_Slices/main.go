package main

import (
	"fmt"
	"strings"
)

func main() {
	// Slices, mirip dengan array, perbedaan utamanya adalah slices fleksibel untuk jumlah data yang dapat di simpan, tidak perlu di define berapa besarannya
	primes := [6]int{2,3,5,7,11,13}

	// Slices mengambil memory address dari array. SLice dapat dibuat dengan menyediakan dua indeks (indeks rendah dan indeks tertinggi)
	var s []int = primes[0:6]
	fmt.Println(s)

	// slice array primes ke a
	a := primes[0:1]
	fmt.Println(a)
	// ubah value pada variabel a untuk indeks pertama ke 90
	a[0] = 90
	fmt.Println(a)
	// dikarenakan slice mirip seperti pointer ke variabel, maka slice mirip seperti referensi ke array, apabila dilakukan perubahan pada slices maka array yang menjadi sumber utamanya juga akan berubah
	fmt.Println(s)

	// Slice literals (?)
	r := []bool{true, true, true, false}
	fmt.Println(r)
	r[1] = false
	fmt.Println(r)

	// Slice default, secara default biasanya ketika melakukan slicing, indeks terbawah akan secara default 0 dan indeks tertinggi adalah panjang dari array yg akan di slices
	fmt.Println(r[0:4])
	fmt.Println(r[2:])
	fmt.Println(r[:3])
	fmt.Println(r[:])

	// Slice memiliki batasan untuk indeks pertama maupun indeks terakhir, ini bisa di cek dengan menggunakan func len (lenght) dan cap (capacity)
	fmt.Println(r)
	fmt.Printf("len %d - cap %d \n", len(r), cap(r))
	r = r[1:3]
	fmt.Println(r)
	fmt.Printf("len %d - cap %d \n", len(r), cap(r))

	// Membuat slice menggunakna func make
	b := make([]int, 5) // len nya 5
	// b := make([]int, 0, 5) // len nya 0 dan capacity nya 5
	fmt.Println(b)

	// slices di dalam slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][1] = "x"
	board[2][2] = "x"
	printTictacBoard(board)

	// Menambahkan ke slice
	board = append(board, []string{"_", "_", "_"})
	printTictacBoard(board)
	r = append(r, false, true, false, false)
	fmt.Println(r)
}

func printTictacBoard(board [][]string) {
	fmt.Println("Mencetak tictactoe board yeay")
	for i := 0; i < len(board); i++{
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}