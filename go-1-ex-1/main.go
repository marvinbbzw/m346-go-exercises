package main

import "fmt"

func main() {
	var name string = "Marvin Eggerschwiler"
	var age int = 16
	var height float64 = 1.70
	var zodiacSign string = "\u2648"

	fmt.Println("Name:", name)
	fmt.Println("Alter:", age)
	fmt.Println("Größe:", height, "m")
	fmt.Println("Sternzeichen:", zodiacSign)
}
