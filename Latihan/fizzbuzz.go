package main

import "fmt"

func fizzbuzz(count int) {
	for i := 1; i < count; i++ {
		//print fizzbuzz
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("FizzBuzz =", i)
		}

		if i%3 == 0 {
			fmt.Println("Fizz =", i)
		}

		if i%5 == 0 {
			fmt.Println("Buzz =", i)
		} else {
			fmt.Println(i)
		}

	}
}

func main() {
	fizzbuzz(30)
}
