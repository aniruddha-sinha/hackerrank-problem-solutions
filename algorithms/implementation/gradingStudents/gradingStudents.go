package gradingstudents

import (
	"fmt"
)

func processPassingGrade(number int32) int32 {
	var finalNumber int32
	for i := number; i < (number + 10); i++ {
		if i%5 == 0 && i-number < 3 {
			finalNumber = i
			break
		} else {
			finalNumber = number
		}
	}

	return finalNumber
}

func evalGrade(grade int32) int32 {
	var finalGrade int32
	if grade >= 38 {
		finalGrade = processPassingGrade(grade)
	} else {
		finalGrade = grade
	}

	return finalGrade
}

func gradingStudents(grades []int32) []int32 {
	var resultantGrades []int32 = make([]int32, 0)
	for _, val := range grades {
		resultantGrades = append(resultantGrades, evalGrade(val))
	}

	return resultantGrades
}

func GradingStudentsHelper(grades []int32) {
	roundedOffGrades := gradingStudents(grades)
	for _, g := range roundedOffGrades {
		fmt.Println(g)
	}
}
