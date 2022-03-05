package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		[]int{10,1,1997},
		[]int{16,3,1996},
	}
	fmt.Println(a)
	
	// range digunakan pada for untuk melakukan iterasi pada slice atau map
	for i, v := range a{
		fmt.Printf("%d = %d\n", i, v)
	}

	// looping dg range bisa dilakukan dg cara
	// for i, _ := range a // digunakan untuk skip value nya saja
	// for _, value := range a // digunakan untuk skip indeks nya saja
	// for i := range a // digunakan untuk ngambil indeks nya saja
}