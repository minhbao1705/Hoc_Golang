package main

import "fmt"

func main() {
	// fmt.Print("Hello Bao.")
	// fmt.Print("Goodbye Bao.")

	// fmt.Println("Hello Bao.")
	// fmt.Println("Goodbye Bao.")

	// var fullName = "Vo Minh Bao"
	// var age = 24
	// fmt.Printf("Hello %s and your age is %d. \n", fullName, age)

	// var firstName, lassName string
	// fmt.Print("Please enter your name: ")
	// fmt.Scan(&firstName, &lassName)
	// fmt.Println("My name is: ", firstName, lassName)

	// var firstName, lassName string
	// fmt.Print("Please enter your first name: ")
	// fmt.Scanln(&firstName)
	// fmt.Print("Please enter your last name: ")
	// fmt.Scanln(&lassName)
	// fmt.Println("My name is: ", firstName, lassName)

	// var name string
	// var age int
	// fmt.Println("Please enter your name: ")
	// fmt.Scanf("%s", &name)
	// fmt.Println("Please enter your age: ")
	// fmt.Scanf("%d", &age)
	// fmt.Printf("My name is %s and age is %d \n", name, age)

	// message := fmt.Sprint("My name is ", "Bao")
	// fmt.Println(message)

	// message := fmt.Sprintln("My name is ", "Bao")
	// fmt.Println(message)

	// name := "Bao"
	// age := 24
	// message := fmt.Sprintf("My name is %s and age is %d", name, age)
	// fmt.Println(message)

	ten := "Vo Minh Bao"
	tuoi := 24
	chieucao := 1.73
	daTotNghiep := true
	phanTram := 10

	fmt.Printf("Kieu du lieu cua bien ten %T \n", ten)
	fmt.Printf("Kieu du lieu cua bien ten %T \n", tuoi)
	fmt.Printf("Kieu du lieu cua bien ten %T \n", chieucao)
	fmt.Printf("Kieu du lieu cua bien ten %T \n", daTotNghiep)
	fmt.Printf("Kieu du lieu cua bien ten %T \n", phanTram)

	fmt.Printf("Toi ten la: %v \n", ten)

	fmt.Printf("Toi ten la %s va toi %d \n", ten, tuoi)

	fmt.Printf("Chieu cao cua toi la %.2f \n", chieucao)
	fmt.Printf("Chieu cao cua toi la %.5f \n", chieucao)
	fmt.Printf("Chieu cao cua toi la %.1f \n", chieucao)

	fmt.Printf("Toi da tot nghiep: %t \n", daTotNghiep)
	
	fmt.Printf("Toi da hoc duoc %d%% \n", phanTram)
}
