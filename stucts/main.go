package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

//use this type if you want to access "0 value" of a variable
// func main() {
// 	var mark person

// 	mark.firstName = "Mark"
// 	mark.lastName = "Protsyuk"

// 	fmt.Println(mark)
// 	fmt.Printf("%+v", mark)
// }

// syntax here is not standard but it gives more leaniancy with variables
func main() {
	mark := person{firstName: "Mark", lastName: "Protsyuk"}

	fmt.Println(mark)
}
