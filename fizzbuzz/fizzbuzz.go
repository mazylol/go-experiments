package fizzbuzz

import "fmt"

func Fizzbuzz() {
	for i := 0; i < 100; i++ {
		var out = ""

		if i%3 == 0 {
			out += "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}

		if out == "" {
			fmt.Println(i)
		} else {
			fmt.Println(out)
		}
	}
}
