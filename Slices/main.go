package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Welcome in slices")
	// var flist = []string{}
	// flist = append(flist, "Green","Red", "white","blue","purple")
	// fmt.Println(flist)
	// flist = append(flist[1:])// give all the value that comes after 1st index
	// fmt.Println(flist)
	// flist = append(flist[:1]) // give all the value that comes before 1st index
	// fmt.Println(flist)

	// scores := make([]int, 4)
	// scores[0]=10
	// scores[1]=20
	// scores[2]=30000
	// scores[3]=40
	// // scores[4]=33 this will through error index out of bound
	// scores = append(scores, 333, 3333)
	// fmt.Println(scores)
	// sort.Ints(scores)
	// fmt.Println(scores)
	// fmt.Println(sort.IntsAreSorted(scores))

	// remove slices
	var st = []string{"js", "react", "swift", "rust", "go"}
	fmt.Println(st)
	var index int = 2
	st = append(st[:index], st[index+1:]...);
	fmt.Println(st)
	

}
