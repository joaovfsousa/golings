// structs1
// Make me compile!
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{
		name: "João",
		age:  23,
	}

	fmt.Printf("Person %s and age %d", person.name, person.age)
}
