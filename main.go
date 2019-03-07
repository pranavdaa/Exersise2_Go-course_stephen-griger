package main

import "fmt"

//embading one struct to another
type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	//this is the madic of go we dont have to put person.firstName and lastname is taken by it self
	//alex := person{"Alex", "Anderson"}
	//method2
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	//Magic of go -> if you forgot to assign a valiie to one of the fields then go automatically gives a zero value(eprty string to it)
	//	this will even give you the field anmes also
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9094, //special ,
		}, //, specil when there embaded structs
	}
	//Discriptive pritning as defined above
	fmt.Printf("%v", jim)
}
