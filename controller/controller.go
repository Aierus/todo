package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"project/model"
	"time"
)

// CRUD Create, Read (done), Update, Delete

func CreateJsonTodoList(title string, description string, listID int) []byte {
	data := &model.TodoList{
		Title:       title,
		Description: description,
		Done:        false,
		CurrentTime: time.Now(),
		ListID:      listID,
	}
	jsonData, _ := json.Marshal(data)
	// fmt.Printf("jsonData: %s\n", jsonData)
	return jsonData
}

func CreateTodo(title string, description string, listID int) model.TodoList {
	var todoObject model.TodoList
	data := CreateJsonTodoList(title, description, listID)
	payload := bytes.NewBuffer(data)
	// resp, err := http.NewRequest("POST", "http://localhost:8080/todo", payload)
	// fmt.Printf("in controller.go bytes.NewBuffer(data) == %v", payload)
	resp, err := http.Post("http://localhost:8080/todo", "application/json", payload)
	fmt.Printf("response from POST request == %v", resp)
	if err != nil {
		fmt.Printf("Error with POST request to /todo\n")
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &todoObject)
	if err != nil {
		fmt.Printf("Error\n")
		panic(err)
	}
	fmt.Printf("in controller.go the return value is == %+v\n", todoObject)
	return todoObject
	// fmt.Println(string(body))
}

func ReadTodo() []model.TodoList {
	var todos []model.TodoList
	resp, err := http.Get("http://localhost:8080/todo")
	if err != nil {
		fmt.Printf("Error with GET request to /\n")
		panic(err)
	}
	decoder := json.NewDecoder(resp.Body)
	for {
		var todo model.TodoList
		err := decoder.Decode(&todo)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		todos = append(todos, todo)
		// fmt.Printf("%+v\n", todo)
	}
	return todos
	// fmt.Printf("All Todos = %+v\n", todos)
}

func UpdateTodo(title string, description string, done bool, ListID int) {
	todos := ReadTodo()
	var todo model.TodoList
	for _, v := range todos {
		fmt.Printf("v.ListID == %d", v.ListID)
		if v.ListID == ListID {
			todo = v
			fmt.Printf("the todo that the for loop grabs is == %+v\n", todo)
		}
	}
	// reconstruct TodoList with new values
	todo.Title = title
	todo.Description = description
	todo.Done = done
	todo.CurrentTime = time.Now()
	todo.ListID = ListID
	// marshal params
	updatedTodo, _ := json.Marshal(todo)
	// Delete previous json object
	// DeleteTodo(ListID)
	// send back to server
	resp, err := http.NewRequest("POST", "http://localhost:8080/todo", bytes.NewBuffer(updatedTodo))
	if err != nil {
		fmt.Printf("Error updating todo with ListID %d\n", ListID)
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("%+v\n", todos)
}

// Remove Todo w/ ListID "ListID "from todos
func DeleteTodo(ListID int) {
	todos := ReadTodo()
	for _, v := range todos {
		if v.ListID == ListID {
			v = model.TodoList{}
			emptyTodoList, err := json.Marshal(v)
			if err != nil {
				fmt.Printf("Error marshalling empty todoList")
				panic(err)
			}
			resp, err := http.NewRequest("POST", "http://localhost:8080/todo", bytes.NewBuffer(emptyTodoList))
			if err != nil {
				fmt.Printf("Error deleting todo with ListID %d", ListID)
				panic(err)
			}
			defer resp.Body.Close()
		}
	}
	fmt.Printf("%+v\n", todos)
}

// func main() {
// 	resp, err := http.Get("http://localhost:8080/todo")
// fmt.Printf("GET response == %v\n", resp)
// resp2, err2 := http.Get("http://localhost:8080/todo")
// if err != nil || err2 != nil {
// 	fmt.Printf("Error with GET request to /todo \n")
// 	panic(err)
// }
// defer resp.Body.Close()

// Method 1 using Unmarshal
// body, err := ioutil.ReadAll(resp.Body)
// var todos model.TodoList
// json.Unmarshal(body, &todos)

// Method 2 using Decoder
// var todos2 model.TodoList
// decoder := json.NewDecoder(resp2.Body)
// for {
// 	var tempdo model.TodoList
// 	err := decoder.Decode(&tempdo)
// 	if err == io.EOF {
// 		break
// 	}
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", tempdo)
// }
// err3 := decoder.Decode(&todos2)
// if err3 != nil {
// 	switch v := err.(type) {
// 	case *json.SyntaxError:
// 		fmt.Println(string(body[v.Offset-40 : v.Offset]))
// 	default:
// 		panic(err)
// 	}
// }

// defer resp.Body.Close()
// defer resp2.Body.Close()

// fmt.Printf("GET /todo \n")
// fmt.Printf("JSON data on /todo = %+v\n", resp)
// fmt.Printf("Unmarshaled JSON Data %+v\n", todos)
// fmt.Printf("Decoded JSON Data %+v\n", todos2)
// }
