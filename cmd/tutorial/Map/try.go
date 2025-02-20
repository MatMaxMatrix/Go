package main

import (
	"fmt"
	"strings"
)

func main() {
	Try()
}

func Try() {
	var myString = []rune("Résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\n The length of the string is %v\n", len(myString))

	var myRune = 'a'
	fmt.Printf("%v, %T\n", myRune, myRune)

	//Following is the old way of concatenating strings.
	var strSlice = []string{"a", "b", "c", "d", "e", "f"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i] //We are creating a new string everytime and is not efficient because strings are immutable in go and can't modify.
	}
	fmt.Println(catStr)
	//catStr[0] = "z" //This will throw an error because strings are immutable.

	//Following is the new way of concatenating strings. Much faster and efficient.
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Println(catStr2)
}
