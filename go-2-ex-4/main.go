package main

import "fmt"

// TODO: declare a type for Student (with first and last name)
type Student struct {
	FirstName string
	LastName  string
}

// TODO: declare a type for Class (consisting of multiple students)
type Class struct {
	Name     string
	Students []Student
}

func main() {
	// TODO: declare a type for Student (with first and last name)
	// TODO: declare a type for Class (consisting of multiple students)
	// TODO: declare a map of modules being attended by multiple classes

	// Create students
	classA := Class{
		Name: "Class A",
		Students: []Student{
			{FirstName: "Anna", LastName: "Müller"},
			{FirstName: "Bob", LastName: "Schmidt"},
			{FirstName: "Clara", LastName: "Weber"},
		},
	}

	classB := Class{
		Name: "Class B",
		Students: []Student{
			{FirstName: "David", LastName: "Fischer"},
			{FirstName: "Emma", LastName: "Wagner"},
			{FirstName: "Felix", LastName: "Becker"},
		},
	}

	// Create map of modules with classes attending them
	modules := map[int][]Class{
		104: {classA, classB},
		117: {classA},
		346: {classB, classA},
	}

	// TODO: output everything using fmt.Println()
	fmt.Println("=== Class Management System ===")
	fmt.Println("\nClasses:")
	fmt.Println(classA.Name, ":", classA.Students)
	fmt.Println(classB.Name, ":", classB.Students)

	fmt.Println("\nModules:")
	for moduleNum, classes := range modules {
		fmt.Printf("Module %d attended by: ", moduleNum)
		for i, class := range classes {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(class.Name)
		}
		fmt.Println()
	}

	fmt.Println("\nComplete module data:")
	fmt.Println(modules)
}
