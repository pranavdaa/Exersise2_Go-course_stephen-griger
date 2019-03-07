package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	//this is the madic of go we dont have to put person.firstName and lastname is taken by it self
	//alex := person{"Alex", "Anderson"}
	//method2
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
