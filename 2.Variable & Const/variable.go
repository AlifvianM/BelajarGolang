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

	// CONST
	const firstName_2 string = "john"
	fmt.Print("halo ", firstName_2, "!\n")

	const lastName_2 = "wick"
	fmt.Print("nice to meet you ", lastName_2, "!\n")

	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	//Contoh deklarasi konstanta dengan tipe data dan nilai yang sama:
	const (
		a = "konstanta"
		b
	)

	const (
		today string = "senin"
		sekarang
		isToday2 = true
	)
	/*
	*today dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "senin"
	*sekarang dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "senin"
	*isToday2 dideklarasikan dengan metode type inference dengan tipe data bool dan nilainya true
	 */

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
}
