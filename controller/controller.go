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

func CreateJsonTodoList(title string, description string) []byte {
	i := 0
	data := &model.TodoList{
		Title:       title,
		Description: description,
		Done:        false,
		CurrentTime: time.Now(),
		ListID:      i,
	}
	i = i + 1
	jsonData, _ := json.Marshal(data)
	// fmt.Printf("jsonData: %s\n", jsonData)
	return jsonData
}

func CreateTodo(title string, description string) {
	data := CreateJsonTodoList(title, description)
	resp, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error with POST request to /todo\n")
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error\n")
		panic(err)
	}
	fmt.Println(string(body))
}

func ReadTodo() []model.TodoList {
	var todos []model.TodoList
	resp, err := http.Get("http://localhost:8080/todo")
	if err != nil {
		fmt.Printf("Error with GET request to /todo\n")
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

func UpdateTodo(ListID int) {
	todos := ReadTodo()
	fmt.Printf("%+v\n", todos)
}

func DeleteTodo(ListID int) {
	todos := ReadTodo()
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
