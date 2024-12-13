package main

import (
	"fmt"
)

func computeGrade(gotPoints, maxPoints float64) float64 {
	if maxPoints <= 0 {
		return 1.0
	}
	grade := 1.0 + (5.0 * gotPoints / maxPoints)
	if grade > 6.0 {
		return 6.0
	} else if grade < 1.0 {
		return 1.0
	}
	return grade
}

func main() {
	fmt.Println(computeGrade(17.5, 28.0)) // 4.125
	fmt.Println(computeGrade(28.0, 28.0)) // 6.0
	fmt.Println(computeGrade(0, 28.0))    // 1.0
}
