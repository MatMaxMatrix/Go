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

}
