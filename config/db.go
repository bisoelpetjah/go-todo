package config

type DbConfig struct {
    Host string `envconfig:"DB_HOST" default:"go-todo-dev-mysql"`
    Port string `envconfig:"DB_PORT" default:"3306"`
    User string `envconfig:"DB_USER" default:"dev"`
    Password string `envconfig:"DB_PASSWORD" default:"dev"`
}
