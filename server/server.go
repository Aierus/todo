package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func initServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling a request on /\n")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "Welcome to a to-do app written in Go!\n")
}

func initTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling a request to /todo\n")

	switch r.Method {
	case "GET":
		//
	case "POST":
		respBody, err := ioutil.ReadAll(r.Body)
		// io.WriteString(w, string(respBody))
		// initTodo := controller.CreateJsonTodoList("A Task", "A Description", 21)
		// fmt.Printf("\n the manual json list     shows as == %s\n", initTodo)
		fmt.Printf("\n in server.go the request shows as == %s\n", respBody)
		// w.Write(initTodo) // shows up when curl http://localhost:8080/todo
		w.Write(respBody) // ??? Doesn't show up when curl http://localhost:8080/todo
		// Must be different types!?
		if err != nil {
			fmt.Printf("Error reading request body")
			panic(err)
		}
		_, errw := w.Write(respBody)
		if errw != nil {
			fmt.Printf("err writing\n")
		}
	}
	defer r.Body.Close()
	// initTodo := controller.CreateJsonTodoList("A Task", "A Description", 21)
	// anotherTodo := controller.CreateJsonTodoList("Another Task", "Another Description", 32)
	// w.Write(initTodo)
	// w.Write(anotherTodo)
}

func initApi(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Initializing API endpoint \n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", initServer) // root and homepage
	// mux.HandleFunc("/todo", initTodos) // Where todo json data is stored
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server is closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server %s\n", err)
		panic(err)
	}
}
