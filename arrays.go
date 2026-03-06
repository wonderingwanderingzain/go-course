package main

import "fmt"

func main() {
	//*****  Arrays ***************

	// A fixed-length collection of elements of the same type. we mostly use Slices as arrays are rarely used but still learn. Arrays do not change size

	//  var arrayName=[size]elementType
	// [5]is part of type. so complete type is [5]int. they are not seprate. thats why we cant compare arrays of different size as they are of different type.arr[4]!=arr[5]
	//to find length of array len(arr) . returns size . if you need resizing use slices

	var numbers [5]int
	fmt.Println(numbers)

	numbers[0] = 20
	fmt.Println(numbers)

	numbers[4] = 1
	fmt.Println(numbers)

	fruits := [4]string{"Apples", "Grapes", "Watermelon", "Banana"}
	fmt.Printf("Third element is %v and last is %v\n", fruits[2], fruits[len(fruits)-1])
}
