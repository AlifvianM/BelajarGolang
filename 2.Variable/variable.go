package main

import "fmt"

func main() {
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	// var firstName string = "john"
	// lastName := "wick"
	// fmt.Printf("halo %s %s!\n", firstName, lastName)

	// Diperbolehkan untuk tetap menggunakan keyword var pada saat deklarasi meskipun tanpa menuliskan tipe data,
	// dengan ketentuan tidak menggunakan tanda :=, melainkan tetap menggunakan =
	// lastName := "wick"
	// lastName = "ethan"
	// lastName = "bourne"

	// # Deklarasi Multi Variabel
	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"
	// var fourth, fifth, sixth string = "empat", "lima", "enam"
	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// # Variabel Underscore _
	// _ = "belajar Golang"
	// _ = "Golang itu mudah"
	// name, _ := "john", "wick"
}
