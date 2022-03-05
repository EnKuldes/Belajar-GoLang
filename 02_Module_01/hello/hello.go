package main

import (
	"fmt"
	// Import Module greetings yg di buat
	"02_Module_01/greetings"

	// Import Module Log
	"log"
)

func main() {
	// Case 1
	message := greetings.HelloV1("Farhan")
	fmt.Println(message)

	// Case 2
	message = greetings.HelloV2("Muzaki")
	fmt.Println(message)

	// Mempersiapkans Log
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Case 3
	newMessage, err := greetings.HelloV3("Farhan Muzaki")
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println(newMessage)

	// Case 4
	/*newMessage, err = greetings.HelloV3("")
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println(newMessage)*/

	// Case 5
	names := []string {
		"Farhan",
		"Almas",
		"Hafiz",
		"Muzaki",
	}

	newMessages, errs := greetings.Hellos(names)
	if errs != nil {
		log.Fatal(errs) 
	}
	fmt.Println(newMessages)
}