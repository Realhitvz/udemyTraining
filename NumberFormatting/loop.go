package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}

//line 6: "for" creates a loop.
//the variable "i" is instructed to begin at 0
//to run the looped code if it is less then 200
//and to increase its value by 1 after running the looped code
//this prints the numbers 0 to 199 in decimal, binary, and hex
//%q could be used to show UTF-8 letter values for these numbers
