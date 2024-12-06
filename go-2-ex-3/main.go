package main

import "fmt"

func main() {
	modules := map[int]string{
		346: "Cloud-Lösungen konzipieren und realisieren",
		420: "Datenanalyse und Visualisierung",
		101: "Programmieren mit Go",
	}

	fmt.Println("Module 346:", modules[346])
	fmt.Println("Module 420:", modules[420])
	fmt.Println("Module 101:", modules[101])

	delete(modules, 420)
	modules[501] = "Künstliche Intelligenz entwickeln"
	modules[101] = "Programmieren mit Rust"

	fmt.Println("\nUpdated modules:")
	for key, value := range modules {
		fmt.Printf("Module %d: %s\n", key, value)
	}
}
