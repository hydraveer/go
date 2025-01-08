package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time study in Golang")
	start := time.Now()
	fmt.Println(start)
	fmt.Println(start.Format("01-02-2006 15:04:05 Monday"))
}
