package main

import "fmt"

func main() {
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	var point_2 = 8840.0

	if percent := point_2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var point_3 = 6

	switch point_3 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var point_4 = 6

	switch point_4 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var point_5 = 6

	switch point_5 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	var point_6 = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point_6 < 8) && (point_6 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// Penggunaan Keyword fallthrough Dalam switch
	// Seperti yang sudah dijelaskan sebelumnya, bahwa switch pada Go memiliki perbedaan dengan bahasa lain.
	// Ketika sebuah case terpenuhi, pengecekan kondisi tidak akan diteruskan ke case-case setelahnya.
	// Keyword fallthrough digunakan untuk memaksa proses pengecekan diteruskan ke satu case selanjutnya dengan
	// tanpa menghiraukan nilai kondisinya, jadi satu case di pengecekan selanjutnya tersebut selalu dianggap
	// benar (meskipun aslinya adalah salah). Dalam sebuah switch lebih dari satu fallthrough bisa di tempatkan
	// untuk memaksa melanjutkan proses pengecekan ke satu case setelahnya.

	var point_7 = 6

	switch {
	case point_7 == 8:
		fmt.Println("perfect")
	case (point_7 < 8) && (point_7 > 3):
		fmt.Println("awesome")
		fallthrough
	case point_7 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
