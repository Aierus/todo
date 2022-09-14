package sql

import (
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
	"project/model"
	"strconv"
)

type Repository interface {
	Migrate() error
	Create(todo model.TodoList) (*model.TodoList, error)
	All() ([]model.TodoList, error)
	GetByName(name string) (*model.TodoList, error)
	Update(ListID int, updated model.TodoList) (*model.TodoList, error)
	Delete(ListID int) error
}

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS todos(
        title TEXT NOT NULL UNIQUE,
        description TEXT NOT NULL UNIQUE,
        done TEXT NOT NULL,
		time TEXT NOT NULL,
        listid INTEGER PRIMARY KEY NOT NULL
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(todo model.TodoList) (*model.TodoList, error) {
	res, err := r.db.Exec("INSERT INTO todos(title, description, done, time, listid) values(?,?,?,?,?)", todo.Title, todo.Description, strconv.FormatBool(todo.Done), todo.CurrentTime.String(), todo.ListID)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	todo.ListID = id

	return &todo, nil
}

func (r *SQLiteRepository) All() ([]model.TodoList, error) {
	rows, err := r.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []model.TodoList
	for rows.Next() {
		var todo model.TodoList
		if err := rows.Scan(&todo.Title, &todo.Description, &todo.Done, &todo.CurrentTime, &todo.ListID); err != nil {
			return nil, err
		}
		all = append(all, todo)
	}
	return all, nil
}

func (r *SQLiteRepository) GetByTitle(title string) (*model.TodoList, error) {
	row := r.db.QueryRow("SELECT * FROM todos WHERE title = ?", title)

	var todo model.TodoList
	if err := row.Scan(&todo.Title, &todo.Description, &todo.Done, &todo.CurrentTime, &todo.ListID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &todo, nil
}

func (r *SQLiteRepository) Update(id int64, updated model.TodoList) (*model.TodoList, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	res, err := r.db.Exec("UPDATE websites SET name = ?, url = ?, rank = ? WHERE id = ?", updated.Title, updated.Description, strconv.FormatBool(updated.Done), updated.CurrentTime.String(), id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *SQLiteRepository) Delete(id int64) error {
	res, err := r.db.Exec("DELETE FROM websites WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
