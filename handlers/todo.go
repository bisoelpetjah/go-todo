package handlers

import (
    "strconv"
    "net/http"
    "encoding/json"
    "database/sql"

    "github.com/gorilla/mux"

    "github.com/bisoelpetjah/go-todo/response"
    "github.com/bisoelpetjah/go-todo/models"
    "github.com/bisoelpetjah/go-todo/repositories"
)

type TodoHandler struct {
    repository *repositories.TodoRepository
}

func NewTodoHandler(db *sql.DB) *TodoHandler {
    repository := repositories.NewTodoRepository(db)
    return &TodoHandler {
        repository: repository,
    }
}

func (handler *TodoHandler) GetTodos(res http.ResponseWriter, req *http.Request) {
    todos, err := handler.repository.GetTodos()
    if err != nil {
        response.CreateErrorResponse(res, http.StatusInternalServerError, err.Error())
        return
    }

    response.CreateSuccessResponse(res, todos)
}

func (handler *TodoHandler) GetTodoById(res http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    idString := params["id"]

    id, err := strconv.ParseInt(idString, 10, 64)
    if err != nil {
        response.CreateErrorResponse(res, http.StatusNotFound, "Not found")
        return
    }

    todos, err := handler.repository.GetTodoById(id)
    if err != nil {
        response.CreateErrorResponse(res, http.StatusInternalServerError, err.Error())
        return
    }

    if len(todos) == 0 {
        response.CreateErrorResponse(res, http.StatusNotFound, "Not found")
        return
    }

    response.CreateSuccessResponse(res, todos[0])
}

func (handler *TodoHandler) CreateTodo(res http.ResponseWriter, req *http.Request) {
    newTodo := models.Todo {}
    err := json.NewDecoder(req.Body).Decode(&newTodo)
    if err != nil {
        response.CreateErrorResponse(res, http.StatusBadRequest, "Bad request")
        return
    }

    if !newTodo.Message.Valid {
        response.CreateErrorResponse(res, http.StatusBadRequest, "Message field is required")
        return
    }

    todos, err := handler.repository.CreateTodo(newTodo)
    if err != nil {
        response.CreateErrorResponse(res, http.StatusInternalServerError, "Internal server error")
        return
    }

    response.CreateSuccessResponse(res, todos[0])
}

func (handler *TodoHandler) UpdateTodoById(res http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    idString := params["id"]

    id, err := strconv.ParseInt(idString, 10, 64)
    if err != nil {
        response.CreateErrorResponse(res, http.StatusNotFound, "Not found")
        return
    }

    newTodo := models.Todo {}
    err = json.NewDecoder(req.Body).Decode(&newTodo)
    if (err != nil) {
        response.CreateErrorResponse(res, http.StatusBadRequest, "Bad request")
        return
    }

    todos, err := handler.repository.UpdateTodoById(id, newTodo)
    if err != nil {
        response.CreateErrorResponse(res, http.StatusInternalServerError, "Internal server error")
        return
    }

    if len(todos) == 0 {
        response.CreateErrorResponse(res, http.StatusNotFound, "Not found")
        return
    }

    response.CreateSuccessResponse(res, todos[0])
}
