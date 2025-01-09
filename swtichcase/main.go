package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1
	switch dice{
		case 1:
			fmt.Println("less then 3")
		case 2:
			fmt.Println("greater then 1")
		case 3: 
			fmt.Println("less then 4")
			fallthrough
		case 4:
			fmt.Println("greater then 3") 
		case 5:
			fmt.Println("less then 6")
		case 6: 
		fmt.Println("greater then 5")
	}
}
