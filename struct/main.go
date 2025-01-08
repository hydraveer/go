package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Active bool
}

func main() {
	Mohit := User{"Veer",22,"Veer@12",true}
	fmt.Println(Mohit)
	fmt.Printf("The details of mohit is %+v\n",Mohit)
}
