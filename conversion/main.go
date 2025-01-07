package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome the Zudio")
	fmt.Println("Please rate us 1 to 5")
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your rating")
	input,_ := reader.ReadString('\n');
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input),64);
	if err != nil{
		fmt.Println("There is some error",err)
	}
	fmt.Println(numrating+1)
}
