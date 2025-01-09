package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Active bool
}

func main() {
	user1 := User{"Veer", 22, "@12", true}
	user1.GetUser()
	fmt.Printf("the value of the user is %+v \n", user1)
	user1.GetEmail()
	fmt.Printf("the value of the user is %+v \n", user1)
}
func (u User) GetUser() {
	fmt.Println(u.Active)
}
func (u User) GetEmail() {
	u.Email="setes@12"
	fmt.Println(u.Email)// this will not change the value of it, because u pass as a copy not as a original thing
}
