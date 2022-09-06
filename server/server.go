package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"project/model"
)

func initServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling a request on /\n")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "Welcome to a to-do app written in Go!\n")
}

// func respondToPost(w http.ResponseWriter, r *http.Request, response model.TodoList) {
// 	fmt.Printf("\n in 'respondToPost' the TodoList passed as a parameter shows as == %+v\n", response)
// 	respBody, err := ioutil.ReadAll(r.Body)
// 	fmt.Printf("\n in 'respondToPost' the request shows as == %s\n", respBody)
// 	if err != nil {
// 		panic(err)
// 	}
// 	w.Write(respBody)
// }

func addLiveTodoList(todo model.TodoList) []model.TodoList {
	var liveTodoList []model.TodoList

	// if liveTodoList == nil {
	// 	controller.CreateTodo("Default Todo", "Default Description", 0)
	// }

	liveTodoList = append(liveTodoList, todo)

	fmt.Printf("Live Todo List == %+v\n", liveTodoList)

	return liveTodoList
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling a request to /todo\n")
	w.Header().Set("Content-Type", "application/json")

	var tempTodoList model.TodoList

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &tempTodoList)
	addLiveTodoList(tempTodoList)

	// 	fmt.Fprintf(w, string(data)
	// 	fmt.Printf("%s\n", data)
	// 	w.Write(data)
	// io.WriteString(w, string(data))
	// io.WriteString(w, "testing to see if curl works")

	// When you curl http://localhost:8080/todo this shows up
	// localJsonObject := controller.CreateJsonTodoList("Get Eggs", "From the store", 11)
	// w.Write(localJsonObject)

}

// func initTodos(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Printf("Handling a request to /todo\n")
// 	// var todos model.TodoList

// 	respBody, err := ioutil.ReadAll(r.Body)
// 	w.Write(respBody)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// json.Unmarshal(respBody, &todos)
// 	fmt.Printf("\n in initTodos the responseBody shows as == %s\n", respBody)
// 	// fmt.Printf("in initTodos the 'todos' object when unmarshaled shows as %+v\n", todos)
// 	// jsonObject, _ := json.Marshal(todos)
// 	w.Write(respBody)

// 	// w.Write(jsonObject)

// 	// switch r.Method {
// 	// case "GET":
// 	// 	// Todo (haha)
// 	// case "POST":
// 	// 	respBody, err := ioutil.ReadAll(r.Body)
// 	// 	json.Unmarshal(respBody, &todos)
// 	// 	fmt.Printf("\n in initTodos the request shows as == %s\n", respBody)
// 	// 	fmt.Printf("the todos object when unmarshaled shows as %+v\n", todos)
// 	// 	jsonObject, _ := json.Marshal(todos)
// 	// 	w.Write(jsonObject)

// 	// 	// initTodo := controller.CreateJsonTodoList("A Task", "A Description", 21)
// 	// 	// w.Write(initTodo) // shows up when curl http://localhost:8080/todo
// 	// 	w.Write(respBody) // ??? Doesn't show up when curl http://localhost:8080/todo
// 	// 	// Must be different types!?
// 	// 	if err != nil {
// 	// 		fmt.Printf("Error reading request body")
// 	// 		panic(err)
// 	// 	}
// 	// 	_, errw := w.Write(respBody)
// 	// 	if errw != nil {
// 	// 		fmt.Printf("err writing\n")
// 	// 	}
// 	// }
// 	defer r.Body.Close()

// 	// initTodo := controller.CreateJsonTodoList("A Task", "A Description", 21)
// 	// anotherTodo := controller.CreateJsonTodoList("Another Task", "Another Description", 32)
// 	// w.Write(initTodo)
// 	// w.Write(anotherTodo)
// }

func initApi(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Initializing API endpoint \n")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", initServer)      // root and homepage
	mux.HandleFunc("/todo", handleTodos) // Where todo json data is stored
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server is closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server %s\n", err)
		panic(err)
	}

	// time.Sleep(500 * time.Millisecond)

	// requestUrl := fmt.Sprintf("http://localhost:%d%s", 8080, "/todo")
	// req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	// if err != nil {
	// 	fmt.Printf("could not create request: %s\n", err)
	// 	panic(err)
	// }

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Printf("client: error making http request %s\n", err)
	// }

	// fmt.Printf("Client: status code %d\n", res.StatusCode)
	// resBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("client: response body: %s\n", resBody)
}
