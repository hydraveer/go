package main

import "fmt"

func main() {
	// names := []string{"Veer", "mohit", "Sandy", "Sanjay"}
	// // for i := 0; i < len(name); i++ {
	// // 	fmt.Println(name[i])
	// // }
	// // for i := range name{
	// // 	fmt.Println(name[i])
	// // }
	// for i, name := range names{
	// 	fmt.Printf("The name is %v and the index is %v \n", name, i)
	// }
	value := 1;
	for value<10{
		if value==3{
			goto l
		}
		if value==5{
			value++
			continue
		}
		fmt.Println(value)
		value++
	}
	l:
		fmt.Println("goto hit")
}
