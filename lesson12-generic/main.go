package main

import (
	"cmp"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Box[T any] struct {
	Content     T
	Description T
}

func PrintValue[T any](value T) {
	fmt.Println(value)
}

func IsEqual[T comparable](a, b T) bool {
	return a == b
}

func IsNotEqual[T comparable](a, b T) bool {
	return a != b
}

type Number interface {
	constraints.Integer | constraints.Float
}

func sum[T Number] (a, b T) T {
	return a + b
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MaxLengthString(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}

func main() {
	// fmt.Println(PrintValue("Bao"))
	// fmt.Println(PrintValue(10))
	// fmt.Println(PrintValue(true))

	// result01 := IsEqual(3, 5)
	// PrintValue(result01)

	// result02 := IsEqual("Bao", "Bao")
	// PrintValue(result02)

	// result03 := IsEqual(3.5, 9.1)
	// PrintValue(result03)

	// result04 := IsNotEqual(3.5, 9.1)
	// PrintValue(result04)

	// 	PrintValue(max(5, 7))
	// 	PrintValue(max(5.5, 9.2))
	// PrintValue(MaxLengthString("Bao", "Minh Bao"))

	// stringBox := Box[string]{Content: "Hoc golang Generic", Description: "Mo ta String box"}
	// intBox := Box[int]{Content: 99, Description: 99}
	// PrintValue(stringBox.Content)
	// PrintValue(intBox.Content)
	// PrintValue(stringBox.Description)
	// PrintValue(intBox.Description)

	PrintValue(sum(5.5, 10))
	PrintValue(sum(9, 10))
	PrintValue(sum(3.6, 6.4))
}
