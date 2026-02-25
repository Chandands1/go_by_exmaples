package main

import "fmt"

type person struct{
	name string
	age int
}

func newPerson(name string) *person{
	p := person{name: name}
	p.age = 42
	return &p
}

func main(){
	fmt.Println(person{"Bob",20})
	fmt.Println(person{name:"Alice",age: 50})
	fmt.Println(person{name:"Raj"})

	fmt.Println(&person{name: "ann", age: 40})
	fmt.Println(newPerson("jon"))

	dog := struct{
		name string
		isGood bool
	}{
		"res",
		true,
	}
	fmt.Println(dog)
}