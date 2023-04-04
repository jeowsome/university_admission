package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const defaultRange = 3

func getGrades() []float64 {
	var grades []float64
	grades = make([]float64, defaultRange)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < defaultRange; i++ {
		scanner.Scan()
		line := scanner.Text()
		f, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println(err, f)
		}

		grades[i] = f
	}
	return grades
}

func sum(grades []float64) float64 {
	var sum float64

	for _, grade := range grades {
		sum += grade
	}

	return sum
}

func assessGrade(average float64) {
	fmt.Println(average)
	if average >= 60 {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}
}

func main() {
	grades := getGrades()
	average := sum(grades) / float64(len(grades))

	assessGrade(average)

}
