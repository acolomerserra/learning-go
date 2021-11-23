package main

import (
	"fmt"
	"strconv"
)

var global_i int = 124
var global_j float32 = 124

// var firstName string = "Adrià"
// var lastName string = "Colomer Serra"
// var age int = 36

//Example group vars
var (
	firstName string = "Adrià"
	lastName  string = "Colomer Serra"
	age       int    = 36
)

//Can declare multiple group vars
var (
	counter int = 1
	i       int = 24 // Global scope var only visible inside package
	J       int = 30 //Package scope but can used it everywhere of the app - Uppercase first letter make it visible to others
)

func main() {

	var i int // Function scope var with the same name as a global var pre-declared = Shadowing
	i = 42
	fmt.Println(i)
	i = 27
	fmt.Println(i)

	var j float32 = 53
	fmt.Println(j)
	fmt.Printf("%v, %T", j, j)
	fmt.Println("")

	k := 87
	fmt.Println(k)
	fmt.Printf("%v, %T", k, k)
	fmt.Println("")

	l := 98
	fmt.Printf("%v, %T", l, l)
	fmt.Println("")

	m := 98.
	fmt.Printf("%v, %T", m, m)
	fmt.Println("")

	global_i = 121
	fmt.Printf("%v, %T", global_i, global_i)
	fmt.Println("")

	global_j = 121
	fmt.Printf("%v, %T", global_j, global_j)
	fmt.Println("")

	var theURL string = "https://google.com" // Keep the acronyms in uppercase
	fmt.Println(theURL)

	//Convert vars from one type to other
	var firstNumber int = 20
	var secondNumber float32
	var thirdString string

	secondNumber = float32(firstNumber)
	fmt.Printf("%v, %T", firstNumber, firstNumber)
	fmt.Println("")
	fmt.Printf("%v, %T", secondNumber, secondNumber)
	fmt.Println("")

	thirdString = string(firstNumber)
	fmt.Printf("%v, %T", thirdString, thirdString)
	fmt.Println("")

	thirdString = strconv.Itoa(firstNumber)
	fmt.Printf("%v, %T", thirdString, thirdString)
	fmt.Println("")

}
