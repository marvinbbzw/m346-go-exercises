package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceRoll := rand.Intn(6) + 1
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Ausgabe
	fmt.Fprintln(os.Stdout, "Erwürfelte Zahl:", diceRoll)
	fmt.Fprintln(os.Stderr, "Zeitpunkt des Würfelns:", currentTime)
}
