package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	FirstName string
	LastName  string
	GPA       float64
}

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

func getApplicants(numApplicants int) []Student {
	var students []Student
	students = make([]Student, numApplicants)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < numApplicants; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		fName, lName := line[0], line[1]
		gpa, err := strconv.ParseFloat(line[2], 64)

		if err != nil {
			fmt.Println(err, gpa)
		}
		students[i].FirstName = fName
		students[i].LastName = lName
		students[i].GPA = gpa
	}
	return students
}

func main() {
	var nApplicant, mAccepted int
	fmt.Scan(&nApplicant)
	fmt.Scan(&mAccepted)

	studentList := getApplicants(nApplicant)
	sort.Slice(studentList, func(i, j int) bool {
		if studentList[i].GPA != studentList[j].GPA {
			return studentList[i].GPA > studentList[j].GPA
		}
		return studentList[i].GPA < studentList[j].GPA
	})

	fmt.Println("Successful applicants:")

	for _, student := range studentList[:mAccepted] {
		fmt.Println(student.FirstName, student.LastName)
	}
	//grades := getGrades()
	//average := sum(grades) / float64(len(grades))
	//
	//assessGrade(average)

}
