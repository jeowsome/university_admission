package main

import "fmt"

const (
	Min = 0
	Max = 100
)

func main() {
	constraint := newConstraint(Min, Max)

	var number int
	fmt.Scan(&number)

	fmt.Println(constraint(number))
}

func newConstraint(a, b int) func(num int) int {
	minValue := a
	maxValue := b

	return func(num int) int {
		if num < minValue {
			return minValue
		}
		if num > maxValue {
			return maxValue
		}
		return num
	}
}
