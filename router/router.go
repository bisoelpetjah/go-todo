package router

import (
    "net/http"
    "database/sql"

    "github.com/gorilla/mux"

    "github.com/bisoelpetjah/go-todo/response"
    "github.com/bisoelpetjah/go-todo/handlers"
)

func indexHandler(req http.ResponseWriter, res *http.Request)  {
    response.CreateSuccessResponse(req, map[string]interface {} {
        "message": "API is live!",
    })
}

func initHandlers(router *mux.Router, db *sql.DB) {
    todoHandler := handlers.NewTodoHandler(db)

    router.HandleFunc("/api/todos", todoHandler.GetTodos).Methods("GET")
    router.HandleFunc("/api/todos/{id}", todoHandler.GetTodoById).Methods("GET")
    router.HandleFunc("/api/todos", todoHandler.CreateTodo).Methods("POST")
    router.HandleFunc("/api/todos/{id}", todoHandler.UpdateTodoById).Methods("PATCH")

}

func notFoundHandler(res http.ResponseWriter, req *http.Request)  {
    response.CreateErrorResponse(res, http.StatusNotFound, "Not found")
}

func CreateRouter(db *sql.DB) *mux.Router  {
    router := mux.NewRouter()

    router.HandleFunc("/", indexHandler)

    initHandlers(router, db)

    router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

    return router
}
