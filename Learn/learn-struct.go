package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Birth string
}

func main() {
	var person1 Person
	person1.Name = "inggis"
	person1.Age = 42
	person1.Birth = "Magelang"

	fmt.Println(person1)
}
