package main

import (
	"fmt"
)

func main() {
	// Pointer
	// set variabel
	i := 97

	x := &i // bikin variabel x nge point ke variabel i
	fmt.Println(*x) // untuk mencetak value dari variabel i menggunakan pointer x
	*x = 1997 // untuk set value via pointer
	fmt.Println(*x)
	*x = *x * 2 // modify value nya dg melakukan perkalian
	fmt.Println(*x)

	// set variabel
	var o int = 10
	var y *int
	
	y = &o // bikin variabel y nge point ke variabel o
	fmt.Println(*y)
	y = &i // bikin variabel y nge point ke variabel i
	fmt.Println(*y)

	fmt.Println(y) // menunjukan memory address dari variabel yg di point kan

}