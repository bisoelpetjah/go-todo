package main

import (
    "fmt"
    "net/http"
    "database/sql"

    "github.com/kelseyhightower/envconfig"

    "github.com/bisoelpetjah/go-todo/config"
    "github.com/bisoelpetjah/go-todo/router"
)

var (
    appConfig config.AppConfig
    dbConfig config.DbConfig
)

func createConfig()  {
    envconfig.Process("", &appConfig)
    envconfig.Process("", &dbConfig)
}

func main()  {
    createConfig()

    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/todo?parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port))

    if err != nil {
        return
    }

    defer db.Close()

    router := router.CreateRouter(db)

    http.Handle("/", router)
    http.ListenAndServe(":" + appConfig.Port, nil)
}
