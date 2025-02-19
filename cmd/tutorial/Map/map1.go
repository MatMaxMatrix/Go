package main

import "fmt"

func main() {
	var MyMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(MyMap)

	var myMap2 = map[string]uint8{"Adam": 12, "Sarah": 23, "John": 34}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"])
	fmt.Println(myMap2["John"]) //Map returns 0 if the key is not found. This is a default value for uint8 and becareful with this default value.

	delete(myMap2, "John")

	var age, ok = myMap2["John"]
	var age2, ok2 = myMap2["Adam"]
	fmt.Println(age, ok)
	fmt.Println(age2, ok2)

	if ok {
		fmt.Println("John is in the map")
	} else {
		fmt.Println("John is not in the map")
	}

	for name, age := range myMap2 {
		fmt.Printf("%s is %d years old\n", name, age) //The order is not preserved.
	}

	//Similar to a while loop here
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("-----------------")

	//Similar to a for loop and break

	for {
		fmt.Println(i)
		if i == 0 {
			break
		}

		i--
	}
	fmt.Println("-----------------")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
