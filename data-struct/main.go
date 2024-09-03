package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
}

type Category struct {
	name string
	id   int
}

type Product struct {
	name     string
	quantity int16
	category Category
}

func main() {
	person1 := Person{"Tuan", 18, "dennis@gmail.com"}
	person2 := Person{name: "Vi", age: 100, email: "vi@gmail.com"}
	var person3 Person

	person3.age = 29
	person3.email = "tuanprodd@gmail.com"
	person3.name = "Khanh"

	fmt.Println(person1)
	fmt.Println(person2.age, person2.name, person2.email)
	fmt.Printf("%+v", person3)

	//
	newProduct := Product{
		name:     "IPAD",
		quantity: 2024,
		category: Category{name: "Smart Phone", id: 1},
	}
	fmt.Printf("%+v\n", person3)

	// newProductPointer := &newProduct

	// newProductPointer.updateQuantity(40)
	newProduct.updateQuantity(40)
	newProduct.print()
}

func (pointerP *Product) updateQuantity(value int16) {
	(*pointerP).quantity = value
}

func (p Product) print() {
	fmt.Printf("%+v", p)
}
