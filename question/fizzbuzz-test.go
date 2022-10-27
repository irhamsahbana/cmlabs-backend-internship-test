package question

import "fmt"

func FizzBuzz(from int, to int) {

	for i := from; i <= to; i++ {
		msg := ""

		if i%3 == 0 {
			msg = fmt.Sprintf("%d: %s", i, "Fizz")
		}

		if i%5 == 0 {
			msg = fmt.Sprintf("%d: %s", i, "Buzz")
		}

		if i%3 == 0 && i%5 == 0 {
			msg = fmt.Sprintf("%d: %s", i, "FizzBuzz")
		} else if i%3 != 0 && i%5 != 0 {
			msg = fmt.Sprintf("%d: %d", i, i)
		}

		println(msg)
	}
}
