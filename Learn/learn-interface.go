package main

import "fmt"

type HasName interface {
	GetName() string
}

func SaysHello(hasName HasName) {
	fmt.Println("Hallo my name is ", hasName.GetName())
}

// FIrst struct
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

//Second struct
type Building struct {
	Name     string
	Location string
}

func (building Building) GetName() string {
	return building.Name
}
func main() {
	var inggis Person
	inggis.Name = "Inggis Kurnia"

	var beltway Building
	beltway.Name = "office"

	SaysHello(inggis)
	SaysHello(beltway)
}
