package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"project/db"
)

const dbName = "sqlite.db"

func main() {
	// controller.CreateTodo("Get Eggs", "From the store", 11)
	// controller.CreateTodo("Do Homework", "For my class", 1)
	// controller.ReadTodo()
	// controller.UpdateTodo("Get Beef", "From the store", false, 0)
	// controller.ReadTodo()

	os.Remove(dbName)

	database, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}

	todoRepository := db.NewSQLiteRepository(database)

	if err := todoRepository.Migrate(); err != nil {
		log.Fatal(err)
	}
}
