package main

import (
	"fmt"
	"time"
)

// func process(numchain chan int) {
// 	// for num := range numchain {
// 	// 	fmt.Println("processing time", num)
// 	// 	time.Sleep(time.Second)
// 	// }
// 	// fmt.Println("processing time", <-numchain)

// }
// func sum(result chan int, num1 int, num2 int) {
// 	numres := num1 + num2
// 	result <- numres

// }
func EmailSender(emailQueue chan string, done chan bool) {
	defer func() {
		done <- true
	}()
	for email := range emailQueue {
		fmt.Println("email sending to, ", email)
		time.Sleep(time.Second)
	}
}
func main() {
	//queue implement
	emailQueue := make(chan string, 100) //buffer channel, non blocking
	done := make(chan bool)
	go EmailSender(emailQueue,done)
	for i := 0; i < 10; i++ {
		emailQueue <- fmt.Sprintf("%dgmail.com",i)
	}
	fmt.Println("Data added in queue")
	close(emailQueue)
	<-done
	//receving data
	// result := make(chan int)

	// go sum(result, rand.Intn(10), rand.Intn(10))

	// res := <-result
	// fmt.Println(res)
	//sending data in one way

	// num := make(chan int)

	// go process(num)
	// for {
	// 	num <- rand.Intn(100)
	// }

	// time.Sleep(time.Second * 1)

	// error deadlock because blocking
	// mess := make(chan string)
	// mess <- "Hi there"
	// mes := <-mess
	// fmt.Println(mes)
}
