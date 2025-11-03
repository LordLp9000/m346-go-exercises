package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	eyesFile, _ := os.Create("eyes.txt")
	defer eyesFile.Close()
	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	diceFile, _ := os.Create("dice.log")
	defer diceFile.Close()
	fmt.Fprintln(diceFile, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go TODO
}
