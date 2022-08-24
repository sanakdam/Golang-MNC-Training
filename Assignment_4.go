package main

import (
	"fmt"
)

type Person struct {
	name string
}

func main() {
	var evenNames = func(persons []*Person) {
		for _, person := range persons {
			fmt.Println(person.name)
		}
	}

	names := []*Person{
		{
			name: "Teguh Afdilla",
		},
		{
			name: "Fajar Agus Maulana",
		},
		{
			name: "Kadek Bintang A",
		},
	}

	evenNames(names)
}
