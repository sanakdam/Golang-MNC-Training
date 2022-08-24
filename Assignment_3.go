package main

import (
	"fmt"
)

func main() {
	var evenNames = func(names []*string) []string {
		var result []string

		for _, name := range names {
			result = append(result, *name)
		}

		return result
	}

	teguh := "Teguh Afdilla"
	fajar := "Fajar Agus Maulana"
	kadek := "Kadek Bintang A"

	names := []*string{
		&teguh,
		&fajar,
		&kadek,
	}

	fmt.Println(evenNames(names))
}
