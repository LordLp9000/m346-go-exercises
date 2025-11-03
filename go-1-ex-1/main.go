package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	firstName := "Laurent"
	lastName := "Scherrer"
	dayOfBirth := 9
	monthOfBirth := 4
	yearOfBirth := 2005
	numberOfSiblings := 1
	heightInMeters := 1.72
	zodiacSign := '♉'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
