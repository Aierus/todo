package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"project/controller"
)

func initServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Starting Server on / \n")
	w.Header().Set("Content-Type", "application/json")
	// frontend.CreateForm(w, r)
	io.WriteString(w, "Welcome to a to-do app written in go")
}

func initTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling a request to /todo\n")
	initTodo := controller.CreateJsonTodoList("A Task", "A Description")
	anotherTodo := controller.CreateJsonTodoList("Another Task", "Another Description")
	w.Write(initTodo)
	w.Write(anotherTodo)
}

func initApi(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Initializing API endpoint \n")
}

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", initServer)    // root and homepage
	mux.HandleFunc("/todo", initTodos) // Where todo json data is stored
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server is closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server %s\n", err)
		panic(err)
	}
}
