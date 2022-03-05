// Membuat package bernama "greetings"
package greetings

import (
	"fmt"
	// menambahahkan module error untuk error logging
	"errors"
	// Menambhakn modul math untuk mendapatkan random integer
	"math/rand"
)

/* 
Func HelloV1 menerima string Nama dan return String 
*/
func HelloV1(name string) string {
	// Dalam membuat variable, go bisa seperti PHP dimana bisa di declare lgs dg nilai atau bisa dengan declrare berdasarkan tipe
	// Jika ingin melakukan Declare variabel lalu assign nilainya, dengan carai ini
	var pesan string
	pesan = fmt.Sprintf(randomFormat(), name)

	return pesan
}

/* 
Func HelloV2 menerima string Nama dan return String 
*/
func HelloV2(name string) string {
	// Dalam membuat variable, go bisa seperti PHP dimana bisa di declare lgs dg nilai atau bisa dengan declrare berdasarkan tipe

	// Jika lgs declare variabel dengan cara ini:
	pesan := fmt.Sprintf(randomFormat(), name)
	// Return pesan
	return pesan
}

/* 
Func HelloV3 menerima string Nama dan return String atau Error apabila parameter tidak dikirimkan
*/
func HelloV3(name string) (string, error) {
	// Jika parameter name kosong, return error
	if name == "" {
		return "", errors.New("Parameter nama kosong")
	}
	pesan := fmt.Sprintf(randomFormat(), name)
	// line di bawah untuk cek error code atau ga
	// pesan := fmt.Sprintf(randomFormat())
	
	// Return pesan dan nil (artinya ga ada error)
	return pesan, nil
}

/*
Func Hellos menerima kumpulan string dari Nama dan return kumpulan greetings dari masing masing nama yg diterima di parameter
*/
func Hellos(names []string) (map[string]string, error) {
	// bikin variabel nampung value greeting dg masing masing nama yg di terima
	messages := make(map[string]string)
	for _, name := range names {
		message, err := HelloV3(name)
		if err != nil {
		return nil, err 
		}
		messages[name] = message
	}
	return messages, nil
}

/*
Fungsi Random Generate greeting dari predefined greeting
*/
func randomFormat() string {
	// Potongan dari format pesan yang bisa di return
	formats := []string{
		"Hai, %v. Selamat Datang :3 <3",
		"Seneng berjumpa denganmu %v",
		"HELLLLLOOOOOOOOOOOOO, %v",
	}
	// return random slice dari math/rand
	return formats[rand.Intn(len(formats))]
}