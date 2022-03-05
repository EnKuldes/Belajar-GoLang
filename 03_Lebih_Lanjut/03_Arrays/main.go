package main
import (
	"fmt"
	"math/rand"
)

func main() {
	// dalam pembuatan Array harus di definisikan besar array nya
	var a[5] int
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
}