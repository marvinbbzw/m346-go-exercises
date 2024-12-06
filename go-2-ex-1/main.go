package main

import "fmt"

type Profile struct {
	FirstName    string
	LastName     string
	DayOfBirth   int
	MonthOfBirth string
	YearOfBirth  int
	Siblings     int
}

func main() {
	me := Profile{
		FirstName:    "Marvin",
		LastName:     "Eggerschwiler",
		DayOfBirth:   15,
		MonthOfBirth: "Juli",
		YearOfBirth:  2008,
		Siblings:     2,
	}

	me.Siblings++

	fmt.Printf("Profile: %+v\n", me)
}
