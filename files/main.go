package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content := "This need to go to file"
	file, err := os.Create("./test.txt")
	nilerr(err)
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err) //just shut down the execution and print the error
	}
	fmt.Println(length)
	defer file.Close()
	readFile("./test.txt")
}
func readFile(filePath string) {
	fmt.Println("Reading the file")
	databyte, err := os.ReadFile(filePath)
	nilerr(err)
	fmt.Println("The data inside that file is \n", string(databyte))
}

func nilerr(err error){
	if err != nil {
		panic(err)
	}
}