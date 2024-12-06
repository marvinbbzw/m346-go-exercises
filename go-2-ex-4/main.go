package main

import "fmt"

type Student struct {
	ErstName string
	Nachname string
}

type Class struct {
	Name     string
	Students []Student
}

type Module struct {
	ID      int
	Classes []Class
}

func main() {
	class1 := Class{
		Name: "Class A",
		Students: []Student{
			{ErstName: "Alice", Nachname: "Smith"},
			{ErstName: "Bob", Nachname: "Johnson"},
			{ErstName: "Charlie", Nachname: "Brown"},
		},
	}

	class2 := Class{
		Name: "Class B",
		Students: []Student{
			{ErstName: "Diana", Nachname: "Miller"},
			{ErstName: "Eve", Nachname: "Davis"},
			{ErstName: "Frank", Nachname: "Wilson"},
		},
	}

	module1 := Module{
		ID:      101,
		Classes: []Class{class1},
	}

	module2 := Module{
		ID:      202,
		Classes: []Class{class2},
	}

	module3 := Module{
		ID:      303,
		Classes: []Class{class1, class2},
	}

	fmt.Println("Modules and Classes:")
	modules := []Module{module1, module2, module3}
	for _, module := range modules {
		fmt.Printf("Module %d:\n", module.ID)
		for _, class := range module.Classes {
			fmt.Printf("  %s:\n", class.Name)
			for _, student := range class.Students {
				fmt.Printf("    - %s %s\n", student.ErstName, student.Nachname)
			}
		}
	}
}
