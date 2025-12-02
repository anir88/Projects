//Golang script to demonstrate arrays and slices

package main

import "fmt"

func main() {
	fmt.Println("This script shows the usage of the arrays and slices")
	var arrDemo1 [3]int32 = [3]int32{1, 2, 3}
	autoArrDemo1 := [...]int32{7, 8, 9} // automated way of declaring and describing the array length
	fmt.Printf("The array with hard coded length is %v\nThe length of the array is %v\n", arrDemo1, len(arrDemo1))
	fmt.Printf("The array with autocalculated length is %v\nThe length of the array is %v\n", autoArrDemo1, len(autoArrDemo1))
	var slicedemo1 = []int32{10, 11, 12}
	fmt.Printf("The type of variable slicedemo1 is %T\nThe value of the slice is %v\nThe length of the sice is %v\n", slicedemo1, slicedemo1, len(slicedemo1))

}
