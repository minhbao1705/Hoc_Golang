package main

import (
	"fmt"

	"minhbao.com/hoc-golang/cat"
	"minhbao.com/hoc-golang/dog"
	"minhbao.com/hoc-golang/mouse"
	"minhbao.com/hoc-golang/service"
)

func MakeSound(a service.Animal) {
	fmt.Printf("Day la tieng cua %s: \n", a.GetName())
	fmt.Println(a.Speak())
}

func MakeSoundPlus(a service.AnimalPlus) {
	fmt.Printf("Day la tieng cua %s: \n", a.GetName())
	fmt.Println(a.Speak())
	fmt.Println(a.Eat())
}

func PrintValue(val any) {
	// str, ok := val.(string)
	// if ok {
	// 	fmt.Println(str)
	// } else {
	// 	fmt.Println("Please send a string")
	// }

	// numb, ok := val.(int)
	// if ok {
	// 	fmt.Println(numb)
	// } else {
	// 	fmt.Println("Please send a string")
	// }

	switch val.(type) {
	case int:
		fmt.Println(val)
	case string:
		fmt.Println(val)
	default:
		fmt.Println("invalid!")
	}

}

func main() {
	mydog, err := dog.New("  Bully     ")
	if err != nil {
		panic(err)
	}
	MakeSound(mydog)

	PrintValue("-=-=-=-=-=-=-=-=-=-=-")

	mycat, err := cat.New("Pon")
	if err != nil {
		panic(err)
	}
	MakeSoundPlus(mycat)

	PrintValue("-=-=-=-=-=-=-=-=-=-=-")

	mymouse, err := mouse.New("Mr Ti")
	if err != nil {
		panic(err)
	}
	PrintValue(mymouse.GetName())

	PrintValue(3)
	PrintValue(5.8)
	PrintValue(false)
}
