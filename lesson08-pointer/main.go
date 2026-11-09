package main

import "fmt"

func lythuyetpointer() {
	name := "Vo Minh Bao"
	fmt.Println("-=-=-=-=-=- Information name variable -=-=-=-=-=-")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)

	fmt.Println()
	// Tao con tro (pointer)
	ptrName := &name
	fmt.Println("-=-=-=-=-=- Information ptrName variable -=-=-=-=-=-")
	fmt.Printf("Data type: %T \n", ptrName)
	fmt.Printf("Value: %v \n", ptrName)
	fmt.Printf("Address: %v \n", &ptrName)

	fmt.Printf("Find value name from ptrName: %v \n", *ptrName)

	fmt.Println()

	// Tao con tro
	ptrName2 := &ptrName
	fmt.Println("-=-=-=-=-=- Information ptrName2 variable -=-=-=-=-=-")
	fmt.Printf("Data type: %T \n", ptrName2)
	fmt.Printf("Value: %v \n", ptrName2)
	fmt.Printf("Address: %v \n", &ptrName2)

	fmt.Printf("Find ptrName1 name from ptrName2: %v \n", *ptrName2)
	fmt.Printf("Find value name from ptrName2: %v \n", **ptrName2)
}

func updateName(name *string) {
	*name = "Nguyen Van A"
	fmt.Println("-=-=-=-=-=- Information name variable inside updateName func -=-=-=-=-=-")
	fmt.Printf("Value: %v \n", *name)
	fmt.Printf("Address: %v \n", name)
}

func main() {
	name := "Vo Minh Bao"
	fmt.Println("-=-=-=-=-=- Information name variable -=-=-=-=-=-")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)

	fmt.Println()
	name = "Ara Vo"
	fmt.Println("-=-=-=-=-=- Information name variable after update value -=-=-=-=-=-")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)

	fmt.Println()

	updateName(&name)

	fmt.Println()

	fmt.Println("-=-=-=-=-=- Information name variable after run updateName -=-=-=-=-=-")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)
}
