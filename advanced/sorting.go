package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type ByAge []Person
type ByName []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func sorting() {

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}

	fmt.Println(people)

}
