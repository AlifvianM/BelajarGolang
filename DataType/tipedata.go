package main

import "fmt"

func main() {

	// uint8	0 ↔ 255
	// uint16	0 ↔ 65535
	// uint32	0 ↔ 4294967295
	// uint64	0 ↔ 18446744073709551615
	// uint	sama dengan uint32 atau uint64 (tergantung nilai)
	// byte	sama dengan uint8
	// int8	-128 ↔ 127
	// int16	-32768 ↔ 32767
	// int32	-2147483648 ↔ 2147483647
	// int64	-9223372036854775808 ↔ 9223372036854775807
	// int	sama dengan int32 atau int64 (tergantung nilai)
	// rune	sama dengan int32

	// TIPE DATA INTEGER
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// TIPE DATA NUMERIK DESIMAL
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// TIPE DATA BOOLEAN
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// TIPE DATA STRING
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var message2 = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message2)

	// NILAI NILL ATAU ZERO
	// nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya nil berarti memiliki nilai kosong.

	// Semua tipe data yang sudah dibahas di atas memiliki zero value (nilai default tipe data). Artinya meskipun variabel dideklarasikan dengan tanpa nilai awal, tetap akan ada nilai default-nya.

	// Zero value dari string adalah "" (string kosong).
	// Zero value dari bool adalah false.
	// Zero value dari tipe numerik non-desimal adalah 0.
	// Zero value dari tipe numerik desimal adalah 0.0.
	// Zero value berbeda dengan nil. Nil adalah nilai kosong, benar-benar kosong. nil tidak bisa digunakan pada tipe data yang sudah dibahas di atas. Ada beberapa tipe data yang bisa di-set nilainya dengan nil, di antaranya:

	// pointer
	// tipe data fungsi
	// slice
	// map
	// channel
	// interface kosong atau any (yang merupakan alias dari interface{})
	// Nantinya kita akan sering bertemu dengan nil setelah masuk pada pembahasan-pembahasan tersebut.

}
