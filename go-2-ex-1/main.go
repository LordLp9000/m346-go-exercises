package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	Name             FullName
	Birth            BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		Name: FullName{
			FirstName: "Laurent",
			LastName:  "Scherrer",
		},
		Birth: BirthDate{
			Day:   9,
			Month: 4,
			Year:  2005,
		},
		NumberOfSiblings: 2,   // TODO: adjust
		ZodiacSign:       '♈', // TODO: adjust (Aries for April 9)
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
