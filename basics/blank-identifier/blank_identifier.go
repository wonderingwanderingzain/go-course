package main

import "fmt"

func main() {

	// Blank Identifier  _

	// it is used to ignore some value...Because Go does NOT allow unused variables. This tells compiler: Yes, I know there is a value. I choose not to use it.

	x := 10
	_ = x // compiler doesnt allow unused var

	//ignore index
	arr := [2]string{"Zain", "Ali"}
	for _, value := range arr {
		fmt.Println(value)
	}

	// ignore return value or error
	// value,_ :=someFunction()

	//  learn some advance blank identifier use cases/ patterns as you progress
}
