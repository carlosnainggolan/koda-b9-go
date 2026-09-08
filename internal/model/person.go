package model

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

func Constructor(name string, address string, phone string) *Person {
	return &Person{
		Name : name,
		Address : address,
		Phone : phone,
	}
}

func (data *Person) PrintData() string {
	return fmt.Sprintf("Name: %s, Address: %s, Phone: %s", data.Name, data.Address, data.Phone)
}

func (data *Person) Greet() string {
	return fmt.Sprintf("Hello %s", data.Name)
}

func (data *Person) Setter(nama string) string {
	data.Name = nama
	return data.Name
}