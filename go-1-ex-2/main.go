package main

import "fmt"

func main() {
	var miles float64 = 10
	var fahrenheit float64 = 98.6

	kilometers := miles * 1.60934
	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("%.2f Meilen sind %.2f Kilometer\n", miles, kilometers)
	fmt.Printf("%.2f Fahrenheit sind %.2f Celsius\n", fahrenheit, celsius)

	var km float64 = 16.0934
	var c float64 = 37

	milesFromKm := km / 1.60934
	fahrenheitFromC := c*9/5 + 32

	fmt.Printf("%.2f Kilometer sind %.2f Meilen\n", km, milesFromKm)
	fmt.Printf("%.2f Celsius sind %.2f Fahrenheit\n", c, fahrenheitFromC)
}
