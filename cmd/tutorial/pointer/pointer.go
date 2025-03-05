package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe address of p is : %v", i)
	*p = 42
	fmt.Printf("\nThe value p points to is: %v", *p)
	p = &i
	*p = 21
	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe address of i is : %v", i)
}