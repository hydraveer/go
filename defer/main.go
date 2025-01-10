package main

import "fmt"

func main() {
	// defer fmt.Println("world") // defer put this into last line if we put defer before anything then it will push it into last line
	// fmt.Println("Hello")
	// fmt.Println("Make")
	// fmt.Println("Laugh")
	defer fmt.Println("!")
	defer fmt.Println("world")
	defer fmt.Println("hello")
	fmt.Println("This work like last in first out way")
}
