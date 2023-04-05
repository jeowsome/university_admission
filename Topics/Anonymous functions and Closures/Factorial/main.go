package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	factorial := func() int {
		f := 1
		for i := 1; i <= number; i++ {
			f *= i
		}
		return f
	}()

	fmt.Println(factorial)
}
