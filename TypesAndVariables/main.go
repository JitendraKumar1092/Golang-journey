package main

import "fmt"

const Pi float64 = 3.1415

func main() {
	fmt.Println(Pi) // Pi is accesible
	// Pi = 3.14 // this will cause an error because Pi is a constant
	var name string = "Jeemt"
	// name = 25 // this will cause an error because name is a string
	fmt.Println("Hello from  " + name)           // Hello from  Jeemt
	fmt.Printf("variable's type is: %T\n", name) //	string
	//yeah, we can use %T to get the type of a variable

	var isWorking bool = false

	fmt.Printf("variable's type is: %T\n", isWorking)

	var tinyVal uint8 = 255
	fmt.Println("tinyVal is: ", tinyVal)
	fmt.Printf("variable's type is: %T\n", tinyVal)

	var tinyFloatVal float32 = 22.783563754545
	fmt.Println("tinyVal is: ", tinyFloatVal)
	fmt.Printf("variable's type is: %T\n", tinyFloatVal)

	//aliases
	var newVar int
	fmt.Println(newVar) // 0 is the default value for int
	var newstring string
	fmt.Println(newstring) // "" is the default value for string

	// decaring implicit types
	var nameW = "jeemt ..."
	fmt.Printf("type of name is: %T \n", nameW)

	// without using var , still we can declare variables
	age := 25.663
	fmt.Println("age is: ", age)
}
