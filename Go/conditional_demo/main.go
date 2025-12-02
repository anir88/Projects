//Golang script to show the use of switch case statement

package main

import (
	"fmt"
)

func divisionDemo(numerator int, denominator int) (int, int, error) {
	var err error
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

func main() {
	fmt.Println("Welcome to the divion demo program which showcases the use of conditional switch statements")
	var numerator int = 19
	var denominator int = 2
	var result, remainder, err = divisionDemo(numerator, denominator)
	switch {
	case err != nil:
		fmt.Println(err.Error())

	case remainder == 0:
		fmt.Printf("The result is %v\n", result)

	case remainder != 0:
		fmt.Printf("The result is %v and remainder is %v\n", result, remainder)

	}

}
