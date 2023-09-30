package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street  string
	City    string
	ZipCode string
}

func (p *Person) UpdateName(newName string) {
	p.Name = newName
}

func main() {

	person := Person{
		Name: "John",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "10001",
		},
	}

	fmt.Println("Data awal:")
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Street:", person.Address.Street)
	fmt.Println("City:", person.Address.City)
	fmt.Println("ZipCode:", person.Address.ZipCode)

	person.UpdateName("Alice")

	fmt.Println("\nData setelah manipulasi:")
	fmt.Println("Name:", person.Name)
}
