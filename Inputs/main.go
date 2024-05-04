package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// take input from user in name variable
	// var name string
	// fmt.Println("Enter your name: ")
	// fmt.Scanln(&name)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your salary: ")
	// comma ok syntex or error ok syntax
	name, _ := reader.ReadString('\n')
	fmt.Println(name)
	fmt.Printf("type of name is %T\n", name)
}
