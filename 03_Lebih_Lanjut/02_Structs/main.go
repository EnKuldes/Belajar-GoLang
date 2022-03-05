package main

import (
	"fmt"
)

// Structs adalah kumpulan dari nilai dan tipe
type Cobakan struct {
	X string
	Y int
}

func main() {
	fmt.Println(Cobakan{"Ku Sayang", 10})

	// Structs field/nilai bisa di akses/edit emnggunakan dot
	x := Cobakan{"Nilai Structs", 16}
	fmt.Println(x.X)
	x.X = "Ku Ganti"
	fmt.Println(x.X)
	fmt.Println(x.Y)

	// Menggunakan Pointer ke Structs
	z := &x
	fmt.Println(z.X)
}