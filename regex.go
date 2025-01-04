package main

import (
	"fmt"
	"regexp"
)

func main() {

	exp, err := regexp.Compile("^[a-zA-Z]*$")
	var str string
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter the string to be checked")
	fmt.Scanln(&str)
	matches := exp.MatchString(str)
	fmt.Println("Result : String matched(True) , String not matched(false)")
	fmt.Println(matches) // prints "false"

}
