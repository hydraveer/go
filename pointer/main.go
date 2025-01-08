package main

import "fmt"

func main() {
	fmt.Println("Welcome to study Pointers")
	// var x *int;
	// fmt.Println("The value of this pointer", x)
	myNumber := 30
	numPointer := &myNumber
	fmt.Println("The value of pointer ", numPointer)
	fmt.Println("The value of pointer ", *numPointer)
	*numPointer = *numPointer *2
	fmt.Println(myNumber)

}
