package main

import (
	"fmt"
)

// Following code is to show that the memory location of the array is different for the original and the function which uses more memory

// func main() {
// 	var thing1 = [5]float64{1,2,3,4,5}
// 	fmt.Printf("thing1: %v", thing1)
// 	var result = square(thing1)
// 	fmt.Printf("\nresult: %v", result)
// 	fmt.Printf("\nThe memory location of thing1 is: %p", &thing1)
// }

// func square(thing2  [5]float64) [5]float64 {
// 	fmt.Printf("\nThe memory location of thing2 is: %p", &thing2)
// 	for i := range thing2 {
// 		thing2[i] = thing2[i] * thing2[i]
// 	}
// 	return thing2
// 	}

// Following code is to show using pointers to change the value of the original variable and save memory


func main() {
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("The memory location of thing1 is: %p", &thing1)
	fmt.Printf("\nthing1: %v", thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nresult: %v", result)
	fmt.Printf("\nThe memory location of thing1 is: %p", &thing1)
}


func square(thing2  *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location that thing2 points to is: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
	}
