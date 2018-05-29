package repositories

import (
    "fmt"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"

    "github.com/bisoelpetjah/go-todo/models"
)

var (
    selectionColumns = "id, message, is_done, created_at, updated_at"
)

func mapRowsToTodo(rows *sql.Rows) (models.Todo, error) {
    var res models.Todo
    err := rows.Scan(
        &res.Id,
        &res.Message,
        &res.IsDone,
        &res.CreatedAt,
        &res.UpdatedAt,
    )

    return res, err
}

type TodoRepository struct {
    db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
    return &TodoRepository {
        db: db,
    }
}

func (repository *TodoRepository) GetTodos() ([]models.Todo, error) {
    rows, err := repository.db.Query(fmt.Sprintf("select %s from todo", selectionColumns))
    if err != nil {
        return nil, err
    }


    todos := make([]models.Todo, 0)
    for rows.Next() {
        todo, err := mapRowsToTodo(rows)
        if err != nil {
            return nil, err
        }

        todos = append(todos, todo)
    }

    return todos, nil
}

func (repository *TodoRepository) GetTodoById(id int64) ([]models.Todo, error) {
    statement, err := repository.db.Prepare(fmt.Sprintf("select %s from todo where id=?", selectionColumns))
    if err != nil {
        return nil, err
    }

    rows, err := statement.Query(id)
    if err != nil {
        return nil, err
    }

    todos := make([]models.Todo, 0)
    for rows.Next() {
        todo, err := mapRowsToTodo(rows)
        if err != nil {
            return nil, err
        }

        todos = append(todos, todo)
    }

    return todos, nil
}

func (repository *TodoRepository) CreateTodo(newTodo models.Todo) ([]models.Todo, error) {
    statement, err := repository.db.Prepare("insert into todo(message) values(?)")
    if err != nil {
        return nil, err
    }

    res, err := statement.Exec(newTodo.Message)
    if err != nil {
        return nil, err
    }

    insertId, err := res.LastInsertId()
    if err != nil {
        return nil, err
    }

    return repository.GetTodoById(insertId)
}

func (repository *TodoRepository) UpdateTodoById(id int64, newTodo models.Todo) ([]models.Todo, error) {
    statement, err := repository.db.Prepare("update todo set message=coalesce(?, message), is_done=coalesce(?, is_done), updated_at=current_timestamp where id=?")
    if err != nil {
        return nil, err
    }

    message, _ := newTodo.Message.Value()
    isDone, _ := newTodo.IsDone.Value()

    _, err = statement.Exec(message, isDone, id)
    if err != nil {
        return nil, err
    }

    return repository.GetTodoById(id)
}
