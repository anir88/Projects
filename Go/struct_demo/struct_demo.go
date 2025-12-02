//Golang script to demonstrate the functioning of a struct

package main

import (
	"fmt"
)

func main() {
	fmt.Println("This script shows the working of struct in Golang")
	type Person struct {
		name  string
		age   int
		city  string
		email string
	}
	rec1 := Person{"Anirban", 37, "Bangalore", "anir.siliguri@gmail.com"}
	rec2 := Person{"Avantika", 36, "Bangalore", "avantikabose@gmail.com"}
	fmt.Printf("The records of the first person are %v\n", rec1)
	fmt.Printf("The records of the second person are %v\n", rec2)
	fmt.Printf("The age of the %v is %v\n", rec1.name, rec1.age)

}
