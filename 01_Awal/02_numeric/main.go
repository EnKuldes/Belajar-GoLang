package main

import (
	"fmt"
)

// Mencetak angka seperti desimal, biner dan hexadesimal

func main() {
	// Variable penampung
	var angka = 97
	// mencetak angka desimal
	fmt.Println(angka)

	/* mencetak angka biner
	printf artinya mencetak dg format, dimana %d artinya untuk format desimal, %b untuk format biner
	*/
	fmt.Printf("%d - %b \n", angka, angka)

	/* mencetak angka hexadesimal
	%#X untuk heksadesimal
	*/
	fmt.Printf("%d - %#X \n", angka, angka, angka)
}