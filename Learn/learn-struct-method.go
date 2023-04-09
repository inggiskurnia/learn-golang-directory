package main

import "fmt"

type Animal struct {
	FirstName string
	LastName  string
	Type      string
	Sex       string
}

func (a Animal) GetFullName() string {
	return a.FirstName + " " + a.LastName
}

func method() {
	var cow Animal
	cow.FirstName = "Slamet"
	cow.LastName = "Sutrisno"

	fmt.Println(cow.GetFullName())
}
