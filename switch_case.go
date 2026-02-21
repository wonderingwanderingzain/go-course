package main

import "fmt"

func main() {

	// switch statement practice

	// day := "saturday"

	// switch day {
	// case "monday", "tuesday", "wednesday", "thursday", "friday":
	// 	fmt.Println("Its a weekday")
	// case "saturday":
	// 	fmt.Println("Its Saturday 😘")
	// 	fallthrough //it ignores checking next case condition but will execute next code too
	// case "sunday":
	// 	fmt.Println("Its a weekend")
	// default:
	// 	fmt.Println("Error")
	// }

	checkType(23)

}

func checkType(x any) { // any is same as interface{}...do understand later
	// func checkType(x interface{}) {
	switch x.(type) { // switch type...fallthorugh isnt allowed in switch type as every type is different even int and int32 etc
	case int:
		fmt.Println("Its an int")
	case string:
		fmt.Println("Its a string")
	case float64:
		fmt.Println("Its a float")
	default:
		fmt.Println("Unknown type")
	}
}
