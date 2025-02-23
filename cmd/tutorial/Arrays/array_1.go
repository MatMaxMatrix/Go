package main

import "fmt"

func main() {
	//var intArr [5]int32 = [5]int32{44, 2, 3, 4, 5}
	intArr := [...]int32{44, 2, 3, 4, 5}
	//intArr := [5]int32{44, 2, 3, 4, 5}
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("the length is %d with capacity %d\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("the length is %d with capacity %d\n", len(intSlice), cap(intSlice)) //attention to the capacity which is increased by 3 in this case but append is only adding one element.

	var intSlice2 []int32 = []int32{3, 2}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	fmt.Println("--------------------------------------------------")

	var myMap2 = map[string]uint8{"adam": 20, "John": 30}
	fmt.Println(myMap2)
	fmt.Println(myMap2["adam"])

	var age, ok = myMap2["adam"]

	fmt.Println(age, ok)
	if ok {
		fmt.Printf("adam is %v years old", age)
	} else {
		fmt.Println("adam is not in the map")
	}

	myMap2["adam"] = 30
	fmt.Println(myMap2)

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
		fmt.Printf("Age: %v\n", myMap2[name])
	}

}
