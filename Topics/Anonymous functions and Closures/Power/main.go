package main

import (
	"fmt"
)

func main() {
	var base, exponent int
	fmt.Scan(&base, &exponent)

	pow := func(n, e int) int {
		result := 1
		for e > 0 {
			result *= n
			e--
		}
		return result
	}

	fmt.Println(pow(base, exponent))
}
