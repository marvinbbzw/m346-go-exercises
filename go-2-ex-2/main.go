package main

import "fmt"

func main() {
	fibs := []int{1, 1}

	for i := 2; i < 5; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])
	}

	for i := 5; i < 9; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])
	}

	fmt.Println("Fibonacci sequenz:", fibs)
}
