package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	ceil string
}

func main() {

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "london",
			country: "UK",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "021-462156",
			ceil: "9845663232",
		},
	}

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	p3 := Person{
		firstName: "Jane",
		age:       25,
	}

	// p2.address.city = "New York"
	// p2.address.country = "USA"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.address)
	fmt.Println(p1.address.country)
	fmt.Println(p2.address.country)
	fmt.Println(p1.ceil)
	fmt.Println(p1.home)
	fmt.Println("Are p1 and p2 equall:", p1 == p2)
	fmt.Println("Are p3 and p2 equall:", p3 == p2)

	// Anonymous Structs

	user := struct {
		username string
		email    string
	}{
		username: "user124",
		email:    "pseudoemail@example.com",
	}
	fmt.Println(user.username)

	fmt.Println(p1.fullName())

	fmt.Println("Before", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After", p1.age)

}
