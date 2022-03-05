package main

import (
	"fmt"
)

var namaku string = "Farhan" // Ini Scope nya satu package

const (
	pi = 3.14
	bahasa = "Indonesia"
)

func main() {
	// Variabel via Short Hand
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	// %T untuk menunjukan tipe variabel
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)

	// Variabel via agak ribet
	var a1 = "this is stored in the variable a1"
	var b1, c1 string = "stored in b1", "stored in c1"
	var d1 string
	d1 = "stored in d1"

	fmt.Printf("%v \n", a1)
	fmt.Printf("%v \n", b1)
	fmt.Printf("%v \n", c1)
	fmt.Printf("%v \n", d1)

	fmt.Printf("%v \n", namaku)
	namaku = "Muzaki"
	fmt.Printf("%v \n", namaku)

	const kata = "Pemograman"
	fmt.Printf("%v %v \n", kata, bahasa)
	fmt.Printf("%f \n", pi)
}