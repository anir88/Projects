//Golang code to create a ToDo application

package main

import (
	"bufio"
	"fmt"
	"os"
)

var i int
var todoItems []string

func main() {
	fmt.Println("Welcome to the To Do App")
	todoList()

}

func todoList() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the list of ToDo Items for today:")
	for i = range 5 {
		input, _ := reader.ReadString('\n')
		todoItems = append(todoItems, input)
		//continue
	}
	fmt.Println("Your To Do List:")
	for j, item := range todoItems {
		fmt.Printf("%d. %s", j+1, item)
	}
	return todoItems

}
