package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "Hello World"
	printMe(printValue)
	numerator := 10
	denominator := 0
	result, remainder, err := intDivision(numerator, denominator)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("%d divided by %d is %d\n", numerator, denominator, result)
	default:
		fmt.Printf("%d divided by %d is %d with a remainder of %d\n", numerator, denominator, result, remainder)
	}
	switch remainder {
	case 0:
		fmt.Printf("Devision was successful\n")
	case 1, 2:
		fmt.Printf("devision was close\n")
	default:
		fmt.Printf("devision was not successful\n")
	}
}

// 	if err != nil {
// 		fmt.Printf(err.Error)
// 	}else if  remainder == 0 {
// 		fmt.Printf("%d divided by %d is %d\n", numerator, denominator, result)
// 	}else {
// 	fmt.Printf("%d divided by %d is %d with a remainder of %d\n", numerator, denominator, result, remainder)
// }

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}
