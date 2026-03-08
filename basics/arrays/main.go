package main

import "fmt"

func main() {
	//*****  Arrays ***************

	// A fixed-length collection of elements of the same type. we mostly use Slices as arrays are rarely used but still learn. Arrays do not change size

	//  var arrayName=[size]elementType
	// [5]is part of type. so complete type is [5]int. they are not seprate. thats why we cant compare arrays of different size as they are of different type.arr[4]!=arr[5] therefore arrays can be compared directly unlike slices
	//to find length of array len(arr) . returns size . if you need resizing use slices
	//avoid large arrays and keep them small and better use range for iteration

	var numbers [5]int
	fmt.Println(numbers)

	numbers[0] = 20
	fmt.Println(numbers)

	numbers[4] = 1
	fmt.Println(numbers)

	fruits := [4]string{"Apples", "Grapes", "Watermelon", "Banana"}
	fmt.Printf("Third element is %v and last is %v\n", fruits[2], fruits[len(fruits)-1])

	arr := [...]int{10, 20, 30} //go automatically counts size>>>>length 3
	_ = arr

	// partial initialization
	arr1 := [5]int{1, 2} // it becomes {1,2,0,0,0}
	_ = arr1

	// index specific initialization
	arr2 := [5]int{1: 20, 4: 50} //it becomes {0,10,0,0,50}
	_ = arr2

	a := [3]int{10, 20, 30}
	b := a
	//Arrays are value types...b is copy of a.....modifying b doesnt change a
	//if u want to modify orignl use * and call with &
	// * is pointer and it points to address
	b[0] = 100
	fmt.Println("without * and &")
	fmt.Printf("Orignal array is %v\nwhile copied array is %v\n", a, b)

	c := [3]int{10, 20, 30}
	var d *[3]int
	d = &c
	d[0] = 100
	fmt.Println("with * and &")
	fmt.Printf("Orignal array is %v\nwhile copied array is %v\n", c, *d)

	// looping array with range
	for _, value := range a {
		fmt.Println(value)
	}

	//2D Arrays
	//used for matrices mostly
	// Arrays are used rarely and 2D arrays very rarely but still know that they exist
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)
}
