package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%#X", 42)
}

//prints the number 42 in decimal, binary, and hexidecimal
//%d formats an input number into decimal, %d formats for binary, %x formats for hex
//in line 7, the # denotes I want to preceed the output with "0X" (so I know its hex)
//uppercase X outputs the hex letters in uppercase
