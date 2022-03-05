package main

import (
	"fmt"
)

type Location struct {
	Lat, Long float64
}

var m map[string]Location

// map literal
var n = map[string]Location{
	"Taraju": Location{22.112,31.1111},
}
// map literal cont
var o = map[string]Location{
	"Sidamulya": {22.112,31.1111},
}

func main() {
	// Map melakukan pemetaan akan keys ke value
	m = make(map[string]Location)

	m["Bogor"] = Location{40.123131, 39.1123123,}
	m["Kuningan"] = Location{40.123131, 39.1123123,}

	fmt.Println(m)
	fmt.Println(m["Kuningan"])

	fmt.Println(n)
	fmt.Println(o)

	// Edit Map
	fmt.Println(m["Bogor"].Lat)
}