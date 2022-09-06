package main

import "project/controller"

func main() {

	controller.CreateTodo("Get Eggs", "From the store", 11)
	controller.CreateTodo("Do Homework", "For my class", 1)
	controller.ReadTodo()
	// controller.UpdateTodo("Get Beef", "From the store", false, 0)
	// controller.ReadTodo()

}
