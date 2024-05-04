package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter a number between 100 and 1000: ")

	reader := bufio.NewReader(os.Stdin)
	// comma ok syntex or error ok syntax
	// initializing reader.ReadString('\n') to num
	num, _ := reader.ReadString('\n')
	fmt.Println("number entered is :", num)
	// converting string to float using strconv.ParseFloat

	numFloat, err := strconv.ParseFloat(strings.TrimSpace(num), 32)
	// checking of error
	if err != nil {
		fmt.Println("Error:", err)

	}
	// boom .... done
	fmt.Println("number entered is :", numFloat*2)
}
