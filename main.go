package main

import (
	"project/controller"
)

func main() {

	controller.CreateTodo("Get Eggs", "From the store", 0)
	controller.CreateTodo("Do Homework", "For my class", 1)
	controller.ReadTodo()

}
