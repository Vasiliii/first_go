package main

import (
	"fmt"
)

func main() {
	varible()
	conditionalOperators()
	cycles()

	fmt.Println("\nEnter some text to exit")
	var a string
	fmt.Scan(&a)
}

func cycles(){
	fmt.Printf("\nThat's result call function of cycles\n\n")

	for i := 0; i<5; i++{
		fmt.Println("That's result cycle for")
	}

	fmt.Printf("\nCycle for by range\n")

	num:= []int{1,2,3,4,5,6}
	for index, elements := range num {
		fmt.Printf("Index: %d Element: %d\n", index, elements)
	}

	fmt.Printf("\nRange for matrix\n")

	matrix := [][]int{{1, 2, 3}, {4, 5, 6},{7, 8, 9}}
	for _,row := range matrix{
		for _,coll := range row{
			fmt.Printf("%d ", coll)
		}
		fmt.Printf("\n")
	}

}

func conditionalOperators() {
	fmt.Printf("\nThat's result call function of conditionalOperators\n")

	var num int
	fmt.Printf("Enter the number\n")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Printf("num is greater than 0\n")
	} else if num < 0 {
		fmt.Printf("num is smaller than 0\n")
	} else {
		fmt.Printf("num = 0\n")
	}
}

func varible() {
	fmt.Println("\nThat's result call function of Variable")

	var age int = 18
	var num float64 = 250.674
	age1 := 16
	var name string

	fmt.Println("\nint var age = ", age)
	fmt.Println("float var name = ", num)
	fmt.Println("age1 := ", age1)
	fmt.Printf("What's your name?\n")
	fmt.Scan(&name)
	fmt.Println("string var name:", name)
	fmt.Println("Len of var name", len(name))
}
